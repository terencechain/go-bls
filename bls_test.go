package bls

import (
	"github.com/Nik-U/pbc"
	"testing"

	"crypto/sha256"
)

type data struct {
	msg string
	sig []byte
}

func TestOneTimeVerification(t *testing.T) {

	// Setting up new pairing
	params := pbc.GenerateA(160, 512)
	pairing := params.NewPairing()
	generator := pairing.NewG2().Rand()

	privateKey1 := pairing.NewZr().Rand()
	pubKey1 := pairing.NewG2().PowZn(generator, privateKey1)

	t.Log(privateKey1.Bytes()) // 20 bytes
	t.Log(pubKey1.Bytes())     // 128 bytes

	// Sign a message, hash it to h
	message := "today is a good day"
	h := pairing.NewG1().SetFromStringHash(message, sha256.New())

	sig1 := pairing.NewG2().PowZn(h, privateKey1)

	// To verify, we check that e(h,g^x)=e(sig,g)
	if !pairing.NewGT().Pair(h, pubKey1).Equals(pairing.NewGT().Pair(sig1, generator)) {
		t.Error("Signature verificaiton failed")
	}
}
