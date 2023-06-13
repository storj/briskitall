package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zeebo/errs"
)

func Transfer(opts *bind.TransactOpts, client bind.ContractTransactor, recipient common.Address, amount *big.Int, waiter Waiter) (err error) {
	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}

	nonce, err := client.PendingNonceAt(ctx, opts.From)
	if err != nil {
		return errs.Wrap(err)
	}

	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return errs.Wrap(err)
		}
	}

	rawTx := types.NewTx(&types.LegacyTx{
		To:       &recipient,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      uint64(21000), // standard gas for transfer
		Value:    amount,
	})

	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return errs.Wrap(err)
	}

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return errs.Wrap(err)
	}

	if _, err = waiter.Wait(ctx, signedTx.Hash()); err != nil {
		return errs.Wrap(err)
	}

	return nil
}
