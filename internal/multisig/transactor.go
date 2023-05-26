package multisig

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/eth"
)

type Confirmer = func(ctx context.Context, from, destination common.Address, value *big.Int, data []byte) error

type Transactor struct {
	*Caller
	transactor      *contract.MultiSigWalletWithDailyLimitTransactor
	filterer        *contract.MultiSigWalletWithDailyLimitFilterer
	contractAddress common.Address
	senderAddress   common.Address
	senderSigner    bind.SignerFn
	waiter          eth.Waiter
}

type TransactorBackend interface {
	CallerBackend
	bind.ContractTransactor
	eth.WaitBackend
}

func NewTransactor(backend TransactorBackend, contractAddress, senderAddress common.Address, senderSigner bind.SignerFn, waiter eth.Waiter) (*Transactor, error) {
	if waiter == nil {
		waiter = eth.SilentWaiter(backend)
	}

	caller, err := NewCaller(backend, contractAddress)
	if err != nil {
		return nil, err
	}

	transactor, err := contract.NewMultiSigWalletWithDailyLimitTransactor(contractAddress, backend)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	filterer, err := contract.NewMultiSigWalletWithDailyLimitFilterer(contractAddress, backend)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return &Transactor{
		Caller:          caller,
		transactor:      transactor,
		filterer:        filterer,
		contractAddress: contractAddress,
		senderAddress:   senderAddress,
		senderSigner:    senderSigner,
		waiter:          waiter,
	}, nil
}

func (t *Transactor) Sender() common.Address {
	return t.senderAddress
}

func (t *Transactor) SubmitAddOwner(ctx context.Context, owner common.Address) (uint64, error) {
	if owner == (common.Address{}) {
		return 0, errs.New("owner address is unset")
	}

	data, err := multisigABI.Pack("addOwner", owner)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, t.contractAddress, zero, data)
}

func (t *Transactor) SubmitRemoveOwner(ctx context.Context, owner common.Address) (uint64, error) {
	if owner == (common.Address{}) {
		return 0, errs.New("owner address is unset")
	}

	data, err := multisigABI.Pack("removeOwner", owner)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, t.contractAddress, zero, data)
}

func (t *Transactor) SubmitReplaceOwner(ctx context.Context, oldOwner, newOwner common.Address) (uint64, error) {
	if oldOwner == (common.Address{}) {
		return 0, errs.New("old owner address is unset")
	}
	if newOwner == (common.Address{}) {
		return 0, errs.New("new owner address is unset")
	}

	data, err := multisigABI.Pack("replaceOwner", oldOwner, newOwner)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, t.contractAddress, zero, data)
}

func (t *Transactor) SubmitChangeRequirement(ctx context.Context, requirement uint64) (uint64, error) {
	// TODO: stronger input validation (i.e. mimic the requirements check on
	// the contract function) for improved diagnostics.

	data, err := multisigABI.Pack("changeRequirement", newUint64(requirement))
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, t.contractAddress, zero, data)
}

func (t *Transactor) SubmitETHTransfer(ctx context.Context, recipient common.Address, amount *big.Int) (uint64, error) {
	return t.SubmitTransaction(ctx, recipient, amount, nil)
}

func (t *Transactor) SubmitTokenTransfer(ctx context.Context, tokenContractAddress, recipient common.Address, amount *big.Int) (uint64, error) {
	data, err := erc20ABI.Pack("transfer", recipient, amount)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, tokenContractAddress, zero, data)
}

func (t *Transactor) SubmitTokenTransferFrom(ctx context.Context, tokenContractAddress, from, to common.Address, amount *big.Int) (uint64, error) {
	data, err := erc20ABI.Pack("transferFrom", from, to, amount)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, tokenContractAddress, zero, data)
}

func (t *Transactor) SubmitTokenApprove(ctx context.Context, tokenContractAddress, spender common.Address, amount *big.Int) (uint64, error) {
	data, err := erc20ABI.Pack("approve", spender, amount)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, tokenContractAddress, zero, data)
}

func (t *Transactor) SubmitSetUpgradeMaster(ctx context.Context, tokenContractAddress, master common.Address) (uint64, error) {
	data, err := upgradeableTokenABI.Pack("setUpgradeMaster", master)
	if err != nil {
		return 0, err
	}

	return t.SubmitTransaction(ctx, tokenContractAddress, zero, data)
}

func (t *Transactor) ConfirmTransaction(ctx context.Context, transactionID uint64) (uint64, error) {
	opts := t.transactOpts(ctx)

	tx, err := t.transactor.ConfirmTransaction(opts, newUint64(transactionID))
	if err != nil {
		if !isExecutionReverted(err) {
			return 0, errs.New("failed to confirm transaction: %v", err)
		}
		return 0, errs.New("failed to confirm transaction %d: %v: is the caller an owner?", transactionID, err)
	}
	return t.waiter.Wait(ctx, tx.Hash())
}

func (t *Transactor) ExecuteTransaction(ctx context.Context, transactionID uint64) error {
	opts := t.transactOpts(ctx)

	tx, err := t.transactor.ExecuteTransaction(opts, newUint64(transactionID))
	if err != nil {
		return errs.New("failed to execute transaction: %v", err)
	}
	_, err = t.waiter.Wait(ctx, tx.Hash())
	return err
}

func (t *Transactor) RevokeConfirmation(ctx context.Context, transactionID uint64) error {
	opts := t.transactOpts(ctx)

	tx, err := t.transactor.RevokeConfirmation(opts, newUint64(transactionID))
	if err != nil {
		return errs.New("failed to revoke confirmation: %v", err)
	}
	_, err = t.waiter.Wait(ctx, tx.Hash())
	return err
}

func (t *Transactor) SubmitTransaction(ctx context.Context, destination common.Address, value *big.Int, data []byte) (uint64, error) {
	opts := t.transactOpts(ctx)

	tx, err := t.transactor.SubmitTransaction(opts, destination, value, data)
	if err != nil {
		if isExecutionReverted(err) {
			return 0, errs.New("failed to submit transaction: invalid input: %v", err)
		}
		return 0, errs.New("failed to submit transaction: %v", err)
	}

	blockNumber, err := t.waiter.Wait(ctx, tx.Hash())
	if err != nil {
		return 0, err
	}

	return t.getTransactionIDFromSubmissionEvent(ctx, blockNumber, tx.Hash())
}

func (t *Transactor) getTransactionIDFromSubmissionEvent(ctx context.Context, blockNumber uint64, txHash common.Hash) (uint64, error) {
	it, err := t.filterer.FilterSubmission(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: ctx,
	}, nil)
	if err != nil {
		return 0, errs.New("failed to filter submission event")
	}
	defer func() {
		_ = it.Close()
	}()

	for it.Next() {
		if it.Event.Raw.TxHash != txHash {
			continue
		}
		if !it.Event.TransactionId.IsUint64() {
			// purely defensive, transaction IDs should be very small since
			// they start at zero and are are incremented by one for each
			// submission.
			return 0, errs.New("transaction ID %s does not fit in a uint64", it.Event.TransactionId)
		}
		return it.Event.TransactionId.Uint64(), nil
	}
	if err := it.Error(); err != nil {
		return 0, errs.New("failed to iterate submission events: %v", err)
	}

	// This would be very unexpected.
	return 0, errs.New("transaction ID not found for submission")
}

func (t *Transactor) transactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:    t.senderAddress,
		Signer:  t.senderSigner,
		Context: ctx,
	}
}
