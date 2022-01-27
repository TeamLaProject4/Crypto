package blockchain_test

import (
	blockchain "cryptomunt/blockchain"
	"testing"
)

func TestWhenNewBlockchainCreatedThenFirstBlockIsGenesis(t *testing.T) {
	chain := blockchain.NewBlockchain()
	genesis := blockchain.CreateGenesisBlock()

	firstBlock := chain.Blocks[0]
	if !firstBlock.Equals(genesis) {
		t.Errorf("Expected '%+v', but got '%+v'", genesis, firstBlock)
	}
}