package main

import (
	"context"

	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/token"
)

type depTokenCaller struct {
	client depClient
	token  depToken
}

func (dep *depTokenCaller) setup(params clingy.Parameters) {
	dep.client.setup(params)
	dep.token.setup(params)
}

func (dep *depTokenCaller) open(ctx context.Context) (*token.Caller, error) {
	client, err := dep.client.open(ctx)
	if err != nil {
		return nil, err
	}
	return token.NewCaller(client, dep.token.contractAddress)
}
