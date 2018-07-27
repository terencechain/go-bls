package bls

import (
	"testing"
	"golang.org/x/crypto/blake2s"
)

func TestOneTimeSignVerify(t *testing.T) {

	// Assume we are using type A param
	// See: https://crypto.stanford.edu/pbc/manual/ch08s03.html
	params := GenParamsTypeA(360, 1024)
	paring := GenPairing(params)
	system , _ := GenSystem(paring)
	pubKey, privKey, err := GenKeys(system)
	if err != nil {
		t.Fatal(err)
	}

	// Hash the message, sign the hash to get signature
	hash, err := blake2s.New256([]byte("I love bls!"))
	if err != nil {
		t.Fatal(err)
	}
	var hashByte [32]byte
	copy(hashByte[:], hash.Sum(nil))
	sig := Sign(hashByte, privKey)

	// Verify the output of signature and the hashed message matches public key
	t.Log(Verify(sig, hashByte, pubKey))
}