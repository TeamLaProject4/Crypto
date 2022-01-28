package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"cryptomunt/utils"
	"encoding/hex"
	"fmt"
	"os"
)

var KEY_LENGTH_BITS = 2048
var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.hex"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.hex"

//Keep wallet keys private
type Wallet struct {
	key rsa.PrivateKey //contains private and public keys
}

//// NewWallet TODO: mnumonic?
func CreateWallet() Wallet {
	return Wallet{key: GetKeyPair()}
}

//get the keypair from a file or generate one
//func GetKeyPair() (string, string) {
func GetKeyPair() rsa.PrivateKey {
	var privateKey rsa.PrivateKey

	//get from file
	privateKey = utils.ReadRsaKeyFile("../keys/wallet.rsa")

	//generate key
	if privateKey.Size() < 1 {
		key, _ := rsa.GenerateKey(rand.Reader, KEY_LENGTH_BITS)
		privateKey = *key
		//TODO: error handling
		utils.WriteRsaKeyToFile(privateKey)
	}

	return privateKey
}

func (wallet *Wallet) sign(data string) string {
	message := []byte(data)
	//Sign hash of message because only small messages can be signed
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(rand.Reader, &wallet.key, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return ""
	}

	return hex.EncodeToString(signature)
}

func (wallet *Wallet) GetPublicKeyHex() string {
	return utils.GetRsaPublicKeyHexValue(&wallet.key.PublicKey)
}

func (wallet *Wallet) createTransaction(
	receiverPublicKey string,
	amount int,
	transactionType TransactionType,
) Transaction {
	transaction := CreateTransaction(
		Transaction{
			SenderPublicKey:   wallet.GetPublicKeyHex(),
			ReceiverPublicKey: receiverPublicKey,
			Amount:            amount,
			Type:              transactionType,
			Id:                "",
			Timestamp:         0,
			Signature:         "",
		})

	signature := wallet.sign(transaction.toJson())
	transaction.Sign(signature)
	return transaction
}

func (wallet *Wallet) CreateBlock(
	transactions []Transaction,
	previousHash string,
	blockCount int,
) Block {
	block := CreateBlock(
		Block{
			Transactions: transactions,
			PreviousHash: previousHash,
			Forger:       "",
			Height:       blockCount,
			Timestamp:    0,
			Signature:    wallet.GetPublicKeyHex(),
		})
	signature := wallet.sign(block.Payload())
	block.Sign(signature)
	return block
}

func IsValidSignature(blockPayload string, signatureHex string, publicKey string) bool {
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false
	}

	message := []byte(blockPayload)
	hashed := sha256.Sum256(message)

	pubKey, _ := utils.GetPublicKeyFromHex(publicKey)
	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], signature)

	if err != nil {
		return false
	}
	return true
}
