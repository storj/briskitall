package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitMultisigOwnerRemoveCmd(t *testing.T) {
	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "multisig", "owner", "remove",
		test.AccountAddress[1],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], true)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], false)
}
