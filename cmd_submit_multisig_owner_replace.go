package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitMultisigOwnerReplace struct {
	transactor depMultiSigTransactor
	oldOwner   common.Address
	newOwner   common.Address
}

func (cmd *cmdSubmitMultisigOwnerReplace) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.oldOwner = addressArg(params, "OLDOWNER", "Address of the owner to be replaced")
	cmd.newOwner = addressArg(params, "NEWOWNER", "Address of the owner to replace with")
}

func (cmd *cmdSubmitMultisigOwnerReplace) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitReplaceOwner(ctx, cmd.oldOwner, cmd.newOwner)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to replace owner %s with %s\n", transactionID, cmd.oldOwner, cmd.newOwner)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
