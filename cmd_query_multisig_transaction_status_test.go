package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultiSigTransactionStatusCmd(t *testing.T) {
	harness := test.Run(t)

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	stdout := requireCmdSuccess(t, harness, "query", "multisig", "transaction", "status", "0")
	assert.Equal(t, stdout, fmt.Sprintf(`Transaction %d:
  Destination    = %s
  Value          = 0
  Data (raw)     = 7065cb48000000000000000000000000%x
  Data (decoded) = addOwner(%s)
  Executed       = false
  Confirmed      = false
  Confirmations(1):
    - %s
  Events(2):
    - Confirmation(%s,0)
    - Submission(0)
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2], test.AccountAddress[2],
		test.AccountAddress[0], test.AccountAddress[0]))

	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "status", "0")
	assert.Equal(t, stdout, fmt.Sprintf(`Transaction %d:
  Destination    = %s
  Value          = 0
  Data (raw)     = 7065cb48000000000000000000000000%x
  Data (decoded) = addOwner(%s)
  Executed       = true
  Confirmed      = true
  Confirmations(2):
    - %s
    - %s
  Events(4):
    - Confirmation(%s,0)
    - Submission(0)
    - Confirmation(%s,0)
    - Execution(0)
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2], test.AccountAddress[2],
		test.AccountAddress[0], test.AccountAddress[1],
		test.AccountAddress[0], test.AccountAddress[1]))
}
