package main

import (
	"cryptomunt/api"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/structs"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"encoding/hex"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"strconv"
)

func nodeFactory(config networking.Config) {
	for i := 0; i < config.NodesToBoot; i++ {
		go startNode(config)
	}
}

func startNode(config networking.Config) {
	apiRequest := make(chan structs.ApiCallMessage)
	apiResponse := make(chan structs.ApiCallMessage)

	node := networking.CreateAndInitCryptoNode(config, apiRequest, apiResponse)

	utils.Logger.Info(len(node.Blockchain.Blocks))
	go node.HandleApiCalls(apiRequest, apiResponse)
	go api.StartApi(node, apiRequest, apiResponse)
	select {}

}

func createWallets(node networking.CryptoNode) {
	publicKeys := make([]string, 10)

	for i := 0; i < 10; i++ {
		mnemonic := wallet.GenerateMnemonic()

		key := wallet.NewMasterKey(mnemonic)
		ecdsaKey := wallet.ConvertBip32ToECDSA(key)

		pubKeyBytes := ethCrypto.FromECDSAPub(&ecdsaKey.PublicKey)
		pubKeyHex := hex.EncodeToString(pubKeyBytes)

		publicKeys = append(publicKeys, pubKeyHex)
	}

	utils.Logger.Info(publicKeys)
}

func transactionsFactory(amountOfTransactions int, node networking.CryptoNode) {
	//founderwallet
	//gives coins to n wallets

	for i := 0; i < amountOfTransactions; i++ {
		node.Wallet.CreateTransaction("public-key-"+strconv.Itoa(i), rand.Intn(1000), blockchain.TRANSFER)
	}
}
