package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdQueryTokenBalance struct {
	caller  depTokenCaller
	address common.Address
}

func (cmd *cmdQueryTokenBalance) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
	cmd.address = addressArg(params, "ADDRESS", "Address to obtain token balance for")
}

func (cmd *cmdQueryTokenBalance) Execute(ctx context.Context) error {
	caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	balance, err := caller.GetBalance(ctx, cmd.address)
	if err != nil {
		return err
	}
	fmt.Fprintln(clingy.Stdout(ctx), balance)
	return nil
}
