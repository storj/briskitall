package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdQueryMultiSigTransactionList struct {
	caller           depMultiSigCaller
	skipZeroConfirms bool
	tail             int
	pending          bool
	executed         bool
	status           bool
}

func (cmd *cmdQueryMultiSigTransactionList) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
	cmd.skipZeroConfirms = toggleFlag(params, "skip-zero-confirms", "If true, don't show transactions with zero confirmations.", true)
	cmd.tail = optionalIntFlag(params, "tail", "If > 0, the maximum number of transactions to list, where the earlier transactions are skipped", 0)
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

	var transactionIDsToShow []uint64
	for i := len(transactionIDs) - 1; i >= 0; i-- {
		transactionID := transactionIDs[i]
		if cmd.skipZeroConfirms {
			confirmations, err := caller.GetTransactionConfirmations(ctx, transactionID)
			if err != nil {
				return err
			}
			if len(confirmations) == 0 {
				continue
			}
		}
		transactionIDsToShow = append(transactionIDsToShow, transactionID)
		if cmd.tail > 0 && len(transactionIDsToShow) >= cmd.tail {
			break
		}
	}

	for i := len(transactionIDsToShow) - 1; i >= 0; i-- {
		transactionID := transactionIDsToShow[i]

		if !cmd.status {
			fmt.Fprintln(clingy.Stdout(ctx), transactionID)
			continue
		}

		if needsSeparator {
			fmt.Fprintln(clingy.Stdout(ctx))
		}
		if err := printTransactionStatus(ctx, client, cmd.caller.client.nicknames, caller, transactionID); err != nil {
			return err
		}
		needsSeparator = true
	}
	return nil
}
