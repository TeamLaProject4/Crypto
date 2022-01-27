package blockchain

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
	Signature    string `json:"signature"`
}

func CreateBlock(newBlock Block) Block {
	if newBlock.Timestamp == 0 {
		newBlock.Timestamp = time.Now().Unix()
	}
	return newBlock
}

func CreateGenesisBlock() Block {
	genesis := new(Block)
	genesis.PreviousHash = "genesis_hash"
	genesis.Forger = "genesis_forger"
	genesis.Height = 0
	genesis.Timestamp = 0
	return *genesis
}

func (block *Block) Equals(blockToCompareTo Block) bool {
	return block.Signature == blockToCompareTo.Signature &&
		(reflect.DeepEqual(block.PreviousHash, blockToCompareTo.PreviousHash) && reflect.DeepEqual(block.Forger, blockToCompareTo.Forger) && reflect.DeepEqual(block.Height, blockToCompareTo.Height) && block.Timestamp == blockToCompareTo.Timestamp)
}

func (block *Block) ToJson() string {
	blockJson, err := json.Marshal(block)
	if err != nil {
		panic("ERROR")
	}
	return string(blockJson)
}

func (block *Block) Payload() string {
	tempSignature := block.Signature
	block.Signature = ""
	blockJson := block.ToJson()

	block.Signature = tempSignature
	return blockJson
}

func (block *Block) Sign(signature string) {
	block.Signature = signature
}
