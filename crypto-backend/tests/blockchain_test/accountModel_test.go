package blockchain_test

import (
	"cryptomunt/blockchain"
	"testing"
)

func TestWhenAccountAddedThenAccountInBalances(t *testing.T) {
	publicKey := "alice"
	accountModel := new(blockchain.AccountModel)
	accountModel.Balances = make(map[string]int)
	accountModel.AddAccount(publicKey)

	got := accountModel.IsAccountInBalances(publicKey)
	want := true
	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenAccountAddedThenAccountStartsWithBalanceOfZero(t *testing.T) {
	publicKey := "alice"
	accountModel := new(blockchain.AccountModel)
	accountModel.Balances = make(map[string]int)
	accountModel.AddAccount(publicKey)

	got := accountModel.GetBalance(publicKey)
	want := 0
	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenBalanceUpdatedThenBalanceHasCorrectAmount(t *testing.T) {
	publicKey := "alice"
	accountModel := new(blockchain.AccountModel)
	accountModel.Balances = make(map[string]int)
	accountModel.UpdateBalance(publicKey, 20)
	accountModel.UpdateBalance(publicKey, -5)

	got := accountModel.GetBalance(publicKey)
	want := 15
	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
