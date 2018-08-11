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

	// Public key distribution
	c1 := make(chan []byte)
	c2 := make(chan []byte)
	c3 := make(chan []byte)

	// Done channel for blocking purpose
	done := make(chan bool)


	// In the end, waiting for all channels communication to finish
	<-done
	<-done
}

