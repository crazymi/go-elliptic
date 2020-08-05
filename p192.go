package elliptic

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var (
	p192     *elliptic.CurveParams
	initonce sync.Once
)

// P192 returns a Curve which implements NIST P-192 (FIPS 186-3, section D.2.1),
// also known as secp192r1 or prime192v1. The CurveParams.Name of this Curve is "P-192".
//
// Multiple invocations of this function will return the same value, so it can
// be used for equality checks and switch statements.
//
// The cryptographic operations do not use constant-time algorithms.
func P192() elliptic.Curve {
	initonce.Do(initP192)
	elliptic.P384()
	return p192
}

func initP192() {
	// See FIPS 186-3, section D.2.1
	p192 = &elliptic.CurveParams{Name: "P-192"}
	p192.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFFFFFFFFFFFF", 16)
	p192.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFF99DEF836146BC9B1B4D22831", 16)
	p192.B, _ = new(big.Int).SetString("64210519E59C80E70FA7E9AB72243049FEB8DEECC146B9B1", 16)
	p192.Gx, _ = new(big.Int).SetString("188DA80EB03090F67CBF20EB43A18800F4FF0AFD82FF1012", 16)
	p192.Gy, _ = new(big.Int).SetString("07192B95FFC8DA78631011ED6B24CDD573F977A11E794811", 16)
	p192.BitSize = 192
}
