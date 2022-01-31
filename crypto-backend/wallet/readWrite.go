package wallet

import (
	"cryptomunt/utils"
)

var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.pem"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.pem"

func WriteKeyToFile(path string, pemEncodedKey string) {
	utils.WriteFile(path, pemEncodedKey)
}

func ReadKeyFromFile(path string) string {
	pemEncodedKey := utils.ReadFileBytes(path)
	return string(pemEncodedKey)
}
