package elliptic

import "encoding/asn1"

var oidNamedCurveP192 = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 1}

func isP192CurveFromOID(oid asn1.ObjectIdentifier) bool {
	return oid.Equal(oidNamedCurveP192)
}
