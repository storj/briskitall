package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kyokomi/emoji/v2"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
)

func waiter(ctx context.Context, client *ethclient.Client) eth.Waiter {
	return eth.ProgressWaiter(client, newWaitProgress(ctx))
}

type waitProgress struct {
	stdout io.Writer
	s      *spinner.Spinner
	hash   common.Hash
}

func newWaitProgress(ctx context.Context) *waitProgress {
	stdout := clingy.Stdout(ctx)
	return &waitProgress{
		stdout: stdout,
		s:      spinner.New(spinner.CharSets[11], 100*time.Millisecond, spinner.WithWriter(stdout)),
	}
}

func (p *waitProgress) Start(hash common.Hash) {
	fmt.Fprintln(p.stdout)
	emoji.Fprintf(p.stdout, "ETH[%s]: Sent :page_facing_up:\n", hash)

	p.hash = hash
	p.s.Prefix = fmt.Sprintf("ETH[%s]: Waiting ", hash)
	p.s.Start()
}

func (p *waitProgress) Canceled() {
	p.s.Stop()
	emoji.Fprintf(p.stdout, "ETH[%s]: Wait cancelled :see_no_evil:(may still confirm).\n", p.hash)
}

func (p *waitProgress) Failed(status uint64) {
	p.s.Stop()
	emoji.Fprintf(p.stdout, "ETH[%s]: Failed :cross_mark:(status=%d)\n", p.hash, status)
}

func (p *waitProgress) Dropped() {
	p.s.Stop()
	emoji.Fprintf(p.stdout, "ETH[%s]: Dropped :frowning: \n", p.hash)
}

func (p *waitProgress) Confirmed() {
	p.s.Stop()
	emoji.Fprintf(p.stdout, "ETH[%s]: Confirmed :check_mark_button:\n", p.hash)
}

func (p *waitProgress) TempError(err error) {
	p.s.Suffix = fmt.Sprintf(" (temperr=%s)", err)
}
