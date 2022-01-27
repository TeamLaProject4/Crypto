package main

import (
	"cryptomunt/utils"
)

func main() {
	utils.InitLogger()
	cryptoNode := CreateCryptoNode()
	//cryptoNode.WriteToTopic("NEED BLOCKCHAIN", BLOCKCHAIN)

	select {}
}
