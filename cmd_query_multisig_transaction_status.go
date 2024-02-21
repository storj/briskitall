package main

import (
	"context"

	"github.com/zeebo/clingy"
)

type cmdQueryMultiSigTransactionStatus struct {
	caller        depMultiSigCaller
	transactionID uint64
}

func (cmd *cmdQueryMultiSigTransactionStatus) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
	cmd.transactionID = uint64Arg(params, "TRANSACTIONID", "The ID of the transaction")
}

func (cmd *cmdQueryMultiSigTransactionStatus) Execute(ctx context.Context) error {
	client, caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	return printTransactionStatus(ctx, client, cmd.caller.client.nicknames, caller, cmd.transactionID)
}
