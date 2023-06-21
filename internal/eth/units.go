package eth

import (
	"math/big"
	"regexp"
	"strings"

	"github.com/zeebo/errs"
)

var (
	suffixRE = regexp.MustCompile(`([a-zA-Z]+)$`)
)

// Returns WEI amount from string. Units accepted are:
// -ETH
// -GWEI
// -WEI (default if unset)
func ParseUnit(s string) (*big.Int, error) {
	if s == "" {
		return nil, errs.New("invalid unit: empty")
	}

	var rawSuffix string
	if m := suffixRE.FindStringSubmatch(s); m != nil {
		rawSuffix = m[1]
	}

	s = s[:len(s)-len(rawSuffix)]

	suffix := strings.ToUpper(rawSuffix)
	if suffix == "" {
		suffix = "WEI"
	}

	var mult int64
	switch suffix {
	case "ETH":
		mult = 1e18
	case "GWEI":
		mult = 1e9
	case "WEI":
		mult = 1
	default:
		return nil, errs.New("unsupported suffix %q", rawSuffix)
	}

	// use float so that we can accept scientific notation
	f, ok := new(big.Float).SetString(s)
	if !ok {
		return nil, errs.New("%s is not a valid %s unit: not a valid float", s, suffix)
	}

	f.Mul(f, new(big.Float).SetInt(big.NewInt(mult)))

	i, a := f.Int(nil)
	// make sure there is no fractional component
	if a != big.Exact {
		return nil, errs.New("%s is not a valid %s unit: must be a whole number", s, suffix)
	}
	return i, nil
}
