package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitMultiSigOwnerAddCmd(t *testing.T) {
	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "multisig", "owner", "add",
		test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
