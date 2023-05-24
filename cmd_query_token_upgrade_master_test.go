package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryTokenUpgradeMasterCmd(t *testing.T) {
	harness := test.Run(t)

	stdout := requireCmdSuccess(t, harness, "query", "token", "upgrade-master")
	assert.Equal(t, stdout, fmt.Sprintln(harness.MultiSig.ContractAddress))
}
