package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptoRand "crypto/rand"
	"crypto/sha256"
	"cryptomunt/utils"
	"encoding/hex"
)

type Wallet struct {
	key ecdsa.PrivateKey
}

func CreateWalletFromKeyFile() Wallet {
	pemEncodedKey := ReadKeyFromFile(PRIVATE_KEY_PATH)
	privateKey := DecodePrivateKey(pemEncodedKey)
	return Wallet{key: *privateKey}
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
	pubKey := wallet.key.PublicKey
	pubKeyBytes := elliptic.Marshal(pubKey, pubKey.X, pubKey.Y)

	return hex.EncodeToString(pubKeyBytes)
}

func IsValidSignature(data string, signature string, publicKey string) bool {
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	message := []byte(data)
	hashed := sha256.Sum256(message)

	pubKeyBytes, _ := hex.DecodeString(publicKey)
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKeyBytes)
	pubKey := new(ecdsa.PublicKey)
	pubKey.Curve = elliptic.P256()
	pubKey.X = x
	pubKey.Y = y

	valid := ecdsa.VerifyASN1(pubKey, hashed[:], sig)

	return valid
}
