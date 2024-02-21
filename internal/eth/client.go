package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"
)

type Client interface {
	bind.ContractBackend
	bind.DeployBackend
	bind.PendingContractCaller
	ethereum.TransactionReader
	ethereum.ChainStateReader

	ChainID(ctx context.Context) (*big.Int, error)
	BlockByNumber(context.Context, *big.Int) (*types.Block, error)
	Close()
}

type ClientOption = func(*clientWrapper)

func GasEstimateMultiplier(m float64) ClientOption {
	return func(c *clientWrapper) {
		c.gasEstimateMultiplier = m
	}
}

func Dial(ctx context.Context, url string, opts ...ClientOption) (Client, error) {
	c, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	w := &clientWrapper{Client: c, gasEstimateMultiplier: 2.0}
	for _, opt := range opts {
		opt(w)
	}
	return w, nil
}

type clientWrapper struct {
	*ethclient.Client
	gasEstimateMultiplier float64
}

func (w *clientWrapper) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	gas, err = w.Client.EstimateGas(ctx, call)
	return uint64(float64(gas) * w.gasEstimateMultiplier), err
}
