package proofOfStake

import (
	"cryptomunt/utils"
)

type Lot struct {
	PublicKey         string
	Iteration         int
	PreviousBlockHash string
}

func CreateLot(lot Lot) Lot {
	return lot
}

func (lot *Lot) Hash() string {
	hash := lot.PublicKey + lot.PreviousBlockHash
	for i := lot.Iteration; i > 0; i-- {
		hash = utils.GetHexadecimalHash(hash)
	}
	return hash
}
