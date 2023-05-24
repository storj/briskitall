package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitTokenSetUpgradeMaster struct {
	transactor depMultiSigTransactor
	token      depToken
	master     common.Address
}

func (cmd *cmdSubmitTokenSetUpgradeMaster) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.token.setup(params)
	cmd.master = addressArg(params, "MASTER", "The new upgrade master")
}

func (cmd *cmdSubmitTokenSetUpgradeMaster) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitSetUpgradeMaster(ctx, cmd.token.contractAddress, cmd.master)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to set the upgrade master to %s\n", transactionID, cmd.master)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
