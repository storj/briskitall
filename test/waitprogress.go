package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"storj.io/briskitall/internal/eth"
)

func WaitProgress(t *testing.T) eth.WaitProgress {
	return &waitProgress{t: t}
}

type waitProgress struct {
	t    *testing.T
	hash common.Hash
}

func (w *waitProgress) Start(hash common.Hash) { w.hash = hash }
func (w *waitProgress) Tick()                  { w.t.Logf("%s: waiting...", w.hash) }
func (w *waitProgress) Canceled(err error)     { w.t.Logf("%s: wait canceled: %v", w.hash, err) }
func (w *waitProgress) Failed(status uint64)   { w.t.Logf("%s: failed: status=%d", w.hash, status) }
func (w *waitProgress) Dropped()               { w.t.Logf("%s: dropped", w.hash) }
func (w *waitProgress) Success()               { w.t.Logf("%s: success", w.hash) }
func (w *waitProgress) TempError(err error)    { w.t.Logf("%s: temp error: %v", w.hash, err) }
