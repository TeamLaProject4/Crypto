package blockchain

import (
	"crypto/rand"
	"crypto/rsa"
	"cryptomunt/utils"
)

var KEY_LENGTH_BITS = 2048
var PRIVATE_KEY_PATH = "../keys/walletPrivateKey.hex"
var PUBLIC_KEY_PATH = "../keys/walletPublicKey.hex"

//Keep wallet keys private
type Wallet struct {
	key rsa.PrivateKey //contains private and public keys
	//privateKey string
	//publicKey  string
	//privateKey rsa.PrivateKey
	//publicKey rsa.PublicKey
}

//func (wallet *Wallet) GetPublicKey() string {
//	return wallet.publicKey
//}

//// NewWallet TODO: mnumonic?
func NewWallet(privateKey string) Wallet {
	return Wallet{key: GetKeyPair()}
	//publicKey, privateKeyFinal := getKeyPair()
	//return Wallet{
	//	privateKey: privateKeyFinal,
	//	publicKey:  publicKey,
	//}
}

//get the keypair from a file or generate one
//func GetKeyPair() (string, string) {
func GetKeyPair() rsa.PrivateKey {
	//privateKey := utils.GetFileContents(PRIVATE_KEY_PATH)
	//publicKey := utils.GetFileContents(PUBLIC_KEY_PATH)
	//var privateKey rsa.PrivateKey

	//get from file

	//generate
	//if privateKey == "" {
	privateKey, err := rsa.GenerateKey(rand.Reader, KEY_LENGTH_BITS)
	if err != nil {
		panic("Failed to make a wallet")
	}
	utils.WriteRsaKeyToFile(*privateKey)
	//publicKey = utils.BigIntToHex(*key.N)
	//privateKey = utils.BigIntToHex(*key.D)

	//save to file
	//utils.WriteFile(PRIVATE_KEY_PATH, privateKey)
	//utils.WriteFile(publicKey, publicKey)
	//}

	return *privateKey
	//return publicKey, privateKey
}

//func (wallet *Wallet) sign(data string) string {
//	dataHash := utils.GetHexadecimalHash(data)
//	//pkcs1
//	//sign using signatureSchemeObject
//
//	message := []byte("message to be signed")
//
//	// Only small messages can be signed directly; thus the hash of a
//	// message, rather than the message itself, is signed. This requires
//	// that the hash function be collision resistant. SHA-256 is the
//	// least-strong hash function that should be used for this at the time
//	// of writing (2016).
//	hashed := sha256.Sum256(message)
//
//	test := utils.HexToBigInt(wallet.privateKey)
//	rsaPrivateKey := rsa.PrivateKey{
//		PublicKey: rsa.PublicKey{},
//		D:         &test,
//	}
//
//	signature, err := rsa.SignPKCS1v15(rand.Reader, &rsaPrivateKey, crypto.SHA256, hashed[:])
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
//		return
//	}
//
//	fmt.Printf("Signature: %x\n", signature)
//
//	//signatureSchemeObject := PKCS1_v1_5.new(wallet.KeyPair)
//	//signature := signatureSchemeObject.sign(dataHash)
//	//return signature.hex()
//}

//func (wallet *Wallet) publicKeyString() string {
//	return string(wallet.KeyPair.public_key().export_key("PEM"))
//}

//func (wallet *Wallet) createTransaction(
//	receiverPublicKey string,
//	amount int,
//	transactionType TransactionType,
//) Transaction {
//
//	transaction := NewTransaction(
//		Transaction{
//			SenderPublicKey:   wallet.publicKeyString(),
//			ReceiverPublicKey: receiverPublicKey,
//			Amount:            amount,
//			TransactionType:   transactionType,
//			Id:                "",
//			Timestamp:         0,
//			Signature:         "",
//		})
//
//	signature := wallet.sign(transaction.ToJson())
//	transaction.Sign(signature)
//	return transaction
//}
//
//func (wallet *Wallet) createBlock(
//	transactions []Transaction,
//	previousHash string,
//	blockCount int,
//) Block {
//	block := CreateBlock(
//		Block{
//			Transactions: transactions,
//			PreviousHash: previousHash,
//			Forger:       "",
//			Height:       blockCount,
//			Timestamp:    0,
//			Signature:    wallet.publicKeyString(),
//		})
//		transactions, previousHash, wallet.publicKeyString(), blockCount)
//	signature := wallet.sign(block.payload())
//	block.sign(signature)
//	return block
//}
//
//func (data *Wallet) isValidSignature(signature string, publicKey string) bool {
//	signature := bytes.fromhex(signature)
//	data_hash := BlockchainUtils.hash(data)
//	public_key := RSA.importKey(publicKey)
//	signature_scheme_object := PKCS1_v1_5.new(public_key)
//	return signature_scheme_object.verify(data_hash, signature)
//}
