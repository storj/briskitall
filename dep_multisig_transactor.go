package main

import (
	"context"

	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/multisig"
)

type depMultiSigTransactor struct {
	client   depClient
	multiSig depMultiSig
	sender   depSender
}

func (dep *depMultiSigTransactor) setup(params clingy.Parameters) {
	dep.client.setup(params)
	dep.multiSig.setup(params)
	dep.sender.setup(params)
}

func (dep *depMultiSigTransactor) open(ctx context.Context) (eth.Client, *multisig.Transactor, func(), error) {
	client, err := dep.client.open(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	opts, done, err := dep.sender.transactOpts(ctx, client)
	if err != nil {
		return nil, nil, nil, err
	}

	transactor, err := multisig.NewTransactor(client, dep.multiSig.contractAddress, opts, waiter(ctx, client))
	if err != nil {
		done()
		return nil, nil, nil, err
	}
	return client, transactor, done, nil
}
