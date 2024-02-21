package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultiSigTransactionListCmd(t *testing.T) {
	today := time.Now().Format("2006-01-02")

	expectedPending := fmt.Sprintf(`Transaction 0:
  Destination = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Call        = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed    = false
  Confirmed   = false
  Confirmations(1):
    - Owner(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
  Events(2):
    - %s: ETH[0]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - %s: ETH[0]: Submission()
`, today, today)

	expectedExecuted := fmt.Sprintf(`Transaction 0:
  Destination = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Call        = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed    = true
  Confirmed   = true
  Confirmations(2):
    - Owner(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - Owner(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
  Events(4):
    - %s: ETH[0]: Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a)
    - %s: ETH[0]: Submission()
    - %s: ETH[1]: Confirmation(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0)
    - %s: ETH[1]: Execution()
`, today, today, today, today)

	harness := test.Run(t)

	// Submit the transaction and assert that it shows up under pending.
	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	list := func(flags ...interface{}) string {
		args := []interface{}{"query", "multisig", "transaction", "list"}
		return requireCmdSuccess(t, harness, append(args, flags...)...)
	}

	assert.Empty(t, list("--pending=false", "--executed=false"))
	assert.Empty(t, list("--pending=false", "--executed=true"))
	assert.Equal(t, "0\n", list("--pending=true", "--executed=false"))
	assert.Equal(t, expectedPending, list("--pending=true", "--executed=false", "--status"))

	// Now confirm the transaction and assert that it shows up under executed.
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	assert.Empty(t, list("--pending=false", "--executed=false"))
	assert.Empty(t, list("--pending=true", "--executed=false"))
	assert.Equal(t, "0\n", list("--pending=false", "--executed=true"))
	assert.Equal(t, expectedExecuted, list("--pending=false", "--executed=true", "--status"))
}
