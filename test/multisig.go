package test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/multisig"
)

type MultiSigHarness struct {
	ContractAddress common.Address
	Caller          *multisig.Caller

	nodeURL              string
	chainID              *big.Int
	client               *ethclient.Client
	tokenContractAddress common.Address
}

func (h *MultiSigHarness) AssertIsOwner(t *testing.T, account common.Address, expected bool) {
	actual, err := h.Caller.IsOwner(context.Background(), account)
	require.NoError(t, err)
	require.Equal(t, expected, actual, "unexpected ownership status of %s", account)
}

func (h *MultiSigHarness) AssertOwners(t *testing.T, expected []common.Address) {
	require.Equal(t, expected, h.ListOwners(t))
}

func (h *MultiSigHarness) AssertConfirmationRequirement(t *testing.T, expected uint64) {
	require.Equal(t, expected, h.GetConfirmationRequirement(t), "unexpected confirmation requirement")
}

func (h *MultiSigHarness) AssertTransactionCounts(t *testing.T, pending int, executed int) {
	t.Helper()
	assert.Len(t, h.ListPendingTransactions(t), pending, "unexpected pending transaction count")
	assert.Len(t, h.ListExecutedTransactions(t), executed, "unexpected executed transaction count")
}

func (h *MultiSigHarness) ConfirmTransaction(t *testing.T, initiator *ecdsa.PrivateKey, transactionID uint64) {
	_, err := h.newTransactor(t, initiator).ConfirmTransaction(context.Background(), transactionID)
	require.NoError(t, err)
}

func (h *MultiSigHarness) SubmitAddOwner(t *testing.T, initiator *ecdsa.PrivateKey, newOwner common.Address) uint64 {
	transactionID, err := h.newTransactor(t, initiator).SubmitAddOwner(context.Background(), newOwner)
	require.NoError(t, err)
	return transactionID
}

func (h *MultiSigHarness) SubmitApprove(t *testing.T, initiator *ecdsa.PrivateKey, spender common.Address, amount string) uint64 {
	transactionID, err := h.newTransactor(t, initiator).SubmitTokenApprove(context.Background(), h.tokenContractAddress, spender, bigString(t, amount))
	require.NoError(t, err)
	return transactionID
}

func (h *MultiSigHarness) SubmitTransfer(t *testing.T, initiator *ecdsa.PrivateKey, recipient common.Address, amount string) uint64 {
	transactionID, err := h.newTransactor(t, initiator).SubmitTokenTransfer(context.Background(), h.tokenContractAddress, recipient, bigString(t, amount))
	require.NoError(t, err)
	return transactionID
}

func (h *MultiSigHarness) ListPendingTransactions(t *testing.T) []multisig.Transaction {
	txs, err := h.Caller.ListTransactions(context.Background(), true, false)
	require.NoError(t, err)
	return txs
}

func (h *MultiSigHarness) ListExecutedTransactions(t *testing.T) []multisig.Transaction {
	txs, err := h.Caller.ListTransactions(context.Background(), false, true)
	require.NoError(t, err)
	return txs
}

func (h *MultiSigHarness) ListOwners(t *testing.T) []common.Address {
	owners, err := h.Caller.ListOwners(context.Background())
	require.NoError(t, err)
	return owners
}

func (h *MultiSigHarness) GetConfirmationRequirement(t *testing.T) uint64 {
	requirement, err := h.Caller.GetConfirmationRequirement(context.Background())
	require.NoError(t, err)
	return requirement
}

func (h *MultiSigHarness) IsTransactionConfirmed(t *testing.T, transactionID uint64) bool {
	confirmed, err := h.Caller.IsTransactionConfirmed(context.Background(), transactionID)
	require.NoError(t, err)
	return confirmed
}

func (h *MultiSigHarness) GetTransactionConfirmations(t *testing.T, transactionID uint64) []common.Address {
	confirmations, err := h.Caller.GetTransactionConfirmations(context.Background(), transactionID)
	require.NoError(t, err)
	return confirmations
}

func (h *MultiSigHarness) DumpEvents(t *testing.T) {
	events, err := h.Caller.GetEvents(context.Background())
	require.NoError(t, err)
	for i, event := range events {
		t.Logf("EVENT(%d): %s", i, event)
	}
}

func (h *MultiSigHarness) NewTransactor(t *testing.T, account int) *multisig.Transactor {
	require.Less(t, account, len(AccountKey), "invalid account index")
	return h.newTransactor(t, AccountKey[account])
}

func (h *MultiSigHarness) newTransactor(t *testing.T, key *ecdsa.PrivateKey) *multisig.Transactor {
	opts, err := bind.NewKeyedTransactorWithChainID(key, h.chainID)
	require.NoError(t, err)
	waiter := eth.ProgressWaiter(h.client, WaitProgress(t))
	transactor, err := multisig.NewTransactor(h.client, h.ContractAddress, opts, waiter)
	require.NoError(t, err)
	return transactor
}

func deployMultiSigContract(t *testing.T, opts *bind.TransactOpts, client *ethclient.Client) common.Address {
	owners := AccountAddress[0:2]
	required := int64(2)
	dailyLimit := int64(0)

	contractAddress, err := multisig.DeployContract(opts, client, owners, required, dailyLimit, eth.ProgressWaiter(client, WaitProgress(t)))
	require.NoError(t, err)
	return contractAddress
}
