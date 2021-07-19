package s

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEcdh25519_ComputeSecret(t *testing.T) {

	dh := X25519()

	secret := make([]byte, 32)
	for i := 0; i < 2; i++ {
		priAlice, pubAlice, err := dh.GenerateKey(nil)
		if err != nil {
			t.Fatalf("alice: key pair generation failed: %s", err)
		}

		priBob, pubBob, err := dh.GenerateKey(nil)
		if err != nil {
			t.Fatalf("alice: key pair generation failed: %s", err)
		}
		secAlice := dh.ComputeSecret(priAlice, pubBob)
		secBob := dh.ComputeSecret(priBob, pubAlice)

		{
			secAlice2 := dh.ComputeSecret(priAlice, pubBob)
			secBob2 := dh.ComputeSecret(priBob, pubAlice)
			secAlice3 := dh.ComputeSecret(priAlice, pubBob)
			secBob3 := dh.ComputeSecret(priBob, pubAlice)

			fmt.Println(secAlice2, secBob2, secAlice3, secBob3)
		}

		if !bytes.Equal(secAlice, secBob) {
			toStr := hex.EncodeToString
			t.Fatalf("DH failed: secrets are not equal:\nAlice got: %s\nBob   got: %s", toStr(secAlice), toStr(secBob))
		}
		if bytes.Equal(secret, secAlice) {
			t.Fatalf("DH generates the same secret all the time")
		}
		copy(secret, secAlice)
	}

}
