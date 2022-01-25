package wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"cryptomunt/utils"
	"io/ioutil"
	"os"
)

var KEY_LENGTH_BITS = 2048

//type KeyPair struct {
//	privateKey string
//	publicKey  string
//}
type Wallet struct {
	privateKey string
	publicKey  string
	//keypair KeyPair
}

var wallet Wallet

//var wallet = make(Wallet{keypair: KeyPair{
//	privateKey: "",
//	publicKey:  "",
//}})

// NewWallet TODO: mnumonic?
func NewWallet(privateKey string) {
	publicKey := ""

	if privateKey == "" {
		key, err := rsa.GenerateKey(rand.Reader, KEY_LENGTH_BITS)
		if err != nil {
			panic("Failed to make a wallet")
		}
		publicKey = utils.BigIntToHex(*key.N)
		privateKey = utils.BigIntToHex(*key.D)

	} else {
		//publicKey =crypto.PublicKey(privateKey)
		//test := crypto.PublicKey(privateKey)
	}

	wallet = Wallet{
		privateKey: publicKey,
		publicKey:  privateKey,
	}
}

func (self *Wallet) import_key(file_path interface{}) {
	f := func() *os.File {
		f, err := os.OpenFile(file_path, os.O_RDONLY, 0o777)
		if err != nil {
			panic(err)
		}
		return f
	}()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	self.KeyPair = RSA.import_key(func() string {
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		return string(content)
	}())
}

func (self *Wallet) sign(data interface{}) interface{} {
	data_hash := BlockchainUtils.hash(data)
	signature_scheme_object := PKCS1_v1_5.new(self.KeyPair)
	signature := signature_scheme_object.sign(data_hash)
	return signature.hex()
}

func (self *Wallet) public_key_string() string {
	return string(self.KeyPair.public_key().export_key("PEM"))
}

func (self *Wallet) create_transaction(
	receiver_public_key string,
	amount int,
	tx_type TxType,
) Transaction {
	transaction := Transaction(self.public_key_string(), receiver_public_key, amount, tx_type)
	signature := self.sign(transaction.to_json())
	transaction.sign(signature)
	return transaction
}

func (self *Wallet) create_block(
	transactions []Transaction,
	previous_hash string,
	block_count int,
) Block {
	block := Block(transactions, previous_hash, self.public_key_string(), block_count)
	signature := self.sign(block.payload())
	block.sign(signature)
	return block
}

func (data *Wallet) is_valid_signature(signature string, public_key_string string) bool {
	signature := bytes.fromhex(signature)
	data_hash := BlockchainUtils.hash(data)
	public_key := RSA.importKey(public_key_string)
	signature_scheme_object := PKCS1_v1_5.new(public_key)
	return signature_scheme_object.verify(data_hash, signature)
}
