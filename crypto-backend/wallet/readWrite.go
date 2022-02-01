package wallet

import (
	"cryptomunt/utils"
)

var PRIVATE_KEY_PATH = "./keys/walletPrivateKey.txt"
var PUBLIC_KEY_PATH = "./keys/walletPublicKey.txt"

func WriteKeyToFile(path string, pemEncodedKey string) {
	utils.WriteFile(path, pemEncodedKey)
}

func ReadKeyFromFile(path string) string {
	pemEncodedKey := utils.ReadFileBytes(path)
	return string(pemEncodedKey)
}
