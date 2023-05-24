package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type cmdQueryTokenAllowance struct {
	caller  depTokenCaller
	owner   common.Address
	spender common.Address
}

func (cmd *cmdQueryTokenAllowance) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
	cmd.owner = addressArg(params, "OWNER", "Address of the owner")
	cmd.spender = addressArg(params, "SPENDER", "Address of the spender")
}

func (cmd *cmdQueryTokenAllowance) Execute(ctx context.Context) error {
	caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	allowance, err := caller.GetAllowance(ctx, cmd.owner, cmd.spender)
	if err != nil {
		return err
	}
	fmt.Fprintln(clingy.Stdout(ctx), allowance)
	return nil
}
