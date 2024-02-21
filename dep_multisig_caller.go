package main

import (
	"context"

	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/multisig"
)

type depMultiSigCaller struct {
	client   depClient
	multiSig depMultiSig
}

func (dep *depMultiSigCaller) setup(params clingy.Parameters) {
	dep.client.setup(params)
	dep.multiSig.setup(params)
}

func (dep *depMultiSigCaller) open(ctx context.Context) (eth.Client, *multisig.Caller, error) {
	client, err := dep.client.open(ctx)
	if err != nil {
		return nil, nil, err
	}
	caller, err := multisig.NewCaller(client, dep.multiSig.contractAddress)
	if err != nil {
		return nil, nil, err
	}
	return client, caller, nil
}
