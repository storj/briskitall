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
	transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitTokenTransferFrom(ctx, cmd.token.contractAddress, cmd.from, cmd.to, cmd.amount)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
