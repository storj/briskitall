package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestSubmitMultisigOwnerRemoveCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "submit", "multisig", "owner", "remove",
		test.AccountAddress[1],
		"--sender-key-file", test.AccountKeyFile[0],
	)
	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction 0 submitted to remove owner %s", test.AccountAddress[1]),
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], true)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[1], false)
}
