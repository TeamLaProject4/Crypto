package blockchain_test

import (
	"cryptomunt/blockchain"
	"testing"
)

func constructRewardTransaction() blockchain.Transaction {
	rewardTx := new(blockchain.Transaction)
	rewardTx.ReceiverPublicKey = "alice"
	rewardTx.Amount = 20
	rewardTx.Type = blockchain.REWARD
	return *rewardTx
}

func TestWhenNewBlockchainCreatedThenFirstBlockIsGenesis(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	genesis := blockchain.CreateGenesisBlock(*new([]blockchain.Transaction))

	firstBlock := chain.Blocks[0]
	if !firstBlock.Equals(genesis) {
		t.Errorf("Expected '%+v', but got '%+v'", genesis, firstBlock)
	}
}

func TestWhenNewBlockHasCorrectPreviousHashThenPreviousHashValid(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
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
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	tx := constructTransaction()

	got := chain.IsTransactionCovered(tx)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// TODO remove test when exchange transaction type removed
func TestWhenTransactionIsExchangeTypeThenTransactionCovered(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	exchangeTx := new(blockchain.Transaction)
	exchangeTx.SenderPublicKey = "exchange"
	exchangeTx.ReceiverPublicKey = "alice"
	exchangeTx.Amount = 10
	exchangeTx.Type = blockchain.EXCHANGE

	got := chain.IsTransactionCovered(*exchangeTx)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

// TODO remove test when exchange transaction type removed
func TestWhenExchangeTransactionExecutedThenReceiverHasCorrectBalance(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	amount := 10

	exchangeTx := new(blockchain.Transaction)
	exchangeTx.SenderPublicKey = "exchange"
	exchangeTx.ReceiverPublicKey = "alice"
	exchangeTx.Amount = amount
	exchangeTx.Type = blockchain.EXCHANGE
	coveredTransactions := chain.GetCoveredTransactions([]blockchain.Transaction{*exchangeTx})
	chain.ExecuteTransactions(coveredTransactions)

	got := chain.GetAccountBalance("alice")
	want := amount - blockchain.CalculateFee(amount)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenTransferTransactionExecutedThenReceiverHasCorrectBalance(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	chain.AccountModel.Balances["alice"] = 10
	amount := 5

	transferTx := new(blockchain.Transaction)
	transferTx.SenderPublicKey = "alice"
	transferTx.ReceiverPublicKey = "bob"
	transferTx.Amount = amount
	transferTx.Type = blockchain.TRANSFER
	coveredTransactions := chain.GetCoveredTransactions([]blockchain.Transaction{*transferTx})
	chain.ExecuteTransactions(coveredTransactions)

	got := chain.GetAccountBalance("bob")
	want := amount - blockchain.CalculateFee(amount)
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenRewardTransactionExecutedThenReceiverHasCorrectBalance(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	rewardTx := constructRewardTransaction()
	coveredTransactions := chain.GetCoveredTransactions([]blockchain.Transaction{rewardTx})
	chain.ExecuteTransactions(coveredTransactions)

	got := chain.GetAccountBalance("alice")
	want := rewardTx.Amount
	if want != got {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}

func TestWhenNoRewardTransactionInBlockThenBlockIsInvalid(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	block := constructBlock()

	got := chain.IsBlockRewardTransactionValid(block)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenMultipleRewardTransactionsInBlockThenBlockIsInvalid(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	block := constructBlock()
	rewardTxs := []blockchain.Transaction{
		constructRewardTransaction(),
		constructRewardTransaction(),
	}
	block.Transactions = append(block.Transactions, rewardTxs...)

	got := chain.IsBlockRewardTransactionValid(block)
	want := false
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestWhenRewardTransactionHasForgerAsReceiverAndCorrectAmountThenBlockIsValid(t *testing.T) {
	chain := blockchain.CreateBlockchain(*new([]blockchain.Transaction))
	block := constructBlock()
	block.Transactions[0].Amount = 11
	rewardTx := constructRewardTransaction()
	rewardTx.Amount = 1
	block.Forger = rewardTx.ReceiverPublicKey
	block.Transactions = append(block.Transactions, rewardTx)

	got := chain.IsBlockRewardTransactionValid(block)
	want := true
	if want != got {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}
