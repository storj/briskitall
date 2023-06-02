package main

import (
	"testing"

	"storj.io/briskitall/test"
)

func TestTestETHTransferCmd(t *testing.T) {
	// 1 ETH
	const amount = "1e18"

	harness := test.Run(t)

	harness.ETH.AssertBalance(t, test.AccountAddress[3], "0 Wei")

	requireCmdSuccess(t, harness, "test", "eth", "transfer",
		test.AccountAddress[3],
		amount,
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.ETH.AssertBalance(t, test.AccountAddress[3], "1 ETH")
}
