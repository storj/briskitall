package main

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/zeebo/errs"
)

func selfDerivationPaths(wallet accounts.Wallet) (out []accounts.DerivationPath) {
	if wallet.URL().Scheme == "ledger" {
		out = append(out, accounts.LegacyLedgerBaseDerivationPath)
	}
	out = append(out, accounts.DefaultBaseDerivationPath)
	return out
}

func enumerateUSBWallets() ([]accounts.Wallet, error) {
	hubFuncs := []func() (*usbwallet.Hub, error){
		usbwallet.NewLedgerHub,
		usbwallet.NewTrezorHubWithHID,
		usbwallet.NewTrezorHubWithWebUSB,
	}

	var wallets []accounts.Wallet
	for _, hubFunc := range hubFuncs {
		hub, err := hubFunc()
		if err != nil {
			return nil, errs.Wrap(err)
		}
		wallets = append(wallets, hub.Wallets()...)
	}
	return wallets, nil
}
