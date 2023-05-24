package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestSubmitTokenTransferFromCmd(t *testing.T) {
	const amount = "100000000"

	harness := test.Run(t)

	// Transfer amount from multisig to account 2 to stage the transfer from
	transferID := harness.MultiSig.SubmitTransfer(t, test.AccountKey[0], test.AccountAddress[2], amount)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transferID)

	// Account 2 approves the multisig account to transfer amount
	harness.Token.Approve(t, test.AccountKey[2], harness.MultiSig.ContractAddress, amount)

	transferFromID := transferID + 1

	// Submits the transferFrom for amount from account 2 to account 3
	stdout := requireCmdSuccess(t, harness, "submit", "token", "transfer-from",
		test.AccountAddress[2], test.AccountAddress[3], amount,
		"--sender-key-file", test.AccountKeyFile[0],
	)
	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction %d submitted to transfer 100000000 from %s to %s",
			transferFromID,
			test.AccountAddress[2],
			test.AccountAddress[3],
		),
	)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "99900000000")
	harness.Token.AssertBalance(t, test.AccountAddress[2], amount)
	harness.Token.AssertBalance(t, test.AccountAddress[3], "0")

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transferFromID)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "99900000000")
	harness.Token.AssertBalance(t, test.AccountAddress[2], "0")
	harness.Token.AssertBalance(t, test.AccountAddress[3], amount)
}
