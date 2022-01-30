package wallet

import (
	"bufio"
	"crypto"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"os"
	"time"
)

func GenerateMnemonic() {
	//Generate a mnemonic for memorization or user-friendly seeds

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "secret")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	writePrivateKeyFile(masterKey)
	writePublicKeyFile(publicKey)
}

func writePrivateKeyFile(private *bip32.Key) {
	file, err := os.OpenFile("../keys/private.key", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%v\n", private)
	w.Flush()
}

func writePublicKeyFile(public *bip32.Key) {
	file, err := os.OpenFile("../keys/public.key", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%v\n", public)
	w.Flush()
}

var KEY_LENGTH_BITS = 2048
var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.hex"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.hex"

type Wallet struct {
	key rsa.PrivateKey //contains private and public keys, keep private!
}

// CreateWallet TODO: mnumonic?
func CreateWallet() Wallet {
	return Wallet{key: GetKeyPair()}
}

//get the keypair from a file or generate one
//func GetKeyPair() (string, string) {
func GetKeyPair() rsa.PrivateKey {
	var privateKey rsa.PrivateKey

	//get from file
	privateKey = utils.ReadRsaKeyFile("./keys/wallet.rsa")

	//generate key
	if privateKey.Size() < 1 {
		key, err := rsa.GenerateKey(cryptoRand.Reader, KEY_LENGTH_BITS)
		if err != nil {
			utils.Logger.Error("generate rsa error", err)
		}
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

	signature, err := rsa.SignPKCS1v15(cryptoRand.Reader, &wallet.key, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return ""
	}

	return hex.EncodeToString(signature)
}

func (wallet *Wallet) GetPublicKeyHex() string {
	return utils.GetRsaPublicKeyHexValue(&wallet.key.PublicKey)
}

func (wallet *Wallet) CreateTransaction(
	receiverPublicKey string,
	amount int,
	transactionType blockchain.TransactionType,
) blockchain.Transaction {
	transaction := blockchain.Transaction{
		SenderPublicKey:   wallet.GetPublicKeyHex(),
		ReceiverPublicKey: receiverPublicKey,
		Amount:            amount,
		Type:              transactionType,
		Id:                uuid.New().String(),
		Timestamp:         time.Now().Unix(),
	}

	signature := wallet.sign(transaction.ToJson())
	transaction.Sign(signature)

	return transaction
}

func (wallet *Wallet) CreateBlock(
	transactions []blockchain.Transaction,
	previousHash string,
	blockCount int,
) blockchain.Block {
	block := blockchain.CreateBlock(
		blockchain.Block{
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
