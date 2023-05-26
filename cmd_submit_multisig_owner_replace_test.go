package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitMultiSigOwnerReplaceCmd(t *testing.T) {
	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "multisig", "owner", "replace",
		test.AccountAddress[1], test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0])

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], true)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], false)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
