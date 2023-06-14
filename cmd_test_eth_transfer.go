package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
)

type cmdTestETHTransfer struct {
	client    depClient
	sender    depSender
	recipient common.Address
	amount    *big.Int
}

func (cmd *cmdTestETHTransfer) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.sender.setup(params)
	cmd.recipient = addressArg(params, "RECIPIENT", "Address to receive ETH")
	cmd.amount = bigIntArg(params, "AMOUNT", "Amount to transfer (in WEI)")
}

func (cmd *cmdTestETHTransfer) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}
	opts, done, err := cmd.sender.transactOpts(ctx, client)
	if err != nil {
		return err
	}
	defer done()

	if err := eth.Transfer(opts, client, cmd.recipient, cmd.amount, waiter(ctx, client)); err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx), "Transferred.")
	return nil
}
