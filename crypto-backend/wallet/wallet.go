package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"cryptomunt/utils"
	"encoding/hex"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Key ecdsa.PrivateKey
}

func CreateWalletFromKeyFile() Wallet {
	//pemEncodedKey := ReadKeyFromFile(PRIVATE_KEY_PATH)
	//privateKey := DecodePrivateKey(pemEncodedKey)
	//return Wallet{Key: *privateKey}
	return Wallet{}
}

func CreateWallet(key ecdsa.PrivateKey) Wallet {
	return Wallet{Key: key}
}

func (wallet *Wallet) Sign(data string) string {
	message := []byte(data)
	hashed := ethCrypto.Keccak256Hash(message)
	signature, err := ethCrypto.Sign(hashed.Bytes(), &wallet.Key)

	//Sign hash of message because only small messages can be signed
	//hashed := sha256.Sum256(message)

	//signature, err := ecdsa.SignASN1(cryptoRand.Reader, &wallet.Key, hashed[:])
	if err != nil {
		utils.Logger.Errorf("Error from signing: %s\n", err)
		return ""
	}

	return hex.EncodeToString(signature)
}

func (wallet *Wallet) GetPublicKeyHex() string {
	pubKey := wallet.Key.PublicKey
	pubKeyBytes := elliptic.Marshal(pubKey, pubKey.X, pubKey.Y)

	return hex.EncodeToString(pubKeyBytes)
}

func IsValidSignature(data string, signature string, publicKey string) bool {
	message := []byte(data)
	hashed := ethCrypto.Keccak256Hash(message)
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		utils.Logger.Error(err)
		return false
	}

	//new
	sigPublicKey, err := ethCrypto.Ecrecover(hashed.Bytes(), sigBytes)
	if err != nil {
		utils.Logger.Error(err)
		return false
	}
	matches2 := bytes.Equal(sigPublicKey, []byte(publicKey))
	utils.Logger.Info("VALID ? ", matches2)

	return matches2
}
