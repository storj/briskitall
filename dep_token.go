package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type depToken struct {
	contractAddress common.Address
}

func (dep *depToken) setup(params clingy.Parameters) {
	dep.contractAddress = requiredAddressEnvFlag(params, "token-contract-address", "Address of the token contract", envTokenContractAddress)
}
