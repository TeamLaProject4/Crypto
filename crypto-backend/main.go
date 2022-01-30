package main

import (
	"bufio"
	"cryptomunt/api"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"fmt"
	"os"
)

func main() {
	utils.InitLogger()

	networking.CreateAndInitCryptoNode()
	go api.StartApi()

	go tempWriteToTopic(networking.Node)

	//go startRestApi()
	//when api call x then write to topic

	//infinite loop
	select {}
}

func tempWriteToTopic(node networking.CryptoNode) {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin", sendData)
			panic(err)
		}

		fmt.Println("performing actions...")

		//reset blockchain
		node.Blockchain.Blocks = *new([]blockchain.Block)
		node.Blockchain.AccountModel = new(blockchain.AccountModel)
		node.Blockchain.ProofOfStake = new(proofOfStake.ProofOfStake)
		//get from network
		node.SetBlockchainUsingNetwork()
		utils.Logger.Info("blockchain", node.Blockchain.AccountModel)
		utils.Logger.Info("blockchain", node.Blockchain.ProofOfStake)
		//api -> key/mnemonic -> wallet
		//wallet := blockchain.CreateWallet()
		//transaction := wallet.CreateTransaction("jeroen", 20, blockchain.TRANSFER)
		//
		//node.WriteToTopic(transaction.ToJson(), networking.TRANSACTION)
	}
}
