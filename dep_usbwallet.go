package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"
)

var (
	ethDerivationPath = accounts.DefaultBaseDerivationPath
)

type depUSBWalletAccount struct {
	account               common.Address
	accountDerivationPath accounts.DerivationPath
}

func (dep *depUSBWalletAccount) enabled() bool {
	return dep.account != (common.Address{})
}

func (dep *depUSBWalletAccount) setup(params clingy.Parameters) {
	dep.account = optionalAddressEnvFlag(params, "usb-wallet-account", "USB wallet account", common.Address{}, envUSBWalletAccount)
	dep.accountDerivationPath = optionalDerivationPathFlag(params, "usb-wallet-account-derivation-path", "Account derivation path", ethDerivationPath)
}

func (dep *depUSBWalletAccount) transactOpts(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, func(), error) {
	wallets, err := enumerateUSBWallets()
	if err != nil {
		return nil, nil, err
	}

	var deriveErrs errs.Group

	// Search each wallet for the specified account. Don't self-derive.
	for _, wallet := range wallets {
		wallet := wallet
		wallet.SelfDerive(nil, nil)
		if err := wallet.Open(""); err != nil {
			return nil, nil, errs.New("failed to open wallet %s: %v", wallet.URL(), err)
		}
		defer func() {
			if err != nil {
				_ = wallet.Close()
			}
		}()

		status, err := wallet.Status()
		if err != nil {
			fmt.Fprintf(clingy.Stderr(ctx), "Failed to determine wallet %q status: %v\n", wallet.URL(), err)
			continue
		}
		fmt.Fprintf(clingy.Stdout(ctx), "%s: status: %s\n", wallet.URL(), status)

		account, err := wallet.Derive(dep.accountDerivationPath, true)
		if err != nil {
			deriveErrs.Add(err)
			continue
		}
		if account.Address != dep.account {
			continue
		}

		fmt.Fprintf(clingy.Stdout(ctx), "%s: using account %s (%s)\n", wallet.URL(), account.Address, strings.TrimPrefix(account.URL.String(), wallet.URL().String()))
		return &bind.TransactOpts{
				From: account.Address,
				Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
					if address != account.Address {
						return nil, bind.ErrNotAuthorized
					}
					return wallet.SignTx(account, transaction, chainID)
				},
				Context: context.Background(),
			}, func() {
				_ = wallet.Close()
			}, nil
	}

	deriveErrs.Add(errs.New("no USB wallet account available"))
	return nil, nil, deriveErrs.Err()
}
