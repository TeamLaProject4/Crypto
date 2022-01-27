package blockchain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mitchellh/hashstructure"
	"time"
)

type Transaction struct {
	SenderPublicKey   string          `json:"sender_pk"`
	ReceiverPublicKey string          `json:"receiver_pk"`
	Amount            int             `json:"amount"`
	TransactionType   TransactionType `json:"tx_type"`
	Id                string          `json:"id"`
	Timestamp         int64           `json:"timestamp"`
	Signature         string          `json:"signature"`
}

func NewTransaction(transaction Transaction) Transaction {
	//check and fill variables if they are empty
	if transaction.Id == "" {
		transaction.Id = uuid.New().String()
	}
	if transaction.Timestamp == 0 {
		transaction.Timestamp = time.Now().Unix()
	}

	return transaction
}

func (transaction *Transaction) Equals(transactionToCompareTo Transaction) bool {
	return transaction.Id == transactionToCompareTo.Id
}

func (transaction *Transaction) hash() uint64 {
	hash, err := hashstructure.Hash(transaction, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func (transaction *Transaction) toJson() string {
	transactionJson, err := json.Marshal(transaction)
	if err != nil {
		panic("ERROR")
	}
	return string(transactionJson)
}

func (transaction *Transaction) Payload() string {
	tempSignature := transaction.Signature
	transaction.Signature = ""

	transactionJson := transaction.toJson()
	transaction.Signature = tempSignature

	return transactionJson
}

func (transaction *Transaction) Sign(signature string) {
	transaction.Signature = signature
}

type TransactionType int64

const (
	TRANSFER TransactionType = iota //transaction
	EXCHANGE                        //buy coins with fiad money
	STAKE                           //stake in the lottery
)
