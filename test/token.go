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

	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/token"
)

var (
	InitialSupply = big.NewInt(100000000000)
)

type TokenHarness struct {
	ContractAddress common.Address
	Caller          *token.Caller

	chainID *big.Int
	client  *ethclient.Client
}

func (h *TokenHarness) AssertBalance(t *testing.T, account common.Address, expected string) {
	t.Helper()
	actual, err := h.Caller.GetBalance(context.Background(), account)
	require.NoError(t, err, "failed to get token balance")
	assert.Equal(t, expected, actual.String(), "unexpected token balance")
}

func (h *TokenHarness) AssertAllowance(t *testing.T, owner, spender common.Address, expected string) {
	t.Helper()
	actual, err := h.Caller.GetAllowance(context.Background(), owner, spender)
	require.NoError(t, err, "failed to get token allowance")
	assert.Equal(t, expected, actual.String(), "unexpected token allowance")
}

func (h *TokenHarness) AssertUpgradeMaster(t *testing.T, expected common.Address) {
	t.Helper()
	actual, err := h.Caller.GetUpgradeMaster(context.Background())
	require.NoError(t, err, "failed to get upgrade master")
	assert.Equal(t, expected, actual, "unexpected upgrade master")
}

func (h *TokenHarness) Transfer(t *testing.T, sender *ecdsa.PrivateKey, recipient common.Address, amount string) {
	session := h.newSession(t, sender)
	tx, err := session.Transfer(recipient, bigString(t, amount))
	require.NoError(t, err)
	h.waitForTransaction(t, tx.Hash())
}

func (h *TokenHarness) Approve(t *testing.T, sender *ecdsa.PrivateKey, spender common.Address, amount string) {
	session := h.newSession(t, sender)
	tx, err := session.Approve(spender, bigString(t, amount))
	require.NoError(t, err)
	h.waitForTransaction(t, tx.Hash())
}

func (h *TokenHarness) newSession(t *testing.T, key *ecdsa.PrivateKey) *contract.CentrallyIssuedTokenSession {
	c, err := contract.NewCentrallyIssuedToken(h.ContractAddress, h.client)
	require.NoError(t, err)
	opts, err := bind.NewKeyedTransactorWithChainID(key, h.chainID)
	require.NoError(t, err)
	return &contract.CentrallyIssuedTokenSession{
		Contract:     c,
		TransactOpts: *opts,
	}
}

func (h *TokenHarness) waitForTransaction(t *testing.T, hash common.Hash) {
	_, err := eth.WaitForTransaction(context.Background(), h.client, hash, WaitProgress(t))
	require.NoError(t, err)
}

// deployTokenContract deploys the token contract with the multisig contract as the owner.
func deployTokenContract(t *testing.T, opts *bind.TransactOpts, client *ethclient.Client, multiSigContractAddress common.Address) common.Address {
	decimals := int64(8)
	waiter := eth.ProgressWaiter(client, WaitProgress(t))

	contractAddress, err := token.DeployContract(opts, client, multiSigContractAddress, "Test Token", "TEST", InitialSupply, decimals, waiter)
	require.NoError(t, err)
	return contractAddress
}

func bigString(t *testing.T, amount string) *big.Int {
	b, ok := new(big.Int).SetString(amount, 10)
	require.True(t, ok, "invalid big int: %s", amount)
	return b
}
