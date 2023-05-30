package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
)

type depMultiSig struct {
	contractAddress common.Address
}

func (dep *depMultiSig) setup(params clingy.Parameters) {
	dep.contractAddress = requiredAddressEnvFlag(params, "multisig-contract-address", "Address of the multisig contract", envMultiSigContractAddress)
}
