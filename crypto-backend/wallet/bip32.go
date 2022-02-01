package wallet

import (
	"crypto/ecdsa"
	"cryptomunt/utils"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/jbenet/go-base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewMasterKey(mnemonic string) bip32.Key {
	seed := bip39.NewSeed(mnemonic, "secret")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	return *masterKey
}

func ConvertBip32ToECDSA(masterKey bip32.Key) ecdsa.PrivateKey {
	childKey, err := masterKey.NewChildKey(2147483648 + 44)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	decoded := base58.Decode(childKey.B58Serialize())
	privateKey := decoded[46:78]

	//Hex private key to ECDSA private key
	privateKeyECDSA, err := ethCrypto.ToECDSA(privateKey)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	return *privateKeyECDSA
}
