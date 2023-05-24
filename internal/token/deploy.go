package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/errs"
	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/eth"
)

func DeployContract(opts *bind.TransactOpts, client *ethclient.Client, owner common.Address, name string, symbol string, totalSupply *big.Int, decimals int64, waiter eth.Waiter) (common.Address, error) {
	contractAddress, contractTx, _, err := contract.DeployCentrallyIssuedToken(opts, client, owner, name, symbol, totalSupply, big.NewInt(decimals))
	if err != nil {
		return common.Address{}, errs.Wrap(err)
	}
	if _, err = eth.WaitForTransaction(opts.Context, client, contractTx.Hash(), nil); err != nil {
		return common.Address{}, errs.Wrap(err)
	}
	return contractAddress, nil
}
