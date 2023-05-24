package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestDebugDeployMultiSigCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "debug", "deploy", "multisig",
		test.AccountAddress[2], test.AccountAddress[3],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	assert.Contains(t, stdout, "Contract Address: 0x564f2b59695a90e38164FC7411abAF6346553701\n")
}
