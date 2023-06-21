package eth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/briskitall/internal/eth"
)

func TestParseUnits(t *testing.T) {
	for _, tc := range []struct {
		in  string
		out string
		err string
	}{
		{
			in:  "1.24ETH",
			out: "1240000000000000000",
		},
		{
			in:  "1.24eth",
			out: "1240000000000000000",
		},
		{
			in:  "1.24GWEI",
			out: "1240000000",
		},
		{
			in:  "1.24gwei",
			out: "1240000000",
		},
		{
			in:  "124WEI",
			out: "124",
		},
		{
			in:  "1.24e18",
			out: "1240000000000000000",
		},
		{
			in:  "",
			err: "invalid unit: empty",
		},
		{
			in:  "foo",
			err: `unsupported suffix "foo"`,
		},
		{
			in:  "1f1eth",
			err: "1f1 is not a valid ETH unit: not a valid float",
		},
		{
			in:  "1.24",
			err: "1.24 is not a valid WEI unit: must be a whole number",
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			out, err := eth.ParseUnit(tc.in)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
				assert.Nil(t, out)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.out, out.String())
		})
	}
}
