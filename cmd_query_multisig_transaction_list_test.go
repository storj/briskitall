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
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Submission()
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
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Submission()
    - ETH[0xa133a7cf683c9b7401c9614fb553b9914deb9ded49224d963ee7bd3e0a946325]: Confirmation(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
    - ETH[0xa133a7cf683c9b7401c9614fb553b9914deb9ded49224d963ee7bd3e0a946325]: Execution()
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
