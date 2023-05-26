package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitTokenSetUpgradeMasterCmd(t *testing.T) {
	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "token", "set-upgrade-master",
		test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.Token.AssertUpgradeMaster(t, harness.MultiSig.ContractAddress)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.Token.AssertUpgradeMaster(t, test.AccountAddress[2])
}
