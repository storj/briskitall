package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdQueryMultiSigOwnerList struct {
	caller depMultiSigCaller
}

func (cmd *cmdQueryMultiSigOwnerList) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
}

func (cmd *cmdQueryMultiSigOwnerList) Execute(ctx context.Context) error {
	_, caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	addresses, err := caller.ListOwners(ctx)
	if err != nil {
		return err
	}
	for _, address := range addresses {
		fmt.Fprintln(clingy.Stdout(ctx), address)
	}
	return nil
}
