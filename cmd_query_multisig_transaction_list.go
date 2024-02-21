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
	cmd.status = !boolFlag(params, "short", "Only show transaction numbers")
}

func (cmd *cmdQueryMultiSigTransactionList) Execute(ctx context.Context) error {
	client, caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}

	transactionIDs, err := caller.ListTransactionIDs(ctx, cmd.pending, cmd.executed)
	if err != nil {
		return err
	}

	needsSeparator := false

	for _, transactionID := range transactionIDs {
		confirmations, err := caller.GetTransactionConfirmations(ctx, transactionID)
		if err != nil {
			return err
		}
		if len(confirmations) == 0 {
			continue
		}

		if !cmd.status {
			fmt.Fprintln(clingy.Stdout(ctx), transactionID)
			continue
		}

		if needsSeparator {
			fmt.Fprintln(clingy.Stdout(ctx))
		}
		if err := printTransactionStatus(ctx, client, caller, transactionID); err != nil {
			return err
		}
		needsSeparator = true
	}
	return nil
}
