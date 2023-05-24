package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdExecute struct {
	transactor    depMultiSigTransactor
	transactionID uint64
}

func (cmd *cmdExecute) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.transactionID = uint64Arg(params, "TRANSACTIONID", "The ID of the confirmed transaction to (re)execute")
}

func (cmd *cmdExecute) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	if err := transactor.ExecuteTransaction(ctx, cmd.transactionID); err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d executed by %s\n", cmd.transactionID, transactor.Sender())
	return printTransactionStatus(ctx, transactor.Caller, cmd.transactionID)
}
