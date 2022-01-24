package proofOfStake

import (
	"crypto/sha256"
	"encoding/hex"
)

type Lot struct {
	PublicKey         string
	Iteration         int
	PreviousBlockHash string
}

var lot = new(Lot)

func NewLot(newLot Lot) {
	lot = &newLot
}

func GetLotHash() string {
	hash := lot.PublicKey + lot.PreviousBlockHash
	hasher := sha256.New()
	hasher.Write([]byte(hash))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)

	//TODO: python has a loop. Is a loop also required here?
	//for i := lot.Iteration; i > 0; i-- {
	//	h.Write([]byte(hash))
	//	hash += string(h.Sum(nil))
	//}
	//h.Write([]byte("hello world\n"))
	//fmt.Printf("%x", h.Sum(nil))
}
