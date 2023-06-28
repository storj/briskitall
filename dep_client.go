package main

import (
	"context"

	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
)

type depClient struct {
	nodeURL               string
	gasEstimateMultiplier float64
}

func (dep *depClient) setup(params clingy.Parameters) {
	dep.nodeURL = requiredStringEnvFlag(params, "node-url", "Ethereum Node URL", envNodeURL)
	dep.gasEstimateMultiplier = float64Flag(params, "gas-estimate-multiplier", "Multiplier on the estimated gas price", 2.0)
}

func (dep *depClient) open(ctx context.Context) (eth.Client, error) {
	return eth.Dial(ctx, dep.nodeURL, eth.GasEstimateMultiplier(dep.gasEstimateMultiplier))
}
