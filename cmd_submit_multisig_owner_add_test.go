package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestSubmitMultiSigOwnerAddCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "submit", "multisig", "owner", "add",
		test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0],
	)
	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction 0 submitted to add owner %s", test.AccountAddress[2]),
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
