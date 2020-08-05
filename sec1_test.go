package elliptic

import (
	"encoding/pem"
	"io/ioutil"
	"testing"
)

func TestParseP192PrivateKey(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/go_elliptic_private_key.pem")
	if err != nil {
		t.Errorf("Failed to read private key: %s", err)
		return
	}

	block, _ := pem.Decode(b)
	priv, err := ParseP192PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("Failed to parse private key: %s", err)
		return
	}

	privByte, err := MarshalP192PrivateKey(priv)
	if err != nil {
		t.Errorf("Failed to marshal private key: %s", err)
		return
	}

	priv2, err := ParseP192PrivateKey(privByte)
	if err != nil {
		t.Errorf("Failed to parse marshaled-private key: %s", err)
		return
	}

	if priv.Curve != priv2.Curve || priv.D.Cmp(priv2.D) != 0 ||
		priv.X.Cmp(priv2.X) != 0 || priv.Y.Cmp(priv2.Y) != 0 {
		t.Errorf("got:%+v want:%+v", priv2, priv)
	}
}
