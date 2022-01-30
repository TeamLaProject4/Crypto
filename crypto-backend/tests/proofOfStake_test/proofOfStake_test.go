package proofofstaketest_test

import (
	"bytes"
	"cryptomunt/proofOfStake"
	"math/rand"
	"testing"
)

func generateRandomString(length int) string {
	var letters bytes.Buffer
	for i := 0; i <= length; i++ {
		randomChar := 'a' + rune(rand.Intn(26))
		letters.WriteString(string(randomChar))
	}
	return letters.String()
}

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

func TestGivenEqualStakeWhenForgerPickedThenItRepresentsStake(t *testing.T) {
	account_1 := "barrie"
	account_2 := "sjonnie"
	pos := proofOfStake.NewProofOfStake()
	pos.UpdateStake(account_1, 100)
	pos.UpdateStake(account_2, 100)
	wins := map[string]int{
		account_1: 0,
		account_2: 0,
	}

	for i := 0; i < 100; i++ {
		seed := generateRandomString(16)
		forger := pos.PickForger(seed)
		if forger == account_1 || forger == account_2 {
			wins[forger] += 1
		}
	}

	got := 40 < wins[account_1] && wins[account_1] < 60 && 40 < wins[account_2] && wins[account_2] < 60
	want := true
	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestGivenUnequalStakeWhenForgerPickedThenItRepresentsStake(t *testing.T) {
	account_1 := "barrie"
	account_2 := "sjonnie"
	pos := proofOfStake.NewProofOfStake()
	pos.UpdateStake(account_1, 10)
	pos.UpdateStake(account_2, 90)
	wins := map[string]int{
		account_1: 0,
		account_2: 0,
	}

	for i := 0; i < 100; i++ {
		seed := generateRandomString(16)
		forger := pos.PickForger(seed)
		if forger == account_1 || forger == account_2 {
			wins[forger] += 1
		}
	}

	got := 0 < wins[account_1] && wins[account_1] < 20 && 80 < wins[account_2] && wins[account_2] < 100
	want := true
	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
