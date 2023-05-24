package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdQueryMultiSigTransactionList struct {
	caller   depMultiSigCaller
	pending  bool
	executed bool
	status   bool
}

func (cmd *cmdQueryMultiSigTransactionList) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
	cmd.pending = toggleFlag(params, "pending", "List pending transactions", true)
	cmd.executed = toggleFlag(params, "executed", "List executed transactions", false)
	cmd.status = boolFlag(params, "status", "Show the transaction status", false)
}

func (cmd *cmdQueryMultiSigTransactionList) Execute(ctx context.Context) error {
	caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}

	transactionIDs, err := caller.ListTransactionIDs(ctx, cmd.pending, cmd.executed)
	if err != nil {
		return err
	}

	for _, transactionID := range transactionIDs {
		if cmd.status {
			if err := printTransactionStatus(ctx, caller, transactionID); err != nil {
				return err
			}
		} else {
			fmt.Fprintln(clingy.Stdout(ctx), transactionID)
		}
	}
	return nil
}
