package eth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type WaitProgress interface {
	Start(hash common.Hash)
	Tick()
	Canceled(err error)
	Failed(status uint64)
	Dropped()
	Success()
	TempError(err error)
}

type NoProgress struct{}

func (NoProgress) Start(hash common.Hash) {}
func (NoProgress) Tick()                  {}
func (NoProgress) Canceled(err error)     {}
func (NoProgress) Failed(status uint64)   {}
func (NoProgress) Dropped()               {}
func (NoProgress) Success()               {}
func (NoProgress) TempError(err error)    {}

type WaitBackend interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	TransactionByHash(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error)
}

type Waiter interface {
	Wait(ctx context.Context, hash common.Hash) (uint64, error)
}

func SilentWaiter(backend WaitBackend) Waiter {
	return silentWaiter{
		backend: backend,
	}
}

type silentWaiter struct{ backend WaitBackend }

func (w silentWaiter) Wait(ctx context.Context, hash common.Hash) (uint64, error) {
	return WaitForTransaction(ctx, w.backend, hash, NoProgress{})
}

func ProgressWaiter(backend WaitBackend, progress WaitProgress) Waiter {
	return progressWaiter{backend: backend, progress: progress}
}

type progressWaiter struct {
	backend  WaitBackend
	progress WaitProgress
}

func (w progressWaiter) Wait(ctx context.Context, hash common.Hash) (uint64, error) {
	return WaitForTransaction(ctx, w.backend, hash, w.progress)
}

func WaitForTransaction(ctx context.Context, backend WaitBackend, hash common.Hash, progress WaitProgress) (uint64, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if progress == nil {
		progress = NoProgress{}
	}
	progress.Start(hash)
	for attempts := 0; ; attempts++ {
		if attempts%10 == 0 {
			progress.Tick()
		}
		select {
		case <-time.After(time.Millisecond * 100):
		case <-ctx.Done():
			progress.Canceled(ctx.Err())
			return 0, ctx.Err()
		}

		receipt, err := backend.TransactionReceipt(ctx, hash)
		switch {
		case err == nil:
			if receipt.Status == types.ReceiptStatusSuccessful {
				progress.Success()
				return receipt.BlockNumber.Uint64(), nil
			}
			progress.Failed(receipt.Status)
			return 0, errors.New("transaction failed")
		case err == ethereum.NotFound:
		default:
			progress.TempError(fmt.Errorf("failed to query for transaction receipt: %+v\n", err))
		}

		_, _, err = backend.TransactionByHash(ctx, hash)
		switch {
		case err == nil:
		case err == ethereum.NotFound:
			progress.Dropped()
			return 0, errors.New("transaction dropped")
		default:
			progress.TempError(fmt.Errorf("failed to query for transaction by hash: %+v\n", err))
		}
	}
}
