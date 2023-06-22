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

	rawTx := types.NewTx(&types.DynamicFeeTx{
		To:        &recipient,
		Nonce:     nonce,
		GasFeeCap: big.NewInt(50000000000), // 50 GWei
		GasTipCap: big.NewInt(2000000000),  // 2 GWei
		Gas:       25000,                   // sends to the multisig contract can be 22511
		Value:     amount,
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
