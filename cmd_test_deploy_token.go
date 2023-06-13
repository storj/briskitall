package main

import (
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/token"
)

type cmdTestDeployToken struct {
	client      depClient
	sender      depSender
	output      depOutput
	owner       common.Address
	name        string
	symbol      string
	totalSupply *big.Int
	decimals    int64
}

func (cmd *cmdTestDeployToken) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.sender.setup(params)
	cmd.output.setup(params)
	cmd.name = stringFlag(params, "name", "Name of the token", "Test Token")
	cmd.symbol = stringFlag(params, "symbol", "Symbol of the token", "TEST")
	cmd.totalSupply = bigIntFlag(params, "supply", "Total supply of the token", big.NewInt(10000000000))
	cmd.decimals = int64Flag(params, "decimals", "Number of decimals in the token", 8)
	cmd.owner = addressArg(params, "OWNER", "Address of the contract owner")
}

func (cmd *cmdTestDeployToken) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}
	opts, err := cmd.sender.transactOpts(ctx, client)
	if err != nil {
		return err
	}

	contractAddress, err := token.DeployContract(opts, client, cmd.owner, cmd.name, cmd.symbol, cmd.totalSupply, cmd.decimals, waiter(ctx, client))
	if err != nil {
		return err
	}

	cmd.output.out(ctx, outTestDeployToken{ContractAddress: contractAddress})
	return nil
}

type outTestDeployToken struct {
	ContractAddress common.Address
}

func (out outTestDeployToken) TextOut(w io.Writer) {
	fmt.Fprintln(w, "Token contract address:", out.ContractAddress)
}
