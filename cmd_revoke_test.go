package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/briskitall/test"
)

func TestRevokeCmd(t *testing.T) {
	harness := test.Run(t)

	transactionID := harness.MultiSig.SubmitAddOwner(t, test.AccountKey[0], test.AccountAddress[2])

	stdout := requireCmdSuccess(t, harness, "revoke", fmt.Sprint(transactionID),
		"--sender-key-file", test.AccountKeyFile[0])
	assert.Contains(t, stdout, fmt.Sprintf("Transaction %d revoked by %s", transactionID, test.AccountAddress[0]))

	stdout = requireCmdSuccess(t, harness, "query", "multisig", "transaction", "status", "0")
	assert.Equal(t, stdout, fmt.Sprintf(`Transaction %d:
  Destination    = %s
  Value          = 0
  Data (raw)     = 7065cb48000000000000000000000000%x
  Data (decoded) = addOwner(%s)
  Executed       = false
  Confirmed      = false
  Confirmations(0):
  Events(3):
    - Confirmation(%s,0)
    - Submission(0)
    - Revocation(%s,0)
`, transactionID,
		harness.MultiSig.ContractAddress,
		test.AccountAddress[2], test.AccountAddress[2],
		test.AccountAddress[0],
		test.AccountAddress[0]))
}
