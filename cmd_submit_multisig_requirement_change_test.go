package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestSubmitMultiSigRequirementChangeCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "submit", "multisig", "requirement", "change", 1,
		"--sender-key-file", test.AccountKeyFile[0])

	assert.Contains(t, stdout, "Transaction 0 submitted to change confirmation requirement to 1")

	harness.MultiSig.AssertConfirmationRequirement(t, 2)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertConfirmationRequirement(t, 1)
}
