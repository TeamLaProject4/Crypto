package main

import (
	"cryptomunt/api"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/structs"
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
	apiResponse := make(chan string)

	node := networking.CreateAndInitCryptoNode(config, apiRequest, apiResponse)
	go api.StartApi(node, apiRequest, apiResponse)
	select {}

}

func innitialCoinOffering() {

}

func transactionsFactory(amountOfTransactions int, node networking.CryptoNode) {
	//founderwallet
	//gives coins to n wallets

	for i := 0; i < amountOfTransactions; i++ {
		node.Wallet.CreateTransaction("public-key-"+strconv.Itoa(i), rand.Intn(1000), blockchain.TRANSFER)
	}
}
