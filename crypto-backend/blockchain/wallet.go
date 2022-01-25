package blockchain

var KEY_LENGTH_BITS = 2048

type Wallet struct {
	privateKey string
	publicKey  string
}

//// NewWallet TODO: mnumonic?
//func NewWallet(privateKey string) Wallet {
//	publicKey, privateKeyFinal := getKeyPair(privateKey)
//	return Wallet{
//		privateKey: privateKeyFinal,
//		publicKey:  publicKey,
//	}
//}
//
//func getKeyPair(privateKey string) (string, string) {
//	publicKey := ""
//
//	if privateKey == "" {
//		key, err := rsa.GenerateKey(rand.Reader, KEY_LENGTH_BITS)
//		if err != nil {
//			panic("Failed to make a wallet")
//		}
//		publicKey = utils.BigIntToHex(*key.N)
//		privateKey = utils.BigIntToHex(*key.D)
//
//	} else {
//		publicKey = crypto.PublicKey(privateKey)
//		//test := crypto.PublicKey(privateKey)
//	}
//
//	return publicKey, privateKey
//}
//
//func (wallet *Wallet) importKey(filePath string) {
//	//key := utils.GetFileContents(filePath)
//	//setkey
//}
//
//func (wallet *Wallet) sign(data string) string {
//	dataHash := utils.GetHexadecimalHash(data)
//	signatureSchemeObject := PKCS1_v1_5.new(wallet.KeyPair)
//	signature := signatureSchemeObject.sign(dataHash)
//	return signature.hex()
//}
//
//func (wallet *Wallet) publicKeyString() string {
//	return string(wallet.KeyPair.public_key().export_key("PEM"))
//}
//
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
//func (data *Wallet) is_valid_signature(signature string, public_key_string string) bool {
//	signature := bytes.fromhex(signature)
//	data_hash := BlockchainUtils.hash(data)
//	public_key := RSA.importKey(public_key_string)
//	signature_scheme_object := PKCS1_v1_5.new(public_key)
//	return signature_scheme_object.verify(data_hash, signature)
//}
