package main

import (
	"cryptomunt/api"
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"math/rand"
	"strconv"
)

func nodeFactory(config Config) {
	for i := 0; i < config.nodesToBoot; i++ {
		go startNode(config.BootNodes)
	}
}

func startNode(bootnodes networking.AddrList) {
	node := networking.CreateAndInitCryptoNode(bootnodes)
	utils.Logger.Info("pubkey hex: ", node.Wallet.GetPublicKeyHex())
	go api.StartApi(node)
}

func transactionsFactory(amountOfTransactions int, node networking.CryptoNode) {
	for i := 0; i < amountOfTransactions; i++ {
		node.Wallet.CreateTransaction("public-key-"+strconv.Itoa(i), rand.Intn(1000), blockchain.TRANSFER)
	}
}
