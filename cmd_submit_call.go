package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"
)

type cmdSubmitContractCall struct {
	transactor depMultiSigTransactor
	address    common.Address
	abiPath    string
	function   string
	args       []string
	value      *big.Int
}

func (cmd *cmdSubmitContractCall) Setup(params clingy.Parameters) {
	cmd.transactor.setup(params)
	cmd.value = bigIntFlag(params, "value", "Ether value to send with the call", big.NewInt(0))
	cmd.address = addressArg(params, "ADDRESS", "Address of the contract to call")
	cmd.abiPath = stringArg(params, "ABIPATH", "Path on disk to the JSON representation of the ABI")
	cmd.function = stringArg(params, "FUNCTION", "Function to call")
	cmd.args = repeatedStringArg(params, "ARG", "Arguments to the contract function")
}

func (cmd *cmdSubmitContractCall) Execute(ctx context.Context) error {
	a, err := loadABI(cmd.abiPath)
	if err != nil {
		return err
	}

	data, err := packArgs(a, cmd.function, cmd.args...)
	if err != nil {
		return err
	}

	client, transactor, done, err := cmd.transactor.open(ctx)
	if err != nil {
		return err
	}
	defer done()

	transactionID, err := transactor.SubmitTransaction(ctx, cmd.address, cmd.value, data)
	if err != nil {
		return err
	}

	fmt.Fprintln(clingy.Stdout(ctx))
	return printTransactionStatus(ctx, client, cmd.transactor.client.nicknames, transactor.Caller, transactionID)
}

func loadABI(path string) (*abi.ABI, error) {
	abiBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	a, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &a, nil
}

func packArgs(a *abi.ABI, fn string, rawArgs ...string) ([]byte, error) {
	m, ok := a.Methods[fn]
	if !ok {
		return nil, fmt.Errorf("contract ABI does not define function %q", fn)
	}

	if len(rawArgs) != len(m.Inputs) {
		return nil, fmt.Errorf("function %q requires %d arguments (passed %d)", fn, len(m.Inputs), len(rawArgs))
	}

	var args []interface{}
	for i, input := range m.Inputs {
		rawArg := rawArgs[i]
		arg, err := convertArg(input.Type, rawArg)
		if err != nil {
			return nil, fmt.Errorf("unable to convert %q to argument %q for function %q: %v", rawArg, input.Name, fn, err)
		}
		args = append(args, arg)
	}

	data, err := a.Pack(fn, args...)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	return data, nil
}

func convertArg(typ abi.Type, arg string) (interface{}, error) {
	switch typ.T {
	case abi.IntTy, abi.UintTy:
		val, ok := new(big.Int).SetString(arg, 10)
		if !ok {
			return nil, errors.New("not a valid integer")
		}
		return val, nil
	case abi.BoolTy:
		return strconv.ParseBool(arg)
	case abi.StringTy:
		return arg, nil
	case abi.AddressTy:
		return transformAddress(arg)
	case abi.FixedBytesTy:
		val, err := hex.DecodeString(arg)
		if err != nil {
			return nil, err
		}
		if len(val) != typ.Size {
			return nil, fmt.Errorf("expected %d bytes but got %d", typ.Size, len(val))
		}
		return val, nil
	case abi.BytesTy:
		return hex.DecodeString(arg)
	case abi.HashTy, abi.FixedPointTy, abi.FunctionTy, abi.SliceTy, abi.ArrayTy, abi.TupleTy:
		return nil, fmt.Errorf("cannot convert arg to unsupported type (%d)", typ.T)
	default:
		return nil, fmt.Errorf("cannot convert arg to unknown type (%d)", typ.T)
	}
}
