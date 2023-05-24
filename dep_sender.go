package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"
)

var (
	defChainID = big.NewInt(1)
)

type depSender struct {
	sender        common.Address
	senderKeyFile string
	chainID       *big.Int
	clefEndpoint  string
}

func (dep *depSender) setup(params clingy.Parameters) {
	dep.sender = optionalAddressEnvFlag(params, "sender", "Sender address", common.Address{}, envSender)
	dep.senderKeyFile = optionalStringEnvFlag(params, "sender-key-file", "Path on disk to the sender private key", "", envSenderKeyFile)
	dep.clefEndpoint = optionalStringEnvFlag(params, "clef-endpoint", "Clef endpoint", "", envClefEndpoint)
	dep.chainID = optionalBigIntEnvFlag(params, "chain-id", "Chain ID", defChainID, envChainID)
}

func (dep *depSender) transactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	// --sender-key-file and --clef-endpoint are mutually exclusive
	switch {
	case dep.clefEndpoint != "" && dep.senderKeyFile == "":
		clef, err := external.NewExternalSigner(dep.clefEndpoint)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		return bind.NewClefTransactor(clef, accounts.Account{Address: dep.sender}), nil
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
		opts, err := bind.NewKeyedTransactorWithChainID(key, dep.chainID)
		if err != nil {
			return nil, errs.Wrap(err)
		}
		return opts, nil
	case dep.clefEndpoint != "" && dep.senderKeyFile != "":
		return nil, errors.New("cannot specify both --sender-key-file and --clef-endpoint")
	}

	return nil, errors.New("must specify one of --sender-key-file or --clef-endpoint")
}
