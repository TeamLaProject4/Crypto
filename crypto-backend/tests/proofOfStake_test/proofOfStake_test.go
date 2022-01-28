package proofofstaketest_test

import (
	"cryptomunt/proofOfStake"
	"testing"
)

func TestWhenNewPOSThenAccountNotInStakers(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()

	got := pos.IsAccountInStakers("satoshi")
	want := false
	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenAccountAddedThenAccountInStakers(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()
	pos.AddAccountToStakers("barrie")

	got := pos.IsAccountInStakers("barrie")
	want := true
	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenAccountAddedThenStakeOfAccountIsZero(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()
	pos.AddAccountToStakers("barrie")

	got := pos.GetStake("barrie")
	want := 0
	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenStakeUpdatedThenStakeOfAccountUpdated(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()
	pos.AddAccountToStakers("barrie")
	pos.UpdateStake("barrie", 50)

	got := pos.GetStake("barrie")
	want := 50
	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenOneStakerAddedThenStakerIsForger(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()
	pos.UpdateStake("barrie", 100)

	got := pos.PickForger("prev_hash")
	want := "barrie"
	if got != want {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestWhenForgerPickedThenResultIsDeterministic(t *testing.T) {
	pos := proofOfStake.NewProofOfStake()
	pos.UpdateStake("barrie", 100)
	pos.UpdateStake("sjonnie", 100)
	seed := "zaadje"

	want := "barrie"
	for i := 0; i < 10; i++ {
		got := pos.PickForger(seed)
		if got != want {
			t.Errorf("Expected '%s', but got '%s'", want, got)
		}
	}
}