package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
)

type cmdQueryETHBalance struct {
	client  depClient
	address common.Address
}

func (cmd *cmdQueryETHBalance) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.address = addressArg(params, "ADDRESS", "Address to obtain ETH balance for")
}

func (cmd *cmdQueryETHBalance) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(ctx, cmd.address, nil)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx), eth.Pretty(balance))
	return nil
}
