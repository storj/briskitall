package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitTokenTransfer struct {
	transactor depMultiSigTransactor
	token      depToken
	recipient  common.Address
	amount     *big.Int
}

func (cmd *cmdSubmitTokenTransfer) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.token.setup(params)
	cmd.recipient = addressArg(params, "RECIPIENT", "Address to receive token")
	cmd.amount = bigIntArg(params, "AMOUNT", "Amount to transfer")
}

func (cmd *cmdSubmitTokenTransfer) Execute(ctx context.Context) error {
	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitTokenTransfer(ctx, cmd.token.contractAddress, cmd.recipient, cmd.amount)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, client, transactor.Caller, transactionID)
}
