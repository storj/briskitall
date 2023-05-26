package main

import (
	"fmt"
	"testing"

	"storj.io/briskitall/test"
)

func TestRevokeCmd(t *testing.T) {
	harness := test.Run(t)

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	requireCmdSuccess(t, harness, "revoke", fmt.Sprint(transactionID),
		"--sender-key-file", test.AccountKeyFile[0])

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	// Should not be an owner since account 0 revoked before it was confirmed
	// by account 1.
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)
}
