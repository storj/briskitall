package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitTokenTransferFrom struct {
	transactor depMultiSigTransactor
	token      depToken
	from       common.Address
	to         common.Address
	amount     *big.Int
}

func (cmd *cmdSubmitTokenTransferFrom) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.token.setup(params)
	cmd.from = addressArg(params, "FROM", "Address to transfer from")
	cmd.to = addressArg(params, "TO", "Address to transfer to")
	cmd.amount = bigIntArg(params, "AMOUNT", "Amount to transfer")
}

func (cmd *cmdSubmitTokenTransferFrom) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitTransferFrom(ctx, cmd.token.contractAddress, cmd.from, cmd.to, cmd.amount)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to transfer %s from %s to %s\n", transactionID, cmd.amount, cmd.from, cmd.to)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
