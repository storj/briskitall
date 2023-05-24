package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/token"
)

type cmdDebugDeployToken struct {
	client      depClient
	sender      depSender
	owner       common.Address
	name        string
	symbol      string
	totalSupply *big.Int
	decimals    int64
}

func (cmd *cmdDebugDeployToken) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.sender.setup(params)
	cmd.name = stringFlag(params, "name", "Name of the token", "STORJ Token")
	cmd.symbol = stringFlag(params, "symbol", "Symbol of the token", "STORJ")
	cmd.totalSupply = bigIntFlag(params, "supply", "Total supply of the token", big.NewInt(10000000000))
	cmd.decimals = int64Flag(params, "decimals", "Number of decimals in the token", 8)
	cmd.owner = addressArg(params, "OWNER", "Address of the contract owner")
}

func (cmd *cmdDebugDeployToken) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}
	opts, err := cmd.sender.transactOpts(ctx)
	if err != nil {
		return err
	}

	contractAddress, err := token.DeployContract(opts, client, cmd.owner, cmd.name, cmd.symbol, cmd.totalSupply, cmd.decimals, waiter(ctx, client))
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx), "Contract Address:", contractAddress)
	return nil
}
