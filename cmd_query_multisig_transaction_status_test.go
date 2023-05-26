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
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Confirmation(%s)
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Submission()
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
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Confirmation(%s)
    - ETH[0xb178bdc2bab8511dec202c1536b9211ef04aa10c5f2230319b6f049e920a2d93]: Submission()
    - ETH[0xa133a7cf683c9b7401c9614fb553b9914deb9ded49224d963ee7bd3e0a946325]: Confirmation(%s)
    - ETH[0xa133a7cf683c9b7401c9614fb553b9914deb9ded49224d963ee7bd3e0a946325]: Execution()
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2],
		test.AccountAddress[0],
		test.AccountAddress[1],
		test.AccountAddress[0],
		test.AccountAddress[1],
	), stdout)
}
