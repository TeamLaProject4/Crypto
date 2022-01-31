package blockchain

import (
	"cryptomunt/utils"
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

func GetTransactionFromJson(jsonData string) Transaction {
	var transaction Transaction
	err := json.Unmarshal([]byte(jsonData), &transaction)
	if err != nil {
		utils.Logger.Error("unmarshal error ", err)
	}
	return transaction
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
	TRANSFER TransactionType = iota //Transactions
	EXCHANGE                        //buy coins with fiad money
	STAKE                           //stake in the lottery
	REWARD
)
