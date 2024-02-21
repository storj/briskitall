package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitTokenApprove struct {
	transactor depMultiSigTransactor
	token      depToken
	spender    common.Address
	amount     *big.Int
}

func (cmd *cmdSubmitTokenApprove) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.token.setup(params)
	cmd.spender = addressArg(params, "SPENDER", "Spender address to approve")
	cmd.amount = bigIntArg(params, "AMOUNT", "Amount to approve")
}

func (cmd *cmdSubmitTokenApprove) Execute(ctx context.Context) error {
	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitTokenApprove(ctx, cmd.token.contractAddress, cmd.spender, cmd.amount)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, client, transactor.Caller, transactionID)
}
