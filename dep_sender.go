package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
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
	sender           common.Address
	senderKeyFile    string
	chainID          *big.Int
	clefEndpoint     string
	skipConfirmation bool
}

func (dep *depSender) setup(params clingy.Parameters) {
	dep.sender = optionalAddressEnvFlag(params, "sender", "Sender address", common.Address{}, envSender)
	dep.senderKeyFile = optionalStringEnvFlag(params, "sender-key-file", "Path on disk to the sender private key", "", envSenderKeyFile)
	dep.clefEndpoint = optionalStringEnvFlag(params, "clef-endpoint", "Clef endpoint", "", envClefEndpoint)
	dep.chainID = optionalBigIntEnvFlag(params, "chain-id", "Chain ID", defChainID, envChainID)
	dep.skipConfirmation = boolEnvFlag(params, "skip-confirmation", "Skips confirmation before transacting", envSkipConfirmation)
}

func (dep *depSender) transactOpts(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts, error) {
	// --sender-key-file and --clef-endpoint are mutually exclusive
	var opts *bind.TransactOpts
	switch {
	case dep.clefEndpoint == "" && dep.senderKeyFile == "":
		return nil, errors.New("must specify one of --sender-key-file or --clef-endpoint")
	case dep.clefEndpoint != "" && dep.senderKeyFile == "":
		clef, err := external.NewExternalSigner(dep.clefEndpoint)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if dep.sender == (common.Address{}) {
			return nil, fmt.Errorf("--sender (or env %s) is required when using clef", envSender)
		}
		opts, err = bind.NewClefTransactor(clef, accounts.Account{Address: dep.sender}), nil
		if err != nil {
			return nil, errs.Wrap(err)
		}
		// The hardware wallets only support "legacy" EIP155 transactions, so
		// we need to define the gas price.
		opts.GasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, errs.Wrap(err)
		}
	case dep.clefEndpoint == "" && dep.senderKeyFile != "":
		key, sender, err := loadETHKey(dep.senderKeyFile, "sender key file")
		if err != nil {
			return nil, errs.Wrap(err)
		}
		switch {
		case dep.sender == (common.Address{}):
			dep.sender = sender
		case dep.sender != sender:
			return nil, fmt.Errorf("sender %q does not match sender key address %q", dep.sender, sender)
		}
		opts, err = bind.NewKeyedTransactorWithChainID(key, dep.chainID)
		if err != nil {
			return nil, errs.Wrap(err)
		}
	case dep.clefEndpoint != "" && dep.senderKeyFile != "":
		return nil, errors.New("cannot specify both --sender-key-file and --clef-endpoint")
	}

	opts.Signer = confirmingSigner(ctx, opts.Signer, dep.skipConfirmation)
	opts.Context = ctx
	return opts, nil
}

func confirmingSigner(ctx context.Context, signer bind.SignerFn, skip bool) bind.SignerFn {
	return bind.SignerFn(func(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
		in := clingy.Stdin(ctx)
		out := clingy.Stdout(ctx)

		call := tryDecodeCall(tx.Data())

		fmt.Fprintf(out, "Preparing to send transaction:\n")
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
		return signer(sender, tx)
	})
}

type nopWriterCloser struct{ io.Writer }

func (nopWriterCloser) Close() error { return nil }
