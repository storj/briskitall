package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdQueryTokenUpgradeMaster struct {
	caller depTokenCaller
}

func (cmd *cmdQueryTokenUpgradeMaster) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
}

func (cmd *cmdQueryTokenUpgradeMaster) Execute(ctx context.Context) error {
	caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	master, err := caller.GetUpgradeMaster(ctx)
	if err != nil {
		return err
	}
	fmt.Fprintln(clingy.Stdout(ctx), master)
	return nil
}
