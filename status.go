package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/zeebo/clingy"
	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/multisig"
)

var (
	abis         []*abi.ABI
	initABIsOnce sync.Once
)

func initABIs() {
	appendABI := func(metadata interface{ GetAbi() (*abi.ABI, error) }) {
		abi, err := metadata.GetAbi()
		if err != nil {
			panic(err)
		}
		abis = append(abis, abi)
	}

	initABIsOnce.Do(func() {
		appendABI(contract.BurnableTokenMetaData)
		appendABI(contract.CentrallyIssuedTokenMetaData)
		appendABI(contract.ERC20MetaData)
		appendABI(contract.FactoryMetaData)
		appendABI(contract.MultiSigWalletMetaData)
		appendABI(contract.MultiSigWalletWithDailyLimitMetaData)
		appendABI(contract.SafeMathMetaData)
		appendABI(contract.StandardTokenMetaData)
		appendABI(contract.UpgradeAgentMetaData)
		appendABI(contract.UpgradeableTokenMetaData)
	})
}

func printTransactionStatus(ctx context.Context, caller *multisig.Caller, transactionID uint64) error {
	out := clingy.Stdout(ctx)

	tx, err := caller.FindTransaction(ctx, transactionID)
	switch {
	case err != nil:
		return err
	case tx == nil:
		fmt.Fprintf(out, "Transaction %d: <empty>\n", transactionID)
		return nil
	}

	confirmed, err := caller.IsTransactionConfirmed(ctx, transactionID)
	if err != nil {
		return err
	}
	confirmations, err := caller.GetTransactionConfirmations(ctx, transactionID)
	if err != nil {
		return err
	}
	events, err := caller.GetTransactionEvents(ctx, transactionID)
	if err != nil {
		return err
	}

	dataRaw := "<empty>"
	dataDecoded := "<empty>"

	if len(tx.Data) > 0 {
		dataRaw = hex.EncodeToString(tx.Data)
		dataDecoded = tryDecodeCall(tx.Data)
		if dataDecoded == "" {
			dataDecoded = "<unable to decode>"
		}
	}

	fmt.Fprintf(out, "Transaction %d:\n", transactionID)
	fmt.Fprintf(out, "  Destination    = %s\n", tx.Destination)
	fmt.Fprintf(out, "  Value          = %s\n", tx.Value)
	fmt.Fprintf(out, "  Data (raw)     = %s\n", dataRaw)
	fmt.Fprintf(out, "  Data (decoded) = %s\n", dataDecoded)
	fmt.Fprintf(out, "  Executed       = %t\n", tx.Executed)
	fmt.Fprintf(out, "  Confirmed      = %t\n", confirmed)
	fmt.Fprintf(out, "  Confirmations(%d):\n", len(confirmations))
	for _, confirmation := range confirmations {
		fmt.Fprintf(out, "    - %s\n", confirmation)
	}
	fmt.Fprintf(out, "  Events(%d):\n", len(events))
	for _, event := range events {
		fmt.Fprintf(out, "    - %s\n", event)
	}
	return nil
}

func tryDecodeCall(data []byte) string {
	initABIs()
	for _, a := range abis {
		if decoded := tryDecodeABICall(a, data); decoded != "" {
			return decoded
		}
	}
	return ""
}

func tryDecodeABICall(abi *abi.ABI, data []byte) string {
	if len(data) < 4 {
		return ""
	}
	for _, m := range abi.Methods {
		if !bytes.Equal(m.ID, data[:4]) {
			continue
		}

		args, err := m.Inputs.Unpack(data[4:])
		if err != nil {
			return ""
		}

		buf := new(bytes.Buffer)
		buf.WriteString(m.Name)
		buf.WriteString("(")
		for i, arg := range args {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprint(buf, arg)
		}
		buf.WriteString(")")
		return buf.String()
	}
	return ""
}
