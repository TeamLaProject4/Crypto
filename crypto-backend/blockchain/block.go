package main

import (
	"reflect"
	"time"
)

type Block struct {
	transactions  []Transaction
	previous_hash string
	forger        string
	height        int
	timestamp     int
	signature     string
}

func NewBlock(
	transactions []Transaction,
	previous_hash string,
	forger string,
	height int,
	timestamp int,
	signature string,
) (self *Block) {
	self = new(Block)
	self.transactions = transactions
	self.previous_hash = previous_hash
	self.forger = forger
	self.height = height
	self.timestamp = func() int {
		if timestamp != nil {
			return timestamp
		}
		return int(float64(time.Now().UnixNano()) / 1000000000.0)
	}()
	self.signature = func() string {
		if signature != nil {
			return signature
		}
		return ""
	}()
	return
}

func (self *Block) Eq(block Block) bool {
	return self.signature == block.signature &&
		reflect.DeepEqual(self.previous_hash, block.previous_hash) &&
		(reflect.DeepEqual(self.previous_hash, block.previous_hash) && reflect.DeepEqual(self.forger, block.forger)) &&
		(reflect.DeepEqual(self.previous_hash, block.previous_hash) && reflect.DeepEqual(self.forger, block.forger) && reflect.DeepEqual(self.height, block.height)) &&
		(reflect.DeepEqual(self.previous_hash, block.previous_hash) && reflect.DeepEqual(self.forger, block.forger) && reflect.DeepEqual(self.height, block.height) && self.timestamp == block.timestamp)
}

func (self *Block) to_json() map[interface{}]interface{} {
	json_dict := map[string][]interface{}{}
	json_dict["previous_hash"] = self.previous_hash
	json_dict["forger"] = self.forger
	json_dict["height"] = self.height
	json_dict["timestamp"] = self.timestamp
	json_dict["signature"] = self.signature
	json_dict["transactions"] = func() (elts []interface{}) {
		for _, tx := range self.transactions {
			elts = append(elts, tx.to_json())
		}
		return
	}()
	return json_dict
}

func (self *Block) payload() map[interface{}]interface{} {
	json_repr := deepcopy(self.to_json())
	json_repr["signature"] = ""
	return json_repr
}

func (self *Block) sign(signature string) {
	self.signature = signature
}

func genesis() Block {
	genesis := NewBlock([]interface{}{}, "genesis_hash", "genesis_forger", 0, 0)
	return genesis
}
