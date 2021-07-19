package s

import (
	"crypto"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/curve25519"
)

type CurveParams struct {
	Name    string // the canonical name of the curve
	BitSize int    // the size of the underlying field
}

// KeyExchange is the interface defining all functions
// necessary for ECDH.
type KeyExchange interface {
	// GenerateKey generates a private/public key pair using entropy from rand.
	// If rand is nil, crypto/rand.Reader will be used.
	GenerateKey(rand io.Reader) (private crypto.PrivateKey, public crypto.PublicKey, err error)

	// Params returns the curve parameters - like the field size.
	Params() CurveParams

	// PublicKey returns the public key corresponding to the given private one.
	PublicKey(private crypto.PrivateKey) (public crypto.PublicKey)

	// Check returns a non-nil error if the peers public key cannot used for the
	// key exchange - for instance the public key isn't a point on the elliptic curve.
	// It's recommended to check peer's public key before computing the secret.
	Check(peersPublic crypto.PublicKey) (err error)

	// ComputeSecret returns the secret value computed from the given private key
	// and the peers public key.
	ComputeSecret(private crypto.PrivateKey, peersPublic crypto.PublicKey) (secret []byte)
}

type ecdh25519 struct{}

var curve25519Params = CurveParams{
	Name:    "Curve25519",
	BitSize: 255,
}

func X25519() KeyExchange {
	return ecdh25519{}
}

func (ecdh25519) GenerateKey(random io.Reader) (private crypto.PrivateKey, public crypto.PublicKey, err error) {
	if random == nil {
		random = rand.Reader
	}

	var pri, pub [32]byte
	_, err = io.ReadFull(random, pri[:])
	if err != nil {
		return
	}

	// From https://cr.yp.to/ecdh.html
	pri[0] &= 248
	pri[31] &= 127
	pri[31] |= 64

	curve25519.ScalarBaseMult(&pub, &pri)

	private = pri
	public = pub
	return
}

func (ecdh25519) Params() CurveParams { return curve25519Params }

func (ecdh25519) PublicKey(private crypto.PrivateKey) (public crypto.PublicKey) {
	var pri, pub [32]byte
	if ok := checkType(&pri, private); !ok {
		panic("ecdh: unexpected type of private key")
	}

	curve25519.ScalarBaseMult(&pub, &pri)

	public = pub
	return
}

func (ecdh25519) Check(peersPublic crypto.PublicKey) (err error) {
	if ok := checkType(new([32]byte), peersPublic); !ok {
		err = errors.New("unexptected type of peers public key")
	}
	return
}

func (ecdh25519) ComputeSecret(private crypto.PrivateKey, peersPublic crypto.PublicKey) (secret []byte) {
	var sec, pri, pub [32]byte
	if ok := checkType(&pri, private); !ok {
		panic("ecdh: unexpected type of private key")
	}
	if ok := checkType(&pub, peersPublic); !ok {
		panic("ecdh: unexpected type of peers public key")
	}

	curve25519.ScalarMult(&sec, &pri, &pub)

	secret = sec[:]
	return
}

func checkType(key *[32]byte, typeToCheck interface{}) (ok bool) {
	switch t := typeToCheck.(type) {
	case [32]byte:
		copy(key[:], t[:])
		ok = true
	case *[32]byte:
		copy(key[:], t[:])
		ok = true
	case []byte:
		if len(t) == 32 {
			copy(key[:], t)
			ok = true
		}
	case *[]byte:
		if len(*t) == 32 {
			copy(key[:], *t)
			ok = true
		}
	}
	return
}
