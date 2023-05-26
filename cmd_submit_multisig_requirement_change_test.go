package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitMultiSigRequirementChangeCmd(t *testing.T) {
	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "multisig", "requirement", "change", 1,
		"--sender-key-file", test.AccountKeyFile[0])

	harness.MultiSig.AssertConfirmationRequirement(t, 2)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertConfirmationRequirement(t, 1)
}
