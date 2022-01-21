package main

import (
	"encoding/json"
	"reflect"
	"time"
)

type Block struct {
	Transactions []Transaction
	PreviousHash string
	Forger       string
	Height       int
	Timestamp    int64
	Signature    string
}

var block = new(Block)

func newBlock(newBlock Block) {
	if newBlock.Timestamp == 0 {
		newBlock.Timestamp = time.Now().Unix()
	}
	block = &newBlock
}

func blockEquals(blockCompare Block) bool {
	return block.Signature == blockCompare.Signature &&
		(reflect.DeepEqual(block.PreviousHash, blockCompare.PreviousHash) && reflect.DeepEqual(block.Forger, blockCompare.Forger) && reflect.DeepEqual(block.Height, blockCompare.Height) && block.Timestamp == blockCompare.Timestamp)
}

func blockToJson(block Block) string {
	blockJson, err := json.Marshal(block)
	if err != nil {
		panic("ERROR")
	}
	return string(blockJson)
}

func blockPayload() string {
	copy_block := Block{
		Transactions: block.Transactions,
		PreviousHash: block.PreviousHash,
		Forger:       block.Forger,
		Height:       block.Height,
		Timestamp:    block.Timestamp,
		Signature:    "",
	}
	return blockToJson(copy_block)
}

func signBlock(signature string) {
	block.Signature = signature
}

func createGenesisBlock() Block {
	genesis := new(Block)
	genesis.PreviousHash = "genesis_hash"
	genesis.Forger = "genesis_forger"
	genesis.Height = 0
	genesis.Timestamp = 0
	return *genesis
}
