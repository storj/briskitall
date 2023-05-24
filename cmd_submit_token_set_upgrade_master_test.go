package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestSubmitTokenSetUpgradeMasterCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "submit", "token", "set-upgrade-master",
		test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction 0 submitted to set the upgrade master to %s",
			test.AccountAddress[2],
		),
	)

	harness.Token.AssertUpgradeMaster(t, harness.MultiSig.ContractAddress)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.Token.AssertUpgradeMaster(t, test.AccountAddress[2])
}
