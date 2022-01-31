package wallet

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"cryptomunt/utils"
	"encoding/hex"
)

type Wallet struct {
	key ecdsa.PrivateKey
}

func CreateWalletFromKeyFile() Wallet {
	key, err := ReadKeyFromFile("./keys/privatekey.bytes")
	if err != nil {
		utils.Logger.Warn(err)
		return Wallet{}
	}

	return Wallet{key: *key}
}

func CreateWallet(key ecdsa.PrivateKey) Wallet {
	return Wallet{key: key}
}

func (wallet *Wallet) Sign(data string) string {
	message := []byte(data)
	//Sign hash of message because only small messages can be signed
	hashed := sha256.Sum256(message)

	signature, err := ecdsa.SignASN1(cryptoRand.Reader, &wallet.key, hashed[:])
	if err != nil {
		utils.Logger.Errorf("Error from signing: %s\n", err)
		return ""
	}

	return hex.EncodeToString(signature)
}

func (wallet *Wallet) GetPublicKeyHex() string {
	pubkey := wallet.key.PublicKey
	pubKeyBytes := elliptic.Marshal(pubkey, pubkey.X, pubkey.Y)

	return hex.EncodeToString(pubKeyBytes)
}

func IsValidSignature(data string, signature string, publicKey string) bool {
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	message := []byte(data)
	hashed := sha256.Sum256(message)

	pubKey, _ := utils.GetPublicKeyFromHex(publicKey)
	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], sig)

	if err != nil {
		return false
	}

	return true
}
