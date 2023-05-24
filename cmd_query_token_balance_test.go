package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryTokenBalanceCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "query", "token", "balance", harness.MultiSig.ContractAddress)
	assert.Equal(t, stdout, "100000000000\n")
}
