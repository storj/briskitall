package multisig

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/eth"
)

func DeployContract(opts *bind.TransactOpts, client bind.ContractBackend, owners []common.Address, required int64, dailyLimit int64, waiter eth.Waiter) (common.Address, error) {
	contractAddress, contractTx, _, err := contract.DeployMultiSigWalletWithDailyLimit(opts, client, owners, big.NewInt(required), big.NewInt(dailyLimit))
	if err != nil {
		return common.Address{}, errs.Wrap(err)
	}
	if _, err = waiter.Wait(opts.Context, contractTx.Hash()); err != nil {
		return common.Address{}, errs.Wrap(err)
	}
	return contractAddress, nil
}
