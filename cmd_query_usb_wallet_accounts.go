package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/zeebo/clingy"
)

type cmdQueryUSBWalletAccountList struct {
	client depClient
}

func (cmd *cmdQueryUSBWalletAccountList) Setup(params clingy.Parameters) {
	cmd.client.setup(params)
}

func (cmd *cmdQueryUSBWalletAccountList) Execute(ctx context.Context) error {
	client, err := cmd.client.open(ctx)
	if err != nil {
		return err
	}

	wallets, err := enumerateUSBWallets()
	if err != nil {
		return err
	}
	if len(wallets) == 0 {
		fmt.Fprintln(clingy.Stdout(ctx), "No USB wallets found.")
		return nil
	}

	for _, wallet := range wallets {
		wallet.SelfDerive(selfDerivationPaths(wallet), client)

		if err := wallet.Open(""); err != nil {
			fmt.Fprintf(clingy.Stderr(ctx), "Failed to open wallet %q: %v\n", wallet.URL(), err)
		}
		defer wallet.Close()

		status, err := wallet.Status()
		if err != nil {
			fmt.Fprintf(clingy.Stderr(ctx), "Failed to determine wallet %q status: %v\n", wallet.URL(), err)
			continue
		}

		fmt.Fprintf(clingy.Stdout(ctx), "%s: status: %s\n", wallet.URL(), status)

		// The wallet code is kind of boneheaded. The goroutine for self derivation
		// kicks off on the call to Open but the call to Accounts() does not
		// wait for self derivation to start and will not wait for the results
		// if it hasn't. Retry a few times if no accounts are returned.
		var walletAccounts []accounts.Account
		for tries := 0; tries < 50; tries++ {
			walletAccounts = wallet.Accounts()
			if len(walletAccounts) != 0 {
				break
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Millisecond * 100):
			}
		}

		if len(walletAccounts) == 0 {
			fmt.Fprintf(clingy.Stdout(ctx), "%s: no accounts: is the ETH app online and the device unlocked?\n", wallet.URL())
			continue
		}

		for _, walletAccount := range walletAccounts {
			fmt.Fprintf(clingy.Stdout(ctx), "%s: %s (%s)\n", wallet.URL(), walletAccount.Address, strings.TrimPrefix(walletAccount.URL.String(), wallet.URL().String()))
		}
	}

	return nil
}
