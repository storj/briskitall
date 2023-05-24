package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryTokenAllowanceCmd(t *testing.T) {
	const amount = "100000000"

	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "query", "token", "allowance",
		harness.MultiSig.ContractAddress, test.AccountAddress[1])
	assert.Equal(t, stdout, "0\n")

	// Approve transfer of amount from multisig to account 1
	approveID := harness.MultiSig.SubmitApprove(t, test.AccountKey[0], test.AccountAddress[1], amount)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], approveID)

	stdout = requireCmdSuccess(t, harness, "query", "token", "allowance",
		harness.MultiSig.ContractAddress, test.AccountAddress[1])
	assert.Equal(t, stdout, amount+"\n")
}
