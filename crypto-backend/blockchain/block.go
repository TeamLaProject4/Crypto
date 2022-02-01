package blockchain

import (
	"cryptomunt/utils"
	"encoding/json"
	"reflect"
	"time"
)

type Block struct {
	Transactions []Transaction `json:"Transactions"`
	PreviousHash string        `json:"previous_hash"`
	Forger       string        `json:"forger"`
	Height       int           `json:"height"`
	Timestamp    int64         `json:"timestamp"`
	Signature    string        `json:"signature"`
}

func CreateBlock(newBlock Block) Block {
	if newBlock.Timestamp == 0 {
		newBlock.Timestamp = time.Now().Unix()
	}
	return newBlock
}

func GetBlocksFromJson(jsonData string) []Block {
	var blocks []Block
	err := json.Unmarshal([]byte(jsonData), &blocks)
	if err != nil {
		utils.Logger.Error("unmarshal error ", err)
	}
	return blocks
}

func CreateGenesisBlock(transactions []Transaction) Block {
	genesis := new(Block)
	genesis.PreviousHash = "genesis_hash"
	genesis.Forger = "genesis_forger"
	genesis.Height = 0
	genesis.Timestamp = 0
	genesis.Transactions = transactions
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
