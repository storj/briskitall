package main

import (
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeebo/errs"
)

func loadETHKey(path, which string) (*ecdsa.PrivateKey, common.Address, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, common.Address{}, errs.New("unable to stat %s key: %v\n", which, err)
	}

	if (fi.Mode() & 0177) != 0 {
		return nil, common.Address{}, errs.New("%s mode %#o is too permissive (set to 0600)\n", path, fi.Mode())
	}

	key, err := crypto.LoadECDSA(path)
	if err != nil {
		return nil, common.Address{}, errs.New("unable to load %s key: %v\n", which, err)
	}
	return key, crypto.PubkeyToAddress(key.PublicKey), nil
}
