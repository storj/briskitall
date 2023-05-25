package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/eth"
)

type cmdQueryETHBalance struct {
	client   depClient
	multiSig depMultiSig
	address  *common.Address
}

func (cmd *cmdQueryETHBalance) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.multiSig.setup(params)
	cmd.address = optAddressArg(params, "ADDRESS", "Address to obtain ETH balance for (multisig contract if unset)")
}

func (cmd *cmdQueryETHBalance) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}

	address := cmd.multiSig.contractAddress
	if cmd.address != nil {
		address = *cmd.address
	}

	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx), eth.Pretty(balance))
	return nil
}
