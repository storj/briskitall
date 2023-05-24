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
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	if _, err := transactor.ConfirmTransaction(ctx, cmd.transactionID); err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d confirmed by %s\n", cmd.transactionID, transactor.Sender())
	return nil
}
