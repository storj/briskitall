package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/multisig"
)

type depMultiSigTransactor struct {
	client           depClient
	multiSig         depMultiSig
	sender           depSender
	skipConfirmation bool
}

func (dep *depMultiSigTransactor) setup(params clingy.Parameters) {
	dep.client.setup(params)
	dep.multiSig.setup(params)
	dep.sender.setup(params)
	dep.skipConfirmation = boolEnvFlag(params, "skip-confirmation", "Skips confirmation before transacting", envSkipConfirmation)
}

func (dep *depMultiSigTransactor) open(ctx context.Context) (*multisig.Transactor, error) {
	client, err := dep.client.open(ctx)
	if err != nil {
		return nil, err
	}
	opts, err := dep.sender.transactOpts(ctx)
	if err != nil {
		return nil, err
	}

	backend := &confirmingBackend{Client: client, skip: dep.skipConfirmation}
	return multisig.NewTransactor(backend, dep.multiSig.contractAddress, opts.From, opts.Signer, waiter(ctx, client))
}

type confirmingBackend struct {
	*ethclient.Client
	skip bool
}

func (b *confirmingBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	in := clingy.Stdin(ctx)
	out := clingy.Stdout(ctx)

	gas, err := b.Client.EstimateGas(ctx, msg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas: %v", err)
	}

	call := tryDecodeCall(msg.Data)

	fmt.Fprintf(out, "Preparing to send transaction:\n")
	fmt.Fprintf(out, "  From...........: %s\n", msg.From)
	fmt.Fprintf(out, "  To.............: %s\n", msg.To)
	switch {
	case len(msg.Data) == 0:
		fmt.Fprintf(out, "  Value..........: %s\n", msg.Value)
	case call != "":
		fmt.Fprintf(out, "  Call...........: %s\n", call)
	default:
		fmt.Fprintf(out, "  Data...........: %x\n", msg.Data)
	}
	fmt.Fprintf(out, "  Gas Estimate...: %d\n", gas)

	if b.skip {
		return gas, nil
	}

	prompt := promptui.Prompt{
		Label:     "Send",
		IsConfirm: true,
		Stdin:     io.NopCloser(in),
		Stdout:    nopWriterCloser{Writer: out},
	}
	if _, err := prompt.Run(); err != nil {
		return 0, errors.New("aborted")
	}
	return gas, nil
}

type nopWriterCloser struct{ io.Writer }

func (nopWriterCloser) Close() error { return nil }
