package test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/eth"
)

type ETHHarness struct {
	client eth.Client
}

func (h *ETHHarness) AssertBalance(t *testing.T, account common.Address, expected string) {
	t.Helper()
	actual, err := h.client.BalanceAt(context.Background(), account, nil)
	require.NoError(t, err, "failed to get ETH balance")
	assert.Equal(t, expected, eth.Pretty(actual), "unexpected ETH balance")
}
