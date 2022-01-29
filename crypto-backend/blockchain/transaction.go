package blockchain

import (
	"encoding/json"
	"github.com/mitchellh/hashstructure"
)

type Transaction struct {
	SenderPublicKey   string          `json:"sender_pk"`
	ReceiverPublicKey string          `json:"receiver_pk"`
	Amount            int             `json:"amount"`
	Type              TransactionType `json:"tx_type"`
	Id                string          `json:"id"`
	Timestamp         int64           `json:"timestamp"`
	Signature         string          `json:"signature"`
}

func (transaction *Transaction) hash() uint64 {
	hash, err := hashstructure.Hash(transaction, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func (transaction *Transaction) ToJson() string {
	transactionJson, err := json.Marshal(transaction)
	if err != nil {
		panic("ERROR")
	}
	return string(transactionJson)
}

func (transaction *Transaction) Payload() string {
	tempSignature := transaction.Signature
	transaction.Signature = ""

	transactionJson := transaction.ToJson()
	transaction.Signature = tempSignature

	return transactionJson
}

func (transaction *Transaction) Sign(signature string) {
	transaction.Signature = signature
}

type TransactionType int64

const (
	TRANSFER TransactionType = iota //transactions
	EXCHANGE                        //buy coins with fiad money
	STAKE                           //stake in the lottery
)
