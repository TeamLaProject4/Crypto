package integration_test

import (
	"cryptomunt/blockchain"
	"fmt"
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
	transaction4 := constructTransaction("martijn", "henk", 32)
	return []blockchain.Transaction{transaction1, transaction2, transaction3, transaction4}
}
func constructBlock(prevHash string) blockchain.Block {
	block := new(blockchain.Block)
	block.Transactions = constructTransactions()
	block.PreviousHash = prevHash
	block.Forger = "forger"
	block.Height = 1
	return *block
}

func TestWhenSetBalancesFromBlockchainThenBalanceHasCorrectAmount(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	block1 := constructBlock(chain.LatestPreviousHash())
	chain.Blocks = []blockchain.Block{block1}

	//create balances using the transactions in the blockchain
	chain.AccountModel.SetBalancesFromBlockChain(chain)
	got := chain.AccountModel.Balances["jeroen"]
	want := 30

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhen(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	block1 := constructBlock(chain.LatestPreviousHash())
	block2 := constructBlock(chain.LatestPreviousHash())
	block3 := constructBlock(chain.LatestPreviousHash())
	chain.Blocks = []blockchain.Block{block1, block2, block3}

	//create balances using the transactions in the blockchain
	chain.AccountModel.SetBalancesFromBlockChain(chain)
	got := 90
	want := chain.AccountModel.Balances["jeroen"]

	transactions := chain.GetAllAccountTransactions("lars")
	fmt.Println("transactions from henk", transactions)

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
