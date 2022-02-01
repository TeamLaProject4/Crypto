package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"cryptomunt/utils"
	"encoding/hex"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Key ecdsa.PrivateKey
}

func CreateGenesisWallet() Wallet {
	mnemonicBytes := utils.ReadFileBytes("./keys/demo-keys/wallet-mnemonic-genesis.txt")
	mnemonic := string(mnemonicBytes)

	bipKey := NewMasterKey(mnemonic)
	key := ConvertBip32ToECDSA(bipKey)
	return CreateWallet(key)
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
	//pubKey := wallet.Key.PublicKey

	//pubKeyBytes := elliptic.Marshal(pubKey, pubKey.X, pubKey.Y)
	//ethCrypto.FromECDSA(wallet.key)

	pubKeyBytes := ethCrypto.FromECDSAPub(&wallet.Key.PublicKey)

	return hex.EncodeToString(pubKeyBytes)
}

func (wallet *Wallet) GetPublicKeyString() string {
	return string(ethCrypto.FromECDSAPub(&wallet.Key.PublicKey))
}

func IsValidSignature(data string, signature string, publicKeyString string) bool {
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
	matches := bytes.Equal(sigPublicKey, []byte(publicKeyString))

	return matches
}
