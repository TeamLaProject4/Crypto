package blockchain_test

import (
	blockchain "cryptomunt/blockchain"
	"testing"
)

func TestWhenTransactionNotAddedThenTransactionNotInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction := constructTransaction()

	got := pool.IsTransactionInPool(transaction)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenTransactionAddedThenTransactionInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction := constructTransaction()
	pool.AddTransaction(transaction)

	got := pool.IsTransactionInPool(transaction)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenMultipleTransactionsAddedThenTransactionsInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction1 := constructTransaction()
	transaction1.Id = "new signature"
	transaction2 := constructTransaction()

	transactions := []blockchain.Transaction{transaction2, transaction1}
	pool.AddTransactions(transactions)

	got := pool.IsTransactionInPool(transaction1)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}

	got = pool.IsTransactionInPool(transaction2)
	want = true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenTransactionAddedTwiceThenOneTransactionInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction1 := constructTransaction()
	transaction2 := constructTransaction()

	pool.AddTransaction(transaction1)
	pool.AddTransaction(transaction2)

	got := pool.GetTransactionsLength()
	want := 1
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestGivenPoolWithTransactionWhenTransactionRemovedTransactionThenTransactionNotInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction1 := constructTransaction()

	pool.AddTransaction(transaction1)
	pool.RemoveTransaction(transaction1)

	got := pool.IsTransactionInPool(transaction1)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestGivenPoolWithTransactionsWhenTransactionsRemovedFromPoolThenTransactionsNotInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction1 := constructTransaction()
	transaction1.Id = "new signature"
	transaction2 := constructTransaction()

	transactions := []blockchain.Transaction{transaction2, transaction1}
	pool.AddTransactions(transactions)
	pool.RemoveTransactions(transactions)

	got := pool.IsTransactionInPool(transaction1)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}

	got = pool.IsTransactionInPool(transaction2)
	want = false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestGivenPoolWithThreeTransactionsWhenLastTwoTransactionsRemovedThenFirstTransactionInPool(t *testing.T) {
	pool := blockchain.CreateMemoryPool()
	transaction1 := constructTransaction()
	transaction1.Id = "new signature"
	transaction2 := constructTransaction()
	transaction3 := constructTransaction()
	transaction3.Id = "new signature 3"

	transactions := []blockchain.Transaction{transaction2, transaction1, transaction3}
	pool.AddTransactions(transactions)
	pool.RemoveTransaction(transaction1)
	pool.RemoveTransaction(transaction2)

	got := pool.IsTransactionInPool(transaction3)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}

}
