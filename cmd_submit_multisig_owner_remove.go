package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitMultisigOwnerRemove struct {
	transactor depMultiSigTransactor
	owner      common.Address
}

func (cmd *cmdSubmitMultisigOwnerRemove) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.owner = addressArg(params, "OWNER", "Address of the owner to remove")
}

func (cmd *cmdSubmitMultisigOwnerRemove) Execute(ctx context.Context) error {
	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitRemoveOwner(ctx, cmd.owner)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, client, cmd.transactor.client.nicknames, transactor.Caller, transactionID)
}
