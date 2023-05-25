package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/eth"
)

type cmdSubmitETHTransfer struct {
	transactor depMultiSigTransactor
	recipient  common.Address
	amount     *big.Int
}

func (cmd *cmdSubmitETHTransfer) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.recipient = addressArg(params, "RECIPIENT", "Address to receive ETH")
	cmd.amount = bigIntArg(params, "AMOUNT", "Amount to transfer (in WEI)")
}

func (cmd *cmdSubmitETHTransfer) Execute(ctx context.Context) error {
	transactor, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	transactionID, err := transactor.SubmitETHTransfer(ctx, cmd.recipient, cmd.amount)
	if err != nil {
		return err
	}
	fmt.Fprintf(clingy.Stdout(ctx), "Transaction %d submitted to transfer %s to %s\n", transactionID, eth.Pretty(cmd.amount), cmd.recipient)
	return printTransactionStatus(ctx, transactor.Caller, transactionID)
}
