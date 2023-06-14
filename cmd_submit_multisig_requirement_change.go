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
	transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitChangeRequirement(ctx, cmd.requirement)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
