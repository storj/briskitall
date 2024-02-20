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
    - ETH[0]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0]: Submission()
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
    - ETH[0]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0]: Submission()
    - ETH[1]: Confirmation(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
    - ETH[1]: Execution()
`

	harness := test.Run(t)

	// Submit the transaction and assert that it shows up under pending.
	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	list := func(flags ...interface{}) string {
		args := []interface{}{"query", "multisig", "transaction", "list"}
		return requireCmdSuccess(t, harness, append(args, flags...)...)
	}

	assert.Empty(t, list("--pending=false", "--executed=false", "--short"))
	assert.Empty(t, list("--pending=false", "--executed=true", "--short"))
	assert.Equal(t, "0\n", list("--pending=true", "--executed=false", "--short"))
	assert.Equal(t, expectedPending, list("--pending=true", "--executed=false"))

	// Now confirm the transaction and assert that it shows up under executed.
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	assert.Empty(t, list("--pending=false", "--executed=false", "--short"))
	assert.Empty(t, list("--pending=true", "--executed=false", "--short"))
	assert.Equal(t, "0\n", list("--pending=false", "--executed=true", "--short"))
	assert.Equal(t, expectedExecuted, list("--pending=false", "--executed=true"))
}
