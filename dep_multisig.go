package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

var (
	defMultiSigContractAddress = common.HexToAddress("0x0F564a2A5fDE73349890e86e9B2aA1639994bF2F")
)

type depMultiSig struct {
	contractAddress common.Address
}

func (dep *depMultiSig) setup(params clingy.Parameters) {
	dep.contractAddress = optionalAddressEnvFlag(params, "multisig-contract-address", "Address of the multisig contract", defMultiSigContractAddress, envMultiSigContractAddress)
}
