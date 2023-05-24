package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/test"
)

func TestSubmitCallCmd(t *testing.T) {
	harness := test.Run(t)

	// Dump the ABI to disk
	abiPath := filepath.Join(t.TempDir(), "abi.json")
	require.NoError(t, os.WriteFile(abiPath, []byte(contract.MultiSigWalletWithDailyLimitMetaData.ABI), 0644))

	stdout := requireCmdSuccess(t, harness, "submit", "call",
		harness.MultiSig.ContractAddress.String(),
		abiPath,
		"addOwner",
		test.AccountAddress[2].String(),
		"--sender-key-file", test.AccountKeyFile[0],
	)
	assert.Contains(t, stdout,
		fmt.Sprintf("Transaction 0 submitted to call %s/addOwner", harness.MultiSig.ContractAddress),
	)

	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], false)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)
	harness.MultiSig.AssertIsOwner(t, test.AccountAddress[2], true)
}
