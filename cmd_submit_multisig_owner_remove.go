package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitMultisigOwnerRemove struct {
	transactor depMultiSigTransactor
	owner      common.Address
}

func (cmd *cmdSubmitMultisigOwnerRemove) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.owner = addressArg(params, "OWNER", "Address of the owner to remove")
}

func (cmd *cmdSubmitMultisigOwnerRemove) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitRemoveOwner(ctx, cmd.owner)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to remove owner %s\n", transactionID, cmd.owner)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
