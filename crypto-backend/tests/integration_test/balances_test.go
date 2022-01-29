package integration_test

import (
	"cryptomunt/blockchain"
	"testing"
)

func constructTransaction(pk string, rk string, amount int) blockchain.Transaction {
	return blockchain.Transaction{
		SenderPublicKey:   pk,
		ReceiverPublicKey: rk,
		Amount:            amount,
		Type:              blockchain.TRANSFER,
	}
}
func constructTransactions() []blockchain.Transaction {
	transaction1 := constructTransaction("lars", "jeroen", 20)
	transaction2 := constructTransaction("johan", "jeroen", 10)
	transaction3 := constructTransaction("martijn", "lars", 32)
	return []blockchain.Transaction{transaction1, transaction2, transaction3}
}
func constructBlock(prev_hash string) blockchain.Block {
	block := new(blockchain.Block)
	block.Transactions = constructTransactions()
	block.PreviousHash = prev_hash
	block.Forger = "forger"
	block.Height = 1
	return *block
}

func TestWhenSetBalancesFromBlockchainThenBalanceHasCorrectAmount(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	chain.Blocks = []blockchain.Block{constructBlock(chain.LatestPreviousHash())}

	got := 30

	//create balances using the transactions in the blockchain
	chain.AccountModel.SetBalancesFromBlockChain(chain)
	want := chain.AccountModel.Balances["jeroen"]

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
