package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestTestDeployTokenCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "test", "deploy", "token",
		test.AccountAddress[2],
		"--sender-key-file", test.AccountKeyFile[0],
	)

	assert.Contains(t, stdout, "Token contract address: 0x564f2b59695a90e38164FC7411abAF6346553701\n")
}
