package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultiSigTransactionStatusCmd(t *testing.T) {
	harness := test.Run(t)
	today := time.Now().Format("2006-01-02")

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	stdout := requireCmdSuccess(t, harness, "query", "multisig", "transaction", "status", "0")
	assert.Equal(t, fmt.Sprintf(`Transaction %d:
  Destination = %s
  Call        = addOwner(%s)
  Executed    = false
  Confirmed   = false
  Confirmations(1):
    - Owner(%s)
  Events(2):
    - %s: ETH[0]: Confirmation(%s)
    - %s: ETH[0]: Submission()
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2],
		test.AccountAddress[0],
		today,
		test.AccountAddress[0],
		today,
	), stdout)

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "status", "0")
	assert.Equal(t, fmt.Sprintf(`Transaction %d:
  Destination = %s
  Call        = addOwner(%s)
  Executed    = true
  Confirmed   = true
  Confirmations(2):
    - Owner(%s)
    - Owner(%s)
  Events(4):
    - %s: ETH[0]: Confirmation(%s)
    - %s: ETH[0]: Submission()
    - %s: ETH[1]: Confirmation(%s)
    - %s: ETH[1]: Execution()
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2],
		test.AccountAddress[0],
		test.AccountAddress[1],
		today,
		test.AccountAddress[0],
		today,
		today,
		test.AccountAddress[1],
		today,
	), stdout)
}
