package bls

import (
	"testing"
	"github.com/Nik-U/pbc"

)

type data struct {
	msg string
	sig []byte
}

func TestOneTimeSignVerify(t *testing.T) {

	// Setting up new paring
	params := pbc.GenerateA(160, 512)
	paring := params.NewPairing()
	generator := paring.NewG1().Rand()

	// We should send the generator to the participants
	sharedParams := params.String()
	sharedGenerator := generator.Bytes()

	// Establish a channel to share data across
	messageChannel := make(chan *data)
}

