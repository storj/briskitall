package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestSubmitETHTransferCmd(t *testing.T) {
	// 1 ETH
	const amount = "1e18"

	harness := test.Run(t)

	requireCmdSuccess(t, harness, "submit", "eth", "transfer",
		test.AccountAddress[3],
		amount,
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.ETH.AssertBalance(t, harness.MultiSig.ContractAddress, "10 ETH")
	harness.ETH.AssertBalance(t, test.AccountAddress[3], "0 Wei")

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], 0)

	harness.ETH.AssertBalance(t, harness.MultiSig.ContractAddress, "9 ETH")
	harness.ETH.AssertBalance(t, test.AccountAddress[3], "1 ETH")
}
