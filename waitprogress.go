package main

import (
	"context"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/eth"
)

func waiter(ctx context.Context, client *ethclient.Client) eth.Waiter {
	return eth.ProgressWaiter(client, newWaitProgress(ctx))
}

type waitProgress struct {
	stdout io.Writer
	stderr io.Writer
}

func newWaitProgress(ctx context.Context) *waitProgress {
	return &waitProgress{
		stdout: clingy.Stdout(ctx),
		stderr: clingy.Stderr(ctx),
	}
}

func (p *waitProgress) Start(hash common.Hash) {
	fmt.Fprintf(p.stdout, "Transaction hash is %s\n", hash.String())
	fmt.Fprintf(p.stdout, "Waiting for transaction to be confirmed...\n")
}

func (p *waitProgress) Tick() {
	fmt.Fprint(p.stdout, ".")
}

func (p *waitProgress) Canceled(err error) {
	fmt.Fprintf(p.stderr, "Wait canceled (%+v). Transaction may still confirm.\n", err)
}

func (p *waitProgress) Failed(status uint64) {
	fmt.Fprintln(p.stdout)
	fmt.Fprintf(p.stderr, "Transaction failed with status %d\n", status)
}

func (p *waitProgress) Dropped() {
	fmt.Fprintln(p.stdout)
	fmt.Fprintln(p.stderr, "Transaction was dropped")
}

func (p *waitProgress) Success() {
	fmt.Fprintln(p.stdout)
}

func (p *waitProgress) TempError(err error) {
	fmt.Fprintln(p.stdout)
	fmt.Fprintln(p.stderr, "Error:", err)
}
