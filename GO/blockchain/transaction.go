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
	timestamp           int
	signature           string
}

func NewTransaction(
	sender_public_key string,
	receiver_public_key string,
	amount int,
	tx_type TxType,
	id string,
	timestamp int,
	signature string,
) (self *Transaction) {
	self = new(Transaction)
	self.sender_public_key = sender_public_key
	self.receiver_public_key = receiver_public_key
	self.amount = amount
	self.tx_type = tx_type

	if id != "" {
		self.id = id
	} else {
		id = uuid.New().String()
	}
	self.timestamp = func() int {
		if timestamp != 0 {
			return timestamp
		}
		return int(float64(time.Now().UnixNano()) / 1000000000.0)
	}()
	self.signature = func() string {
		if signature != "" {
			return signature
		}
		return ""
	}()
	return
}

func (self *Transaction) Eq(transaction Transaction) bool {
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
