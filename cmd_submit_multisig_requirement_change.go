package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdSubmitMultisigRequirementChange struct {
	transactor  depMultiSigTransactor
	requirement uint64
}

func (cmd *cmdSubmitMultisigRequirementChange) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.requirement = uint64Arg(params, "REQUIREMENT", "New number of confirmations to require")
}

func (cmd *cmdSubmitMultisigRequirementChange) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitChangeRequirement(ctx, cmd.requirement)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to change confirmation requirement to %d\n", transactionID, cmd.requirement)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
