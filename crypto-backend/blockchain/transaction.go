package blockchain

import (
	"encoding/json"
	"github.com/mitchellh/hashstructure"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	SenderPublicKey   string
	ReceiverPublicKey string
	Amount            int
	TransactionType   TransactionType
	Id                string
	Timestamp         int64
	Signature         string
}

var transaction = new(Transaction)

func NewTransaction(newTransaction Transaction) *Transaction {
	//check and fill variables if they are empty
	if newTransaction.Id == "" {
		newTransaction.Id = uuid.New().String()
	}
	if newTransaction.Timestamp == 0 {
		newTransaction.Timestamp = time.Now().Unix()
	}

	transaction = &newTransaction
	return transaction
}

func transactionEquals(transactionToCompareTo Transaction) bool {
	return transaction.Id == transactionToCompareTo.Id
}

func hashTransaction() uint64 {
	hash, err := hashstructure.Hash(transaction, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func transactionToJson(transaction Transaction) string {
	transactionJson, err := json.Marshal(transaction)
	if err != nil {
		panic("ERROR")
	}
	return string(transactionJson)
}

func transactionPayload() string {
	copyTransaction := Transaction{
		SenderPublicKey:   transaction.SenderPublicKey,
		ReceiverPublicKey: transaction.ReceiverPublicKey,
		Amount:            transaction.Amount,
		TransactionType:   transaction.TransactionType,
		Id:                transaction.Id,
		Timestamp:         transaction.Timestamp,
		Signature:         "",
	}
	return transactionToJson(copyTransaction)
}

func signTransaction(signature string) {
	transaction.Signature = signature
}

type TransactionType int64

const (
	TRANSFER TransactionType = iota
	EXCHANGE
	STAKE
)
