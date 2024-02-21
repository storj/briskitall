package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/clingy"

	"storj.io/briskitall/internal/contract"
	"storj.io/briskitall/internal/eth"
	"storj.io/briskitall/internal/multisig"
)

var (
	abis         []*abi.ABI
	initABIsOnce sync.Once

	ethereumDecimals = 18
	knownDecimals    = map[common.Address]int{
		// STORJ
		common.HexToAddress("0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac"): 8,
	}
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

func printTransactionStatus(ctx context.Context, client eth.Client, nicknames multisig.Nicknames, caller *multisig.Caller, transactionID uint64) error {
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

	blockTimestamps := map[uint64]time.Time{}
	blockTimestamp := func(blockNo uint64) (time.Time, error) {
		if ts, exists := blockTimestamps[blockNo]; exists {
			return ts, nil
		}
		block, err := client.BlockByNumber(ctx, (&big.Int{}).SetUint64(blockNo))
		if err != nil {
			return time.Time{}, err
		}
		blockTimestamps[blockNo] = time.Unix(int64(block.Time()), 0)
		return blockTimestamps[blockNo], nil
	}

	call := tryDecodeCall(nicknames, tx.Data, knownDecimals[tx.Destination])

	fmt.Fprintf(out, "Transaction %d:\n", transactionID)
	fmt.Fprintf(out, "  Destination = %s\n", nicknames.Lookup(tx.Destination))
	switch {
	case len(tx.Data) == 0:
		fmt.Fprintf(out, "  Value       = %s\n", humanInt(tx.Value, ethereumDecimals))
	case call != "":
		fmt.Fprintf(out, "  Call        = %s\n", call)
	default:
		fmt.Fprintf(out, "  Data (raw)  = %x\n", tx.Data)
	}
	fmt.Fprintf(out, "  Executed    = %t\n", tx.Executed)
	fmt.Fprintf(out, "  Confirmed   = %t\n", confirmed)
	fmt.Fprintf(out, "  Confirmations(%d):\n", len(confirmations))
	for _, confirmation := range confirmations {
		fmt.Fprintf(out, "    - Owner(%s)\n", nicknames.Lookup(confirmation))
	}
	fmt.Fprintf(out, "  Events(%d):\n", len(events))
	for _, event := range events {
		ts, err := blockTimestamp(event.BlockNumber())
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "    - %s: %s\n", ts.Format("2006-01-02"), event.StringWithNicknames(nicknames))
	}
	return nil
}

func tryDecodeCall(nicknames multisig.Nicknames, data []byte, decimals int) string {
	initABIs()
	for _, a := range abis {
		if decoded := tryDecodeABICall(nicknames, a, data, decimals); decoded != "" {
			return decoded
		}
	}
	return ""
}

func tryDecodeABICall(nicknames multisig.Nicknames, abi *abi.ABI, data []byte, decimals int) string {
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
			if argData, ok := arg.([]byte); ok {
				if argCall := tryDecodeCall(nicknames, argData, decimals); argCall != "" {
					fmt.Fprint(buf, argCall)
				} else {
					fmt.Fprintf(buf, "%x", arg)
				}
				continue
			}
			switch argT := arg.(type) {
			case common.Address:
				fmt.Fprint(buf, nicknames.Lookup(argT))
			case *big.Int:
				fmt.Fprint(buf, humanInt(argT, decimals))
			default:
				fmt.Fprint(buf, arg)
			}
		}
		buf.WriteString(")")
		return buf.String()
	}
	return ""
}

// humanInt will turn v into a string representation of a base 10 integer
// with US locale commas for thousands, millions, etc. if decimals is greater
// than zero, this integer is assumed to be a decimal value, where the least
// significant 'decimals' amount of digits are after a period.
//
// Examples:
//
//	humanInt(big.NewInt(1000), 0) -> "1,000"
//	humanInt(big.NewInt(100000000), 4) -> "10,000.0000"
func humanInt(v *big.Int, decimals int) string {
	val := v.Text(10)
	negative := val[0] == '-'
	if negative {
		val = val[1:]
	}

	if len(val) < decimals+1 {
		val = strings.Repeat("0", decimals+1-len(val)) + val
	}

	out := make([]byte, len(val)+len(val)/3+2)
	j := len(out) - 1

	var i int
	for i = 0; i < decimals; i++ {
		out[j] = val[len(val)-1-i]
		j--
	}
	if decimals > 0 {
		out[j] = '.'
		j--
	}

	for k := 0; i < len(val); i, k = i+1, k+1 {
		if k%3 == 0 && k > 0 {
			out[j] = ','
			j--
		}
		out[j] = val[len(val)-1-i]
		j--
	}

	result := string(bytes.Trim(out, "\x00"))
	if negative {
		return "-" + result
	}
	return result
}
