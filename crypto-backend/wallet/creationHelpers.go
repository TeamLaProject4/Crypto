package wallet

import (
	"cryptomunt/blockchain"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"time"
)

func (wallet *Wallet) CreateTransaction(receiverPublicKey string, amount int,
	transactionType blockchain.TransactionType) blockchain.Transaction {

	pubKeyString := string(ethCrypto.FromECDSAPub(&wallet.Key.PublicKey))

	transaction := blockchain.Transaction{
		SenderPublicKey:       wallet.GetPublicKeyHex(),
		SenderPublicKeyString: pubKeyString,
		ReceiverPublicKey:     receiverPublicKey,
		Amount:                amount,
		Type:                  transactionType,
		Id:                    uuid.New().String(),
		Timestamp:             time.Now().Unix(),
	}

	//signature := wallet.Sign(transaction.ToJson())
	signature := wallet.Sign(transaction.Payload())
	transaction.Sign(signature)

	return transaction
}

func (wallet *Wallet) CreateBlock(transactions []blockchain.Transaction, previousHash string,
	blockCount int) blockchain.Block {

	block := blockchain.Block{
		Transactions: transactions,
		PreviousHash: previousHash,
		Height:       blockCount,
		Timestamp:    time.Now().Unix(),
	}

	signature := wallet.Sign(block.Payload())
	block.Sign(signature)

	return block
}
