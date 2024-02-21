package main

import (
	"math/big"
	"testing"
)

func TestHumanInt(t *testing.T) {
	for _, tt := range []struct {
		val      *big.Int
		decimals int
		expected string
	}{
		{big.NewInt(0), 0, "0"},
		{big.NewInt(1), 0, "1"},
		{big.NewInt(10), 0, "10"},
		{big.NewInt(100), 0, "100"},
		{big.NewInt(1000), 0, "1,000"},
		{big.NewInt(10000), 0, "10,000"},
		{big.NewInt(100000), 0, "100,000"},
		{big.NewInt(1000000), 0, "1,000,000"},
		{big.NewInt(-1), 0, "-1"},
		{big.NewInt(-100), 0, "-100"},
		{big.NewInt(-1000), 0, "-1,000"},
		{big.NewInt(-100000), 0, "-100,000"},
		{big.NewInt(-1000000), 0, "-1,000,000"},
		{big.NewInt(-10000000000), 4, "-1,000,000.0000"},
		{big.NewInt(0), 1, "0.0"},
		{big.NewInt(0), 2, "0.00"},
		{big.NewInt(0), 10, "0.0000000000"},
		{big.NewInt(1), 1, "0.1"},
		{big.NewInt(1), 2, "0.01"},
		{big.NewInt(1), 5, "0.00001"},
		{big.NewInt(10), 1, "1.0"},
		{big.NewInt(10), 2, "0.10"},
		{big.NewInt(100), 1, "10.0"},
		{big.NewInt(100), 3, "0.100"},
		{big.NewInt(100), 4, "0.0100"},
		{big.NewInt(100), 2, "1.00"},
		{big.NewInt(1000), 1, "100.0"},
		{big.NewInt(10000), 1, "1,000.0"},
		{big.NewInt(1000000), 2, "10,000.00"},
	} {
		if humanInt(tt.val, tt.decimals) != tt.expected {
			t.Fatalf("expected %q (%d) to be %q", tt.val.String(), tt.decimals, tt.expected)
		}
	}
}
