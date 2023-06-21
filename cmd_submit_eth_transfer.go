package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdSubmitETHTransfer struct {
	transactor depMultiSigTransactor
	recipient  common.Address
	amount     *big.Int
}

func (cmd *cmdSubmitETHTransfer) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.recipient = addressArg(params, "RECIPIENT", "Address to receive ETH")
	cmd.amount = ethUnitArg(params, "AMOUNT", "Amount to transfer (in ETH/GWEI/WEI)")
}

func (cmd *cmdSubmitETHTransfer) Execute(ctx context.Context) error {
	transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitETHTransfer(ctx, cmd.recipient, cmd.amount)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
