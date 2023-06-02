package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultiSigTransactionListCmd(t *testing.T) {
	const expectedPending = `Transaction 0:
  Destination = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Call        = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed    = false
  Confirmed   = false
  Confirmations(1):
    - Owner(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
  Events(2):
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Submission()
`
	const expectedExecuted = `Transaction 0:
  Destination = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Call        = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed    = true
  Confirmed   = true
  Confirmations(2):
    - Owner(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - Owner(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
  Events(4):
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0x43e7062198defbef1d15aaf5af8b952d2a0442c7f90ac59efe668cdfef6a8694]: Submission()
    - ETH[0xcf7cd4509886c47f9e0fb416ee996e45190d75fb5fccb3561d8bd60c5554f2ce]: Confirmation(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
    - ETH[0xcf7cd4509886c47f9e0fb416ee996e45190d75fb5fccb3561d8bd60c5554f2ce]: Execution()
`

	harness := test.Run(t)

	// Submit the transaction and assert that it shows up under pending.
	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	assert.Empty(t, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=false"))

	assert.Empty(t, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=true"))

	assert.Equal(t, "0\n", requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=true", "--executed=false"))

	assert.Equal(t, expectedPending, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=true", "--executed=false", "--status"))

	// Now confirm the transaction and assert that it shows up under executed.
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	assert.Empty(t, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=false"))

	assert.Empty(t, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=true", "--executed=false"))

	assert.Equal(t, "0\n", requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=true"))

	assert.Equal(t, expectedExecuted, requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=true", "--status"))
}
