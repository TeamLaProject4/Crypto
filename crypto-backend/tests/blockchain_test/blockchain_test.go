package blockchain_test

import (
	blockchain "cryptomunt/blockchain"
	"testing"
)

func TestWhenNewBlockchainCreatedThenFirstBlockIsGenesis(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	genesis := blockchain.CreateGenesisBlock()

	firstBlock := chain.Blocks[0]
	if !firstBlock.Equals(genesis) {
		t.Errorf("Expected '%+v', but got '%+v'", genesis, firstBlock)
	}
}

func TestWhenNewBlockHasCorrectPreviousHashThenPreviousHashValid(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	prev_hash := chain.LatestPreviousHash()
	tx := constructTransaction()
	block := new(blockchain.Block)
	block.Transactions = []blockchain.Transaction{tx}
	block.PreviousHash = prev_hash
	block.Forger = "forger"
	block.Height = 1

	got := chain.IsValidPreviousBlockHash(*block)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenNewBlockAddedThenLatestBlockHeightIsCorrect(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	prev_hash := chain.LatestPreviousHash()
	tx := constructTransaction()
	block := new(blockchain.Block)
	block.Transactions = []blockchain.Transaction{tx}
	block.PreviousHash = prev_hash
	block.Forger = "forger"
	block.Height = 1
	chain.AddBlock(*block)

	got := chain.LatestBlockHeight()
	want := 1
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestGivenNewBlockchainWhenBlockAddedThenSecondBlockIsBlock(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	prev_hash := chain.LatestPreviousHash()
	tx := constructTransaction()
	block := new(blockchain.Block)
	block.Transactions = []blockchain.Transaction{tx}
	block.PreviousHash = prev_hash
	block.Forger = "forger"
	block.Height = 1
	chain.AddBlock(*block)

	got := chain.Blocks[1]
	want := block
	if !want.Equals(got) {
		t.Errorf("Expected '%+v', but got '%+v'", want, got)
	}
}

func TestWhenThirdBlockHasLowerBlockHeightThenBlockHeightInvalid(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	tx := constructTransaction()

	prev_hash_1 := chain.LatestPreviousHash()
	block_1 := new(blockchain.Block)
	block_1.Transactions = []blockchain.Transaction{tx}
	block_1.PreviousHash = prev_hash_1
	block_1.Forger = "forger"
	block_1.Height = 1
	chain.AddBlock(*block_1)

	prev_hash_2 := chain.LatestPreviousHash()
	block_2 := new(blockchain.Block)
	block_2.Transactions = []blockchain.Transaction{tx}
	block_2.PreviousHash = prev_hash_2
	block_2.Forger = "forger"
	block_2.Height = 1

	got := chain.IsValidBlockHeight(*block_2)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenThirdBlockHasHigherBlockHeightThenBlockHeightInvalid(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	tx := constructTransaction()

	prev_hash_1 := chain.LatestPreviousHash()
	block_1 := new(blockchain.Block)
	block_1.Transactions = []blockchain.Transaction{tx}
	block_1.PreviousHash = prev_hash_1
	block_1.Forger = "forger"
	block_1.Height = 1
	chain.AddBlock(*block_1)

	prev_hash_2 := chain.LatestPreviousHash()
	block_2 := new(blockchain.Block)
	block_2.Transactions = []blockchain.Transaction{tx}
	block_2.PreviousHash = prev_hash_2
	block_2.Forger = "forger"
	block_2.Height = 3

	got := chain.IsValidBlockHeight(*block_2)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenThirdBlockHasCorrectBlockHeightThenBlockHeightInvalid(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	tx := constructTransaction()

	prev_hash_1 := chain.LatestPreviousHash()
	block_1 := new(blockchain.Block)
	block_1.Transactions = []blockchain.Transaction{tx}
	block_1.PreviousHash = prev_hash_1
	block_1.Forger = "forger"
	block_1.Height = 1
	chain.AddBlock(*block_1)

	prev_hash_2 := chain.LatestPreviousHash()
	block_2 := new(blockchain.Block)
	block_2.Transactions = []blockchain.Transaction{tx}
	block_2.PreviousHash = prev_hash_2
	block_2.Forger = "forger"
	block_2.Height = 2

	got := chain.IsValidBlockHeight(*block_2)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestGivenNewBlockchainWhenRandomTransactionCreatedThenTransactionNotCovered(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	tx := constructTransaction()

	got := chain.IsTransactionCovered(tx)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// TODO remove test when exchange transaction type removed
func TestWhenTransactionIsExchangeTypeThenTransactionCovered(t *testing.T) {
	chain := blockchain.CreateBlockchain()
	exchangeTx := new(blockchain.Transaction)
	exchangeTx.SenderPublicKey = "exchange"
	exchangeTx.ReceiverPublicKey = "alice"
	exchangeTx.Amount = 10
	exchangeTx.TxType = blockchain.EXCHANGE

	got := chain.IsTransactionCovered(*exchangeTx)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// TODO remove test when exchange transaction type removed
func TestWhenExchangeTransactionExecutedThenReceiverHasCorrectBalance(t *testing.T) {
	chain := blockchain.CreateBlockchain()

	exchangeTx := new(blockchain.Transaction)
	exchangeTx.SenderPublicKey = "exchange"
	exchangeTx.ReceiverPublicKey = "alice"
	exchangeTx.Amount = 10
	exchangeTx.TxType = blockchain.EXCHANGE
	coveredTransactions := chain.GetCoveredTransactions([]blockchain.Transaction{*exchangeTx})
	chain.ExecuteTransactions(coveredTransactions)

	got := chain.GetAccountBalance("alice")
	want := 10
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenTransferTransactionExecutedThenReceiverHasCorrectBalance(t *testing.T) {
	chain := blockchain.CreateBlockchain()

	exchangeTx := new(blockchain.Transaction)
	exchangeTx.SenderPublicKey = "exchange"
	exchangeTx.ReceiverPublicKey = "alice"
	exchangeTx.Amount = 10
	exchangeTx.TxType = blockchain.EXCHANGE
	coveredTransactions := chain.GetCoveredTransactions([]blockchain.Transaction{*exchangeTx})
	chain.ExecuteTransactions(coveredTransactions)

	transferTx := new(blockchain.Transaction)
	transferTx.SenderPublicKey = "alice"
	transferTx.ReceiverPublicKey = "bob"
	transferTx.Amount = 5
	transferTx.TxType = blockchain.TRANSFER
	coveredTransactions = chain.GetCoveredTransactions([]blockchain.Transaction{*transferTx})
	chain.ExecuteTransactions(coveredTransactions)

	got := chain.GetAccountBalance("bob")
	want := 5
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
