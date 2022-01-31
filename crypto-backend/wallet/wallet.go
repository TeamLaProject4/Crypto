package wallet

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"encoding/hex"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/jbenet/go-base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"time"
)

func GenerateMnemonic() string {
	//Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	return mnemonic
}

func (wallet *Wallet) SetWalletKeyAndWritePrivateKeyFile(mnemonic string) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "secret")

	master, err := bip32.NewMasterKey(seed)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	// m/44'
	key, err := master.NewChildKey(2147483648 + 44)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	decoded := base58.Decode(key.B58Serialize())
	privateKey := decoded[46:78]

	//Hex private key to ECDSA private key
	privateKeyECDSA, err := ethCrypto.ToECDSA(privateKey)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	//set wallet using mnemonic
	wal := CreateWallet(*privateKeyECDSA)
	wallet = &wal

	utils.WriteEDCSAToFile(privateKeyECDSA)
}

var KEY_LENGTH_BITS = 2048
var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.hex"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.hex"

type Wallet struct {
	key ecdsa.PrivateKey
}

func CreateWallet(key ecdsa.PrivateKey) Wallet {
	return Wallet{key: key}
}

func (wallet *Wallet) sign(data string) string {
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
	utils.Logger.Info("pubkey", wallet.key.PublicKey)
	pubkey := wallet.key.PublicKey
	pubKeyBytes := elliptic.Marshal(pubkey, pubkey.X, pubkey.Y)
	return hex.EncodeToString(pubKeyBytes)
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
			Height:       blockCount,
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
