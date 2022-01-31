package wallet

import (
	"cryptomunt/utils"
	"crypto/ecdsa"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.pem"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.pem"

func WriteKeyToFile(path string, privateKey ecdsa.PrivateKey) {
	bytes := ethCrypto.FromECDSA(&privateKey)
	utils.WriteFile(path, string(bytes))
}

func ReadKeyFromFile(path string) (*ecdsa.PrivateKey, error) {
	privkey := utils.ReadFileBytes(path)
	return ethCrypto.ToECDSA(privkey)
}
