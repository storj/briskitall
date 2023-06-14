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
	transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	if err := transactor.ExecuteTransaction(ctx, cmd.transactionID); err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, transactor.Caller, cmd.transactionID)
}
