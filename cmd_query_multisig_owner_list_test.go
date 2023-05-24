package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultisigOwnerListCmd(t *testing.T) {
	harness := test.Run(t)
	stdout := requireCmdSuccess(t, harness, "query", "multisig", "owner", "list")
	assert.Equal(t,
		fmt.Sprintf("%s\n%s\n", test.AccountAddress[0], test.AccountAddress[1]),
		stdout)
}
