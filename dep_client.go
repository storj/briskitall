package main

import (
	"context"

	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/multisig"
)

type depClient struct {
	nodeURL               string
	gasEstimateMultiplier float64
	nicknames             multisig.Nicknames
}

func (dep *depClient) setup(params clingy.Parameters) {
	dep.nodeURL = requiredStringEnvFlag(params, "node-url", "Ethereum Node URL", envNodeURL)
	dep.gasEstimateMultiplier = float64Flag(params, "gas-estimate-multiplier", "Multiplier on the estimated gas price", 2.0)
	dep.nicknames = optionalNicknameMap(params, "address-nicknames", "A comma separated list of wallet address nickname definitions, like name1:address,name2:address", nil, envAddressNicknames)
}

func (dep *depClient) open(ctx context.Context) (eth.Client, error) {
	return eth.Dial(ctx, dep.nodeURL, eth.GasEstimateMultiplier(dep.gasEstimateMultiplier))
}
