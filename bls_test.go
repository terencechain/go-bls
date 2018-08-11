package bls

import (
	"testing"
	"github.com/Nik-U/pbc"

)

func TestOneTimeSignVerify(t *testing.T) {

	// Setting up new paring
	params := pbc.GenerateA(160, 512)
	paring := params.NewPairing()
	g1 := paring.NewG1().Rand()
}