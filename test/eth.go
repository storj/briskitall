package test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"storj.io/briskitall/internal/eth"
)

type ETHHarness struct {
	client *ethclient.Client
}

func (h *ETHHarness) AssertBalance(t *testing.T, account common.Address, expected string) {
	t.Helper()
	actual, err := h.client.BalanceAt(context.Background(), account, nil)
	require.NoError(t, err, "failed to get ETH balance")
	assert.Equal(t, expected, eth.Pretty(actual), "unexpected ETH balance")
}

//func (h *ETHHarness) Transfer(t *testing.T, sender *ecdsa.PrivateKey, recipient common.Address, amount string) {
//	tx, err := h.client.SendTransaction(recipient, bigString(t, amount))
//	require.NoError(t, err)
//	h.waitForTransaction(t, tx.Hash())
//}
