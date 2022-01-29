package main

import (
	"bufio"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"fmt"
	"os"
)

func main() {
	utils.InitLogger()

	node := networking.CreateCryptoNode()

	//node.WriteToTopic("STRING", networking.BLOCKCHAIN)

	//api := startRestApi()
	//listens to api calls
	//when api call x then write to topic
	//
	go tempWriteToTopic(node)

	//infinite loop
	select {}
}

func tempWriteToTopic(node networking.CryptoNode) {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		fmt.Println("writing to topic..." + sendData)

		//api -> key/mnemonic -> wallet
		wallet := blockchain.CreateWallet()
		transaction := wallet.CreateTransaction("jeroen", 20, blockchain.TRANSFER)

		node.WriteToTopic(transaction.ToJson(), networking.TRANSACTION)
	}
}
