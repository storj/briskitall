package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitTokenTransferCmd(t *testing.T) {
	const amount = "100000000"

	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "token", "transfer",
		test.AccountAddress[2],
		amount,
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "100000000000")
	harness.Token.AssertBalance(t, test.AccountAddress[2], "0")

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "99900000000")
	harness.Token.AssertBalance(t, test.AccountAddress[2], amount)
}
