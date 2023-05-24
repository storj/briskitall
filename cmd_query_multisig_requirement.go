package main

import (
	"context"
	"fmt"

	"github.com/zeebo/clingy"
)

type cmdQueryMultiSigRequirement struct {
	caller depMultiSigCaller
}

func (cmd *cmdQueryMultiSigRequirement) Setup(params clingy.Parameters) {
	cmd.caller.setup(params)
}

func (cmd *cmdQueryMultiSigRequirement) Execute(ctx context.Context) error {
	caller, err := cmd.caller.open(ctx)
	if err != nil {
		return err
	}
	required, err := caller.GetConfirmationRequirement(ctx)
	if err != nil {
		return err
	}
	fmt.Fprintln(clingy.Stdout(ctx), required)
	return nil
}
