package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitTokenApproveCmd(t *testing.T) {
	const amount = "100000000"

	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "token", "approve",
		test.AccountAddress[2],
		amount,
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.Token.AssertAllowance(t, harness.MultiSig.ContractAddress, test.AccountAddress[2], "0")
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.Token.AssertAllowance(t, harness.MultiSig.ContractAddress, test.AccountAddress[2], amount)
}
