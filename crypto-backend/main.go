package main

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"fmt"
)

func main() {
	//testTransaction := blockchain.Transaction{
	//	SenderPublicKey:   "",
	//	ReceiverPublicKey: "",
	//	Amount:            0,
	//	TransactionType:   0,
	//	Id:                "",
	//	Timestamp:         0,
	//	Signature:         "signature-value",
	//}
	//blockchain.NewTransaction(testTransaction)

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
	//proofOfStake.NewLot(proofOfStake.Lot{
	//	PublicKey:         "hello",
	//	Iteration:         0,
	//	PreviousBlockHash: "world",
	//})
	//
	//fmt.Println(proofOfStake.GetHash())

	proofOfStake.NewProofOfStake()
	//proofOfStake.SetGenesisNodeStake()
	//fmt.Println(proofOfStake.IsAccountInStakers("false"))

	//proofOfStake.AddAccountToStakers("hello")
	//proofOfStake.PrintStakers()

	//proofOfStake.UpdateStake("moi", 20)
	//proofOfStake.PrintStakers()
	//proofOfStake.UpdateStake("moi", 5)
	//proofOfStake.PrintStakers()
	//fmt.Println(proofOfStake.GetStake("moi"))

	//proofOfStake.UpdateStake("tammo", 3000)
	//proofOfStake.UpdateStake("moi", 1)
	//proofOfStake.UpdateStake("henk", 1)

	//lots := proofOfStake.GenerateLots("seed")
	//fmt.Println(lots)

	//fmt.Println(proofOfStake.PickForger("FF11asnF"))

	//blockchain.NewBlockchain()
	//fmt.Println(blockchain.GetBlockChain())
	//fmt.Println(blockchain.ToJson())

	//fmt.Println(blockchain.GetKeyPair())
	privKey := utils.ReadRsaKeyFile("../keys/wallet.rsa")
	fmt.Println(privKey.PublicKey)

}
