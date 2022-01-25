package blockchain

import (
	"encoding/json"
	"reflect"
)

type Block struct {
	Transactions []Transaction
	PreviousHash string
	Forger       string
	Height       int
	Timestamp    int64
	Signature    string
}

//var block = new(Block)

//func NewBlock(newBlock Block) {
//	if newBlock.Timestamp == 0 {
//		newBlock.Timestamp = time.Now().Unix()
//	}
//	block = &newBlock
//}

func blockEquals(blokA Block, blockB Block) bool {
	return blokA.Signature == blockB.Signature &&
		(reflect.DeepEqual(blokA.PreviousHash, blockB.PreviousHash) && reflect.DeepEqual(blokA.Forger, blockB.Forger) && reflect.DeepEqual(blokA.Height, blockB.Height) && blokA.Timestamp == blockB.Timestamp)
}

func blockToJson(block Block) string {
	blockJson, err := json.Marshal(block)
	if err != nil {
		panic("ERROR")
	}
	return string(blockJson)
}

func getBlockPayload(block Block) string {
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

func signBlock(block Block, signature string) {
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
