package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

var (
	defTokenContractAddress = common.HexToAddress("0xB64ef51C888972c908CFacf59B47C1AfBC0Ab8aC")
)

type depToken struct {
	contractAddress common.Address
}

func (dep *depToken) setup(params clingy.Parameters) {
	dep.contractAddress = optionalAddressEnvFlag(params, "token-contract-address", "Address of the token contract", defTokenContractAddress, envTokenContractAddress)
}
