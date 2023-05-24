package multisig

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
)

const (
	invalidInputErrorCode = -32000
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)

	multisigABI = func() *abi.ABI {
		a, err := contract.MultiSigWalletWithDailyLimitMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		return a
	}()

	erc20ABI = func() *abi.ABI {
		a, err := contract.ERC20MetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		return a
	}()

	upgradeableTokenABI = func() *abi.ABI {
		a, err := contract.UpgradeableTokenMetaData.GetAbi()
		if err != nil {
			panic(err)
		}
		return a
	}()
)

type Transaction struct {
	ID          uint64
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}

func (t Transaction) String() string {
	return fmt.Sprintf("tx(id=%d,dst=%s,val=%s,exec=%t,data=%x)", t.ID, t.Destination, t.Value, t.Executed, t.Data)
}

func getTransaction(ctx context.Context, caller *contract.MultiSigWalletWithDailyLimitCaller, id *big.Int) (Transaction, error) {
	tx, err := findTransaction(ctx, caller, id)
	if err != nil {
		return Transaction{}, err
	}
	if tx == nil {
		return Transaction{}, errs.New("no such transaction")
	}
	return *tx, nil
}

func findTransaction(ctx context.Context, caller *contract.MultiSigWalletWithDailyLimitCaller, id *big.Int) (*Transaction, error) {
	tx, err := caller.Transactions(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, errs.New("failed to retrieve transaction %s: %v", id.String(), err)
	}
	if tx.Destination == (common.Address{}) {
		return nil, nil
	}
	return &Transaction{
		ID:          id.Uint64(),
		Destination: tx.Destination,
		Value:       tx.Value,
		Data:        tx.Data,
		Executed:    tx.Executed,
	}, nil
}

func isInvalidInput(err error) bool {
	if ec, ok := err.(interface{ ErrorCode() int }); ok {
		return ec.ErrorCode() == invalidInputErrorCode
	}
	return false
}

func defaultCallOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}

func defaultTransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{Context: ctx}
}

func newUint64(u uint64) *big.Int {
	return new(big.Int).SetUint64(u)
}
