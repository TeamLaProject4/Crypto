package blockchain_test

import (
	blockchain "cryptomunt/blockchain"
	"testing"
	"time"
)

func constructBlock() blockchain.Block {
	tx := constructTransaction()
	return blockchain.Block{
		Transactions: []blockchain.Transaction{tx},
		PreviousHash: "prev_hash",
		Forger:       "forger",
		Height:       1,
	}
}

func TestWhenNewBlockConstructedThenSignatureIsEmpty(t *testing.T) {
	block := constructBlock()

	got := block.Signature
	want := ""
	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}

func TestWhenDuplicateBlocksConstructedThenBlocksAreEqual(t *testing.T) {
	tx := constructTransaction()
	timestamp := time.Now().Unix()

	block1 := blockchain.Block{
		Transactions: []blockchain.Transaction{tx},
		PreviousHash: "prev_hash",
		Forger:       "forger",
		Height:       1,
		Timestamp:	  timestamp,
	}
	block2 := blockchain.Block{
		Transactions: []blockchain.Transaction{tx},
		PreviousHash: "prev_hash",
		Forger:       "forger",
		Height:       1,
		Timestamp:	  timestamp,
	}

	if !block1.Equals(block2) {
		t.Errorf("Expected '%+v', but got '%+v'", block1, block2)
	}
}
