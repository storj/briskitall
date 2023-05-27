package main

import (
	"context"
	"encoding/json"
	"io"

	"github.com/zeebo/clingy"
)

type outputtable interface {
	TextOut(w io.Writer)
}

type depOutput struct {
	jsonOut bool
}

func (dep *depOutput) setup(params clingy.Parameters) {
	dep.jsonOut = boolFlag(params, "jsonout", "Output JSON instead of text")
}

func (dep *depOutput) out(ctx context.Context, o outputtable) {
	w := clingy.Stdout(ctx)
	if !dep.jsonOut {
		o.TextOut(w)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(o)
}
