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

	//getBlockChainFromNetwork()
	node.Blockchain = blockchain.CreateBlockchain()

	//node.GetBlockChainFromNetwork()

	go tempWriteToTopic(node)

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
		fmt.Println("writing to topic...")
		node.WriteToTopic(sendData, networking.BLOCKCHAIN)
	}
}
