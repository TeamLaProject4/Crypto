package main

import (
	"fmt"
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

type ComplexStruct struct {
	Name string
}

var transaction = new(Transaction)

func NewTransaction(
	newTransaction Transaction,
) {
	//check and fill variables if they are empty
	if newTransaction.Id == "" {
		newTransaction.Id = uuid.New().String()
	}
	if newTransaction.Timestamp == 0 {
		newTransaction.Timestamp = time.Now().Unix()
	}
	//transaction.signature emptystring?

	transaction = &newTransaction
	//return transaction
}

func TransactionEquals(transactionToCompareTo Transaction) bool {
	return transaction.Id == transaction.Id
}

func Hash() uint64 {
	hash, err := hashstructure.Hash(transaction, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", hash)
	return hash
}

//func (self *Transaction) to_json() map[interface{}]interface{} {
//	return self.__dict__
//}
//
//func (self *Transaction) payload() map[interface{}]interface{} {
//	json_repr := deepcopy(self.to_json())
//	json_repr["signature"] = ""
//	return json_repr
//}

func sign(signature string) {
	transaction.Signature = signature
}

type TransactionType int64

const (
	TRANSFER TransactionType = iota
	EXCHANGE
	STAKE
)
