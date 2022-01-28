package proofofstaketest_test

import (
	pos "cryptomunt/proofOfStake"
	"testing"
)

func TestWhenLotHashedThenHasIsCorrect(t *testing.T) {
	lot := new(pos.Lot)
	lot.PublicKey = "alice"
	lot.Iteration = 1
	lot.PreviousBlockHash = "prev_hash"

	// "aliceprev_hash" -> sha256
	hash := "6c98859ec5fb8528e5af26658a8fdc2b4fd367771355beba5f96aec536357181"

	got := lot.Hash()
	want := hash
	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
