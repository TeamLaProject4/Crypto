package blockchain_test

import (
	"cryptomunt/blockchain"
	"testing"
)

func TestWhenBigAmountGivenWhenFeeIsCorrect(t *testing.T) {
	amount := 1000
	fee := 4

	got := blockchain.CalculateFee(amount)
	want := fee
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenSmallestAmountGivenWhenFeeIsCorrect(t *testing.T) {
	amount := 1
	fee := 1

	got := blockchain.CalculateFee(amount)
	want := fee
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenBigAmountIncludingFeeGivenThenInitialAmountIsCorrect(t *testing.T) {
	amountIncludingFee := 1004
	initialAmount := 1000

	got := blockchain.CalculateInitialAmount(amountIncludingFee)
	want := initialAmount
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenSmallestAmountIncludingFeeGivenThenInitialAmountIsCorrect(t *testing.T) {
	amountIncludingFee := 2
	initialAmount := 1

	got := blockchain.CalculateInitialAmount(amountIncludingFee)
	want := initialAmount
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenRewardTransactionCreatedForgerIsSetAsReceiver(t *testing.T) {
	tx1 := constructTransaction()
	transactions := []blockchain.Transaction{tx1}
	forger := "forger"
	rewardTx := blockchain.CreateRewardTransaction(forger, transactions)

	got := rewardTx.ReceiverPublicKey
	want := forger
	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestWhenRewardTransactionCreatedRewardAmountIsCorrect(t *testing.T) {
	tx1 := constructTransaction()
	tx1.Amount = 11
	tx2 := constructTransaction()
	tx2.Amount = 11
	tx3 := constructTransaction()
	tx3.Amount = 2008
	transactions := []blockchain.Transaction{tx1, tx2, tx3}
	rewardTx := blockchain.CreateRewardTransaction("forger", transactions)

	got := rewardTx.Amount
	want := (1 + 1 + 8)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
