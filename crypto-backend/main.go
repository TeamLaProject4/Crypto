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

func main() {
	utils.InitLogger()

	//createWallets()
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
//	//init
//	utils.InitLogger()
//	config := parseFlags()
//	node := networking.CreateAndInitCryptoNode(config)
//	key, _ := ethCrypto.GenerateKey()
//	node.Wallet.Key = *key
//
//	//test
//	transaction := node.Wallet.CreateTransaction("receiverpubkey", 20, blockchain.TRANSFER)
//
//	payload := transaction.Payload()
//	signature := transaction.Signature
//	senderPublicKeyString := transaction.SenderPublicKeyString
//	utils.Logger.Info("valid signature", wallet.IsValidSignature(payload, signature, senderPublicKeyString))
//}

//func main() {
//	utils.InitLogger()
//
//	//creat transaction
//	key, _ := ethCrypto.GenerateKey()
//	publicKeyBytes := ethCrypto.FromECDSAPub(&key.PublicKey)
//
//	data := []byte("hello")
//	hash := ethCrypto.Keccak256Hash(data)
//	fmt.Println(hash.Hex())
//	signature, _ := ethCrypto.Sign(hash.Bytes(), key)
//
//	//verify transaction
//
//	sigPublicKey, err := ethCrypto.Ecrecover(hash.Bytes(), signature)
//
//	pubkeystring := string(ethCrypto.FromECDSAPub(&key.PublicKey))
//	fmt.Println(pubkeystring)
//	matches2 := bytes.Equal(sigPublicKey, publicKeyBytes)
//	fmt.Println(matches2) // true
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
//	fmt.Println(matches) // true
//
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
