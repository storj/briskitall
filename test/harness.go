package test

import (
	_ "embed"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/multisig"
	"storj.io/briskitall/internal/token"
)

func Run(t *testing.T) *Harness {
	geth := runGeth(t)

	// Fund the accounts that will issue transactions with ETH
	geth.Fund(t, rootAddress, AccountAddress[0], AccountAddress[1], AccountAddress[2])

	multiSigContractAddress := deployMultiSigContract(t, geth.RootTransactor, geth.Client)
	tokenContractAddress := deployTokenContract(t, geth.RootTransactor, geth.Client, multiSigContractAddress)

	t.Logf("MultiSig Contract address: %s", multiSigContractAddress)
	t.Logf("Token Contract address: %s", tokenContractAddress)

	multiSigCaller, err := multisig.NewCaller(geth.Client, multiSigContractAddress)
	require.NoError(t, err)

	tokenCaller, err := token.NewCaller(geth.Client, tokenContractAddress)
	require.NoError(t, err)

	geth.Fund(t, multiSigContractAddress)

	return &Harness{
		ChainID: geth.ChainID,
		NodeURL: geth.NodeURL,
		ETH: ETHHarness{
			client: geth.Client,
		},
		MultiSig: MultiSigHarness{
			ContractAddress:      multiSigContractAddress,
			Caller:               multiSigCaller,
			chainID:              geth.ChainID,
			nodeURL:              geth.NodeURL,
			client:               geth.Client,
			tokenContractAddress: tokenContractAddress,
		},
		Token: TokenHarness{
			ContractAddress: tokenContractAddress,
			Caller:          tokenCaller,
			chainID:         geth.ChainID,
			client:          geth.Client,
		},
	}
}

type Harness struct {
	NodeURL  string
	ChainID  *big.Int
	ETH      ETHHarness
	MultiSig MultiSigHarness
	Token    TokenHarness
}
