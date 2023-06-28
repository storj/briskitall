package main

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/test"
)

func requireCmdSuccess(t *testing.T, harness *test.Harness, rawArgs ...interface{}) string {
	stdout, stderr, err := runCmd(t, harness, rawArgs...)
	require.NoError(t, err)
	assert.Empty(t, stderr, "stderr should be empty")

	return stdout
}

func runCmd(t *testing.T, harness *test.Harness, rawArgs ...interface{}) (string, string, error) {
	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)

	args := make([]string, len(rawArgs))
	for i, arg := range rawArgs {
		args[i] = fmt.Sprint(arg)
	}

	env := clingy.Environment{
		Stdout: stdoutBuf,
		Stderr: stderrBuf,
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
			case envSkipConfirmation:
				return "true"
			}
			return ""
		},
	}

	ok, err := env.Run(context.Background(), commands)
	require.True(t, ok, "command dispatch failed: %s", stdoutBuf.String())

	stdout := replaceETHTX(stdoutBuf.String())
	stderr := stderrBuf.String()

	t.Logf("stdout=%s", stdout)
	t.Logf("stderr=%s", stderr)

	return stdout, stderr, err
}

var ethTXRe = regexp.MustCompile(`ETH\[(0x[[:xdigit:]]{64})\]`)

func replaceETHTX(s string) string {
	subs := make(map[string]string)
	return ethTXRe.ReplaceAllStringFunc(s, func(m string) string {
		sub, ok := subs[m]
		if !ok {
			sub = fmt.Sprintf("ETH[%d]", len(subs))
			subs[m] = sub
		}
		return sub
	})
}
