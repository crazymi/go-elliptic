package elliptic

import (
	"math/big"
	"testing"
)

func TestP192_OnCurve(t *testing.T) {
	p192 := P192()

	tcs := []struct {
		Points    [2]*big.Int
		IsOnCurve bool
	}{
		{[2]*big.Int{p192.Params().Gx, p192.Params().Gy}, true},
		{[2]*big.Int{new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)}, false},
	}

	for _, tc := range tcs {
		if tc.IsOnCurve != p192.IsOnCurve(tc.Points[0], tc.Points[1]) {
			t.Errorf("claim that off curve point is on curve: (%v, %v)", tc.Points[0], tc.Points[1])
		}
	}
}
