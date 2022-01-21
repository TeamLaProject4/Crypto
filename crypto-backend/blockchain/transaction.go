package main

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	sender_public_key   string
	receiver_public_key string
	amount              int
	tx_type             TxType
	id                  string
	timestamp           int64
	signature           string
}

var transaction = new(Transaction)

func NewTransaction(
	newTransaction Transaction,
) {
	//check and fill variables if they are empty
	if newTransaction.id == "" {
		newTransaction.id = uuid.New().String()
	}
	if newTransaction.timestamp == 0 {
		newTransaction.timestamp = time.Now().Unix()
	}
	//transaction.signature emptystring?

	transaction = &newTransaction
	//return transaction
}

func (self *Transaction) TransactionEquals(transaction Transaction) bool {
	return reflect.DeepEqual(self.id, transaction.id)
}

/*
func (self *Transaction) Hash() {
	hash(fmt.Sprintf("%#v", self))
}


func (self *Transaction) to_json() map[interface{}]interface{} {
	return self.__dict__
}

func (self *Transaction) payload() map[interface{}]interface{} {
	json_repr := deepcopy(self.to_json())
	json_repr["signature"] = ""
	return json_repr
}
*/
func (self *Transaction) sign(signature string) {
	self.signature = signature
}

type TxType struct{}
