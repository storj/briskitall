package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdConfirm struct {
	transactor    depMultiSigTransactor
	transactionID uint64
}

func (cmd *cmdConfirm) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.transactionID = uint64Arg(params, "TRANSACTIONID", "The ID of the pending transaction to confirm")
}

func (cmd *cmdConfirm) Execute(ctx context.Context) error {
	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	if _, err := transactor.ConfirmTransaction(ctx, cmd.transactionID); err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	printTransactionStatus(ctx, client, transactor.Caller, cmd.transactionID)
	return nil
}
