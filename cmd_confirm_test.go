package main

import (
	"fmt"
	"testing"

	"storj.io/briskitall/test"
)

func TestConfirmCmd(t *testing.T) {
	harness := test.Run(t)

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)

	requireCmdSuccess(t, harness, "confirm",
		// ARGS
		fmt.Sprint(transactionID),
		// FLAGS
		"--sender-key-file", test.AccountKeyFile[1])

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
