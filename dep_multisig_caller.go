package main

import (
	"context"

	"github.com/zeebo/clingy"

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

func (dep *depMultiSigCaller) open(ctx context.Context) (*multisig.Caller, error) {
	client, err := dep.client.open(ctx)
	if err != nil {
		return nil, err
	}
	return multisig.NewCaller(client, dep.multiSig.contractAddress)
}
