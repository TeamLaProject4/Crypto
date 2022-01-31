package utils

import (
	"crypto/ecdsa"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

func WriteEDCSAToFile(privateKey *ecdsa.PrivateKey) {
	bytes := ethCrypto.FromECDSA(privateKey)
	WriteFile("./keys/privatekey.bytes", string(bytes))
}

func ReadEDCSAFromtFile() (*ecdsa.PrivateKey, error) {
	privkey := ReadFileBytes("./keys/privatekey.bytes")
	return ethCrypto.ToECDSA(privkey)
}
