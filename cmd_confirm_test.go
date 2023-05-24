package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestConfirmCmd(t *testing.T) {
	harness := test.Run(t)

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)

	stdout := requireCmdSuccess(t, harness, "confirm",
		// ARGS
		fmt.Sprint(transactionID),
		// FLAGS
		"--sender-key-file", test.AccountKeyFile[1])
	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction %d confirmed by %s", transactionID, test.AccountAddress[1]),
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
