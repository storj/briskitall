package main

import (
	"fmt"
	"testing"

	"storj.io/briskitall/test"
)

func TestExecuteCmd(t *testing.T) {
	amount := test.InitialSupply.String()
	harness := test.Run(t)

	// First transfer everything out of the multisig account to account 2
	txid1 := harness.MultiSig.SubmitTransfer(t, test.AccountKey[0], test.AccountAddress[2], amount)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], txid1)

	// Now try and transfer more to account 3
	txid2 := harness.MultiSig.SubmitTransfer(t, test.AccountKey[0], test.AccountAddress[3], amount)
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], txid2)

	// This transaction will have failed execution
	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "0")
	harness.Token.AssertBalance(t, test.AccountAddress[2], amount)
	harness.Token.AssertBalance(t, test.AccountAddress[3], "0")

	// Now transfer back to the multisig contract
	harness.Token.Transfer(t, test.AccountKey[2], harness.MultiSig.ContractAddress, amount)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, amount)
	harness.Token.AssertBalance(t, test.AccountAddress[2], "0")
	harness.Token.AssertBalance(t, test.AccountAddress[3], "0")

	// Now execute the transaction to transfer to account 3 again
	requireCmdSuccess(t, harness, "execute", fmt.Sprint(txid2),
		"--sender-key-file", test.AccountKeyFile[0],
	)

	harness.Token.AssertBalance(t, harness.MultiSig.ContractAddress, "0")
	harness.Token.AssertBalance(t, test.AccountAddress[2], "0")
	harness.Token.AssertBalance(t, test.AccountAddress[3], amount)
}
