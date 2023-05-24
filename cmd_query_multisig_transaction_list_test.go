package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestQueryMultiSigTransactionListCmd(t *testing.T) {
	harness := test.Run(t)

	// Submit the transaction and assert that it shows up under pending.
	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	stdout := requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=false")
	assert.Empty(t, stdout)

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=true")
	assert.Empty(t, stdout)

	// pending with no status
	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=true", "--executed=false")
	assert.Equal(t, "0\n", stdout)

	// pending with status
	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		// FLAGS
		"--pending=true", "--executed=false", "--status")
	assert.Equal(t, fmt.Sprintf(`Transaction %d:
  Destination    = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Value          = 0
  Data (raw)     = 7065cb48000000000000000000000000d370bbc286487cd41e18c3561c0fc9c1a986516b
  Data (decoded) = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed       = false
  Confirmed      = false
  Confirmations(1):
    - 0x46f40B6B0dFDa35A8b6247489669a83c69804F3a
  Events(2):
    - Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a,0)
    - Submission(0)
`, transactionID), stdout)

	// Now confirm the transaction and assert that it shows up under executed.
	harness.MultiSig.ConfirmTransaction(t, test.AccountKey[1], transactionID)

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		// FLAGS
		"--pending=false", "--executed=false")
	assert.Empty(t, stdout)

	// executed with no status
	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		"--pending=false", "--executed=true")
	assert.Equal(t, "0\n", stdout)

	// executed with status
	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		// FLAGS
		"--pending=false", "--executed=true", "--status")
	assert.Equal(t, fmt.Sprintf(`Transaction %d:
  Destination    = 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe
  Value          = 0
  Data (raw)     = 7065cb48000000000000000000000000d370bbc286487cd41e18c3561c0fc9c1a986516b
  Data (decoded) = addOwner(0xD370Bbc286487CD41E18c3561c0Fc9C1a986516B)
  Executed       = true
  Confirmed      = true
  Confirmations(2):
    - 0x46f40B6B0dFDa35A8b6247489669a83c69804F3a
    - 0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0
  Events(4):
    - Confirmation(0x46f40B6B0dFDa35A8b6247489669a83c69804F3a,0)
    - Submission(0)
    - Confirmation(0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0,0)
    - Execution(0)
`, transactionID), stdout)

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "list",
		// FLAGS
		"--pending=true", "--executed=false")
	assert.Empty(t, stdout)
}
