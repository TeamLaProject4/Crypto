package main

import (
	"cryptomunt/blockchain"
	"cryptomunt/proofOfStake"
	"fmt"
)

func main() {
	testTransaction := blockchain.Transaction{
		SenderPublicKey:   "",
		ReceiverPublicKey: "",
		Amount:            0,
		TransactionType:   0,
		Id:                "",
		Timestamp:         0,
		Signature:         "signature-value",
	}
	blockchain.NewTransaction(testTransaction)

	//testBlock := Block{
	//	Transactions: []Transaction{*transaction},
	//	PreviousHash: "",
	//	Forger:       "",
	//	Height:       0,
	//	Timestamp:    0,
	//	Signature:    "block-signature-value",
	//}
	//newBlock(testBlock)
	//
	//println(block)
	//println(blockToJson(*block))
	//println(blockPayload())
	//println(blockEquals(testBlock))

	//println(Hash())
	//println(TransactionToJson())
	//println(payload())n
	proofOfStake.NewLot(proofOfStake.Lot{
		PublicKey:         "hello",
		Iteration:         0,
		PreviousBlockHash: "world",
	})

	fmt.Println(proofOfStake.GetLotHash())

}

//1839806922502695369

//7314111797897617339
//7314111797897617339
