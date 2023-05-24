package test

import (
	"path/filepath"
	"runtime"
)

func packageDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}

func check[R any](r R, err error) R {
	if err != nil {
		panic(err)
	}
	return r
}
