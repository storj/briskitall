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
	assert.Equal(t, fmt.Sprintf(`Transaction %d:
  Destination = %s
  Call        = addOwner(%s)
  Executed    = false
  Confirmed   = false
  Confirmations(1):
    - Owner(%s)
  Events(2):
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Confirmation(%s)
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Submission()
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2],
		test.AccountAddress[0],
		test.AccountAddress[0],
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
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Confirmation(%s)
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Submission()
    - ETH[0xcf7cd4509886c47f9e0fb416ee996e45190d75fb5fccb3561d8bd60c5554f2ce]: Confirmation(%s)
    - ETH[0xcf7cd4509886c47f9e0fb416ee996e45190d75fb5fccb3561d8bd60c5554f2ce]: Execution()
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2],
		test.AccountAddress[0],
		test.AccountAddress[1],
		test.AccountAddress[0],
		test.AccountAddress[1],
	), stdout)
}
