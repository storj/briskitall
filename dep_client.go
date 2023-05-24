package main

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"
)

type depClient struct {
	nodeURL string
}

func (dep *depClient) setup(params clingy.Parameters) {
	dep.nodeURL = requiredStringEnvFlag(params, "node-url", "Ethereum Node URL", envNodeURL)
}

func (dep *depClient) open(ctx context.Context) (*ethclient.Client, error) {
	ethClient, err := ethclient.DialContext(ctx, dep.nodeURL)
	return ethClient, errs.Wrap(err)
}
