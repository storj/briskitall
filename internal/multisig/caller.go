package multisig

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
)

type Caller struct {
	caller   *contract.MultiSigWalletWithDailyLimitCaller
	filterer *contract.MultiSigWalletWithDailyLimitFilterer
}

type CallerBackend interface {
	bind.ContractCaller
	bind.ContractFilterer
}

func NewCaller(backend CallerBackend, contractAddress common.Address) (*Caller, error) {
	caller, err := contract.NewMultiSigWalletWithDailyLimitCaller(contractAddress, backend)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	filterer, err := contract.NewMultiSigWalletWithDailyLimitFilterer(contractAddress, backend)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &Caller{
		caller:   caller,
		filterer: filterer,
	}, nil
}

func (c *Caller) IsOwner(ctx context.Context, account common.Address) (bool, error) {
	opts := c.callOpts(ctx)
	return c.caller.IsOwner(opts, account)
}

func (c *Caller) ListOwners(ctx context.Context) ([]common.Address, error) {
	opts := c.callOpts(ctx)
	maxOwnerCount, err := c.caller.MAXOWNERCOUNT(opts)
	if err != nil {
		return nil, errs.New("failed to get max owner count: %v", err)
	}

	var owners []common.Address
	index := big.NewInt(0)
	for index.Cmp(maxOwnerCount) != 0 {
		owner, err := c.caller.Owners(opts, index)
		if err != nil {
			if isExecutionReverted(err) {
				return owners, nil
			}
			return nil, errs.New("failed to get owner %d: %v", index, err)
		}
		index = index.Add(index, one)
		owners = append(owners, owner)
	}
	return owners, nil
}

func (c *Caller) GetConfirmationRequirement(ctx context.Context) (uint64, error) {
	opts := c.callOpts(ctx)
	required, err := c.caller.Required(opts)
	if err != nil {
		return 0, errs.New("failed to check required count: %v", err)
	}
	if !required.IsUint64() {
		// This should be purely defensive. The required count of confirmations
		// should be a very small amount.
		return 0, errs.New("required count (%s) does not fit in a uint64", required)
	}
	return required.Uint64(), nil
}

func (c *Caller) IsTransactionConfirmed(ctx context.Context, transactionID uint64) (bool, error) {
	opts := c.callOpts(ctx)
	confirmed, err := c.caller.IsConfirmed(opts, newUint64(transactionID))
	if err != nil {
		return false, errs.New("failed to get confirmation status for transaction %d: %v", transactionID, err)
	}
	return confirmed, nil
}

func (c *Caller) FindTransaction(ctx context.Context, transactionID uint64) (*Transaction, error) {
	return findTransaction(ctx, c.caller, newUint64(transactionID))
}

func (c *Caller) GetTransactionConfirmations(ctx context.Context, transactionID uint64) ([]common.Address, error) {
	opts := c.callOpts(ctx)

	confirmations, err := c.caller.GetConfirmations(opts, newUint64(transactionID))
	if err != nil {
		return nil, errs.New("failed to get confirmations for transaction %d: %v", transactionID, err)
	}
	return confirmations, nil
}

func (c *Caller) GetTransactionEvents(ctx context.Context, transactionID uint64) ([]Event, error) {
	opts := c.filterOpts(ctx)

	filterByID := []*big.Int{newUint64(transactionID)}

	events, err := getTransactionEvents(opts, c.filterer, filterByID)
	if err != nil {
		return nil, err
	}

	sortEvents(events)
	return events, nil
}

func (c *Caller) GetEvents(ctx context.Context) ([]Event, error) {
	opts := c.filterOpts(ctx)

	events, err := getAllEvents(opts, c.filterer)
	if err != nil {
		return nil, err
	}

	sortEvents(events)
	return events, nil
}

func (c *Caller) ListTransactionIDs(ctx context.Context, pending, executed bool) ([]uint64, error) {
	opts := c.callOpts(ctx)
	count, err := c.caller.GetTransactionCount(opts, pending, executed)
	if err != nil {
		return nil, errs.New("failed to get transaction count: %v", err)
	}

	ids, err := c.caller.GetTransactionIds(opts, zero, count, pending, executed)
	if err != nil {
		return nil, errs.New("failed to get transaction ids: %v", err)
	}

	var out []uint64
	for _, id := range ids {
		out = append(out, id.Uint64())
	}
	return out, nil
}

func (c *Caller) ListTransactions(ctx context.Context, pending, executed bool) ([]Transaction, error) {
	opts := c.callOpts(ctx)
	count, err := c.caller.GetTransactionCount(opts, pending, executed)
	if err != nil {
		return nil, errs.New("failed to get transaction count: %v", err)
	}

	ids, err := c.caller.GetTransactionIds(opts, zero, count, true, false)
	if err != nil {
		return nil, errs.New("failed to get transaction ids: %v", err)
	}

	txs := make([]Transaction, 0, len(ids))
	for _, id := range ids {
		tx, err := getTransaction(ctx, c.caller, id)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (c *Caller) callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func (c *Caller) filterOpts(ctx context.Context) *bind.FilterOpts {
	return &bind.FilterOpts{Context: ctx}
}
