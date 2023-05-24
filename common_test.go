package main

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/test"
)

func requireCmdSuccess(t *testing.T, harness *test.Harness, rawArgs ...interface{}) string {
	stdout, stderr, err := runCmd(t, harness, rawArgs...)
	require.NoError(t, err)
	assert.Empty(t, stderr.String())
	return stdout.String()
}

func runCmd(t *testing.T, harness *test.Harness, rawArgs ...interface{}) (*bytes.Buffer, *bytes.Buffer, error) {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	args := make([]string, len(rawArgs))
	for i, arg := range rawArgs {
		args[i] = fmt.Sprint(arg)
	}

	env := clingy.Environment{
		Stdout: stdout,
		Stderr: stderr,
		Name:   "Test",
		Args:   args,
		Getenv: func(key string) string {
			switch key {
			case envChainID:
				return harness.ChainID.String()
			case envNodeURL:
				return harness.NodeURL
			case envMultiSigContractAddress:
				return harness.MultiSig.ContractAddress.String()
			case envTokenContractAddress:
				return harness.Token.ContractAddress.String()
			}
			return ""
		},
	}

	ok, err := env.Run(context.Background(), commands)
	require.True(t, ok, "command dispatch failed: %s", stdout.String())

	return stdout, stderr, err
}
