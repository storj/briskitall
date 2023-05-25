package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/zeebo/clingy"
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

	fmt.Fprintln(clingy.Stdout(ctx), prettyETH(balance))
	return nil
}

// prettyETH is lovingly borrowed from storj/crypto-batch-payment
func prettyETH(wei *big.Int) string {
	switch {
	case wei.Cmp(big.NewInt(1_000_000_000_000_000)) > 0:
		return fmt.Sprintf("%s ETH", decimal.NewFromBigInt(wei, -18))
	case wei.Cmp(big.NewInt(1_000_000_0)) > 0:
		return fmt.Sprintf("%s GWei", decimal.NewFromBigInt(wei, -9))
	default:
		return fmt.Sprintf("%s Wei", wei)
	}
}
