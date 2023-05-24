package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/multisig"
)

type cmdDebugDeployMultiSig struct {
	client     depClient
	sender     depSender
	owners     []common.Address
	required   int64
	dailyLimit int64
}

func (cmd *cmdDebugDeployMultiSig) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
	cmd.sender.setup(params)
	cmd.required = int64Flag(params, "required", "Number of owner confirmations needed to execute a transaction", 2)
	cmd.dailyLimit = int64Flag(params, "daily-limit", "Daily limit", 0)
	cmd.owners = repeatedAddressArg(params, "OWNER", "Address of the contract owner (need at least two)")
}

func (cmd *cmdDebugDeployMultiSig) Execute(ctx context.Context) error {
	if len(cmd.owners) < 2 {
		return errors.New("at least two OWNER's are required")
	}

	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}
	opts, err := cmd.sender.transactOpts(ctx)
	if err != nil {
		return err
	}

	contractAddress, err := multisig.DeployContract(opts, client, cmd.owners, cmd.required, cmd.dailyLimit, waiter(ctx, client))
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx), "Contract Address:", contractAddress)
	return nil
}
