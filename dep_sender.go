package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"
)

var (
	defChainID = big.NewInt(1)
)

type depSender struct {
	usbWallet        depUSBWalletAccount
	senderKeyFile    string
	chainID          *big.Int
	skipConfirmation bool
}

func (dep *depSender) setup(params clingy.Parameters) {
	dep.usbWallet.setup(params)
	dep.senderKeyFile = optionalStringEnvFlag(params, "sender-key-file", "Path on disk to the sender private key", "", envSenderKeyFile)
	dep.chainID = optionalBigIntEnvFlag(params, "chain-id", "Chain ID", defChainID, envChainID)
	dep.skipConfirmation = boolEnvFlag(params, "skip-confirmation", "Skips confirmation before transacting", envSkipConfirmation)
}

func (dep *depSender) transactOpts(ctx context.Context, client *ethclient.Client) (opts *bind.TransactOpts, done func(), err error) {
	done = func() {}

	senderChoices := 0
	eip155Only := false
	isUSBWallet := false

	if dep.senderKeyFile != "" {
		senderChoices++
		eip155Only = false
	}

	if dep.usbWallet.enabled() {
		senderChoices++
		eip155Only = true
		isUSBWallet = true
	}

	if senderChoices == 0 {
		return nil, nil, errors.New("must specify one of --sender-key-file or --usb-wallet-account")
	}
	if senderChoices > 1 {
		return nil, nil, errors.New("can only specify one of --sender-key-file or --usb-wallet-account")
	}

	if dep.senderKeyFile != "" {
		key, err := loadETHKey(dep.senderKeyFile)
		if err != nil {
			return nil, nil, errs.Wrap(err)
		}
		opts, err = bind.NewKeyedTransactorWithChainID(key, dep.chainID)
		if err != nil {
			return nil, nil, errs.Wrap(err)
		}
	}

	if dep.usbWallet.enabled() {
		opts, done, err = dep.usbWallet.transactOpts(ctx, dep.chainID)
		if err != nil {
			return nil, nil, errs.Wrap(err)
		}
		// Ensure the wallet is closed if there is an error
		defer func() {
			if err != nil {
				done()
			}
		}()
	}

	// For transactors that only support EIP155 transactions we need to
	// define the gas price to trigger a "legacy" transaction to be signed.
	if eip155Only {
		opts.GasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, nil, errs.Wrap(err)
		}
	}

	opts.Signer = confirmingSigner(ctx, opts.Signer, dep.skipConfirmation, isUSBWallet)
	opts.Context = ctx
	return opts, done, nil
}

func confirmingSigner(ctx context.Context, signer bind.SignerFn, skip, isUSBWallet bool) bind.SignerFn {
	return bind.SignerFn(func(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
		in := clingy.Stdin(ctx)
		out := clingy.Stdout(ctx)

		call := tryDecodeCall(tx.Data())

		fmt.Fprintf(out, "Preparing to send transaction:\n")
		fmt.Fprintf(out, "  Type...........: %s\n", txType(tx.Type()))
		fmt.Fprintf(out, "  From...........: %s\n", sender)
		fmt.Fprintf(out, "  To.............: %s\n", tx.To())
		switch {
		case len(tx.Data()) == 0:
			fmt.Fprintf(out, "  Value..........: %s\n", tx.Value())
		case call != "":
			fmt.Fprintf(out, "  Call...........: %s\n", call)
		default:
			fmt.Fprintf(out, "  Data...........: %x\n", tx.Data())
		}
		fmt.Fprintf(out, "  Gas Estimate...: %d\n", tx.Gas())
		fmt.Fprintf(out, "  Gas Price......: %s\n", tx.GasPrice())
		fmt.Fprintf(out, "  Gas Tip Cap....: %s\n", tx.GasTipCap())
		fmt.Fprintf(out, "  Gas Fee Cap....: %s\n", tx.GasFeeCap())

		if !skip {
			prompt := promptui.Prompt{
				Label:     "Send",
				IsConfirm: true,
				Stdin:     io.NopCloser(in),
				Stdout:    nopWriterCloser{Writer: out},
			}
			if _, err := prompt.Run(); err != nil {
				return nil, errors.New("aborted")
			}
		}

		if isUSBWallet {
			fmt.Fprintln(out)

			s := spinner.New(spinner.CharSets[11], 100*time.Millisecond, spinner.WithWriter(out))
			s.Prefix = "Waiting for confirmation on the USB wallet "
			s.Start()
			defer s.Stop()
		}
		return signer(sender, tx)
	})
}

type nopWriterCloser struct{ io.Writer }

func (nopWriterCloser) Close() error { return nil }

func txType(t uint8) string {
	switch t {
	case types.LegacyTxType:
		return "Legacy"
	case types.AccessListTxType:
		return "AccessList"
	case types.DynamicFeeTxType:
		return "Dynamic"
	default:
		return "Unknown"
	}
}
