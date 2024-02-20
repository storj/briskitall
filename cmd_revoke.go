package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdRevoke struct {
	transactor    depMultiSigTransactor
	transactionID uint64
}

func (cmd *cmdRevoke) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.transactionID = uint64Arg(params, "TRANSACTIONID", "The ID of the pending transaction to revoke")
}

func (cmd *cmdRevoke) Execute(ctx context.Context) error {
	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	if err := transactor.RevokeConfirmation(ctx, cmd.transactionID); err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, client, cmd.transactor.client.nicknames, transactor.Caller, cmd.transactionID)
}
