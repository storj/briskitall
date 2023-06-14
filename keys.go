package main

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

func loadETHKey(path string) (*ecdsa.PrivateKey, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, errs.New("unable to stat key: %v\n", err)
	}

	if (fi.Mode() & 0177) != 0 {
		return nil, errs.New("%s mode %#o is too permissive (set to 0600)\n", path, fi.Mode())
	}

	key, err := crypto.LoadECDSA(path)
	if err != nil {
		return nil, errs.New("unable to load key: %v\n", err)
	}
	return key, nil
}
