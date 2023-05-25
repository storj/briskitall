package token

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
)

type Caller struct {
	caller *contract.CentrallyIssuedTokenCaller
}

type CallerBackend interface {
	bind.ContractCaller
}

func NewCaller(backend CallerBackend, contractAddress common.Address) (*Caller, error) {
	caller, err := contract.NewCentrallyIssuedTokenCaller(contractAddress, backend)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &Caller{
		caller: caller,
	}, nil
}

func (c *Caller) GetAllowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	opts := c.callOpts(ctx)
	allowance, err := c.caller.Allowance(opts, owner, spender)
	if err != nil {
		return nil, errs.New("failed to get allowance for %s from %s: %v", spender, owner, err)
	}
	return allowance, nil
}

func (c *Caller) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	opts := c.callOpts(ctx)
	balance, err := c.caller.BalanceOf(opts, address)
	if err != nil {
		return nil, errs.New("failed to get balance for %s: %v", address, err)
	}
	return balance, nil
}

func (c *Caller) GetUpgradeMaster(ctx context.Context) (common.Address, error) {
	opts := c.callOpts(ctx)
	master, err := c.caller.UpgradeMaster(opts)
	if err != nil {
		return common.Address{}, errs.New("failed to get upgrade master: %v", err)
	}
	return master, nil
}

func (c *Caller) callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}
