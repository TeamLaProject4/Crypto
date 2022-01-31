package main

import (
	"bufio"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"flag"
	"fmt"
	"os"
)

func main2() {
	utils.InitLogger()

	config := parseFlags()
	if config.NodesToBoot != 0 {
		nodeFactory(config)
	} else {
		go startNode(config)
	}

	//infinite loop
	select {}
}

//func main() {
//	utils.InitLogger()
//	wallet.GenerateMnemonic()
//	key, err := utils.ReadEDCSAFromtFile()
//	if err != nil {
//		return
//	}
//	utils.Logger.Info(key)
//}

func parseFlags() networking.Config {
	config := networking.Config{}
	flag.Var(&config.BootNodes, "peer", "Peer multiaddress for peer discovery")
	flag.IntVar(&config.NodesToBoot, "amount", 0, "amount of nodesToBoot using a factory, 0 is for making a bootnode")
	flag.Parse()
	return config
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
