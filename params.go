package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/eth"
)

func int64Arg(params clingy.Parameters, name, desc string) int64 {
	return params.Arg(name, desc, clingy.Transform(transformInt64)).(int64)
}

func uint64Arg(params clingy.Parameters, name, desc string) uint64 {
	return params.Arg(name, desc, clingy.Transform(transformUint64)).(uint64)
}

func bigIntArg(params clingy.Parameters, name, desc string) *big.Int {
	return params.Arg(name, desc, clingy.Transform(transformBigInt)).(*big.Int)
}

func ethUnitArg(params clingy.Parameters, name, desc string) *big.Int {
	return params.Arg(name, desc, clingy.Transform(eth.ParseUnit)).(*big.Int)
}

func stringArg(params clingy.Parameters, name, desc string) string {
	return params.Arg(name, desc).(string)
}

func repeatedStringArg(params clingy.Parameters, name, desc string) []string {
	return params.Arg(name, desc, clingy.Repeated).([]string)
}

func addressArg(params clingy.Parameters, name, desc string) common.Address {
	return params.Arg(name, desc, clingy.Transform(transformAddress)).(common.Address)
}

func optAddressArg(params clingy.Parameters, name, desc string) *common.Address {
	return params.Arg(name, desc, clingy.Optional, clingy.Transform(transformAddress)).(*common.Address)
}

func repeatedAddressArg(params clingy.Parameters, name, desc string) []common.Address {
	return params.Arg(name, desc, clingy.Repeated, clingy.Transform(transformAddress)).([]common.Address)
}

func toggleFlag(params clingy.Parameters, name, desc string, def bool) bool {
	return params.Flag(name, desc, def, clingy.Transform(strconv.ParseBool)).(bool)
}

func boolFlag(params clingy.Parameters, name, desc string) bool {
	return params.Flag(name, desc, false, clingy.Transform(strconv.ParseBool), clingy.Boolean).(bool)
}

func boolEnvFlag(params clingy.Parameters, name, desc, env string) bool {
	return params.Flag(name, desc, false, clingy.Getenv(env), clingy.Transform(strconv.ParseBool), clingy.Boolean).(bool)
}

func stringFlag(params clingy.Parameters, name, desc, def string) string {
	return params.Flag(name, desc, def).(string)
}

func int64Flag(params clingy.Parameters, name, desc string, def int64) int64 {
	return params.Flag(name, desc, def, clingy.Transform(transformInt64)).(int64)
}

func uint64Flag(params clingy.Parameters, name, desc string, def uint64) uint64 {
	return params.Flag(name, desc, def, clingy.Transform(transformUint64)).(uint64)
}

func float64Flag(params clingy.Parameters, name, desc string, def float64) float64 {
	return params.Flag(name, desc, def, clingy.Transform(transformFloat64)).(float64)
}

func bigIntFlag(params clingy.Parameters, name, desc string, def *big.Int) *big.Int {
	return params.Flag(name, desc, def, clingy.Transform(transformBigInt)).(*big.Int)
}

func optionalStringEnvFlag(params clingy.Parameters, name, desc, def, env string) string {
	return params.Flag(name, desc, def, clingy.Getenv(env)).(string)
}

func requiredStringEnvFlag(params clingy.Parameters, name, desc, env string) string {
	return params.Flag(name, desc, clingy.Required, clingy.Getenv(env)).(string)
}

func optionalAddressEnvFlag(params clingy.Parameters, name, desc string, def common.Address, env string) common.Address {
	return params.Flag(name, desc, def, clingy.Getenv(env), clingy.Transform(transformAddress)).(common.Address)
}

func requiredAddressEnvFlag(params clingy.Parameters, name, desc string, env string) common.Address {
	return params.Flag(name, desc, clingy.Required, clingy.Getenv(env), clingy.Transform(transformAddress)).(common.Address)
}

func optionalBigIntEnvFlag(params clingy.Parameters, name, desc string, def *big.Int, env string) *big.Int {
	return params.Flag(name, desc, def, clingy.Getenv(env), clingy.Transform(transformBigInt)).(*big.Int)
}

func requiredBigIntEnvFlag(params clingy.Parameters, name, desc string, env string) *big.Int {
	return params.Flag(name, desc, clingy.Required, clingy.Getenv(env), clingy.Transform(transformBigInt)).(*big.Int)
}

func optionalDerivationPathFlag(params clingy.Parameters, name, desc string, def accounts.DerivationPath) accounts.DerivationPath {
	return params.Flag(name, desc, def, clingy.Transform(accounts.ParseDerivationPath)).(accounts.DerivationPath)
}

func transformInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func transformUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func transformFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func transformAddress(s string) (common.Address, error) {
	var address common.Address
	if err := address.UnmarshalText([]byte(s)); err != nil {
		return common.Address{}, err
	}
	return address, nil
}

func transformBigInt(s string) (*big.Int, error) {
	// use float so that we can accept scientific notation
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return nil, fmt.Errorf("not a valid big integer: %q", s)
	}
	i, a := f.Int(nil)
	// make sure it there is no fractional component
	if a != big.Exact {
		return nil, fmt.Errorf("not a valid big integer: %q", s)
	}
	return i, nil
}
