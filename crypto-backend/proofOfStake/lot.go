package proofOfStake

import (
	"cryptomunt/utils"
	"fmt"
	"strconv"
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

func GetLotHash(lot Lot) string {
	hash := lot.PublicKey + lot.PreviousBlockHash + strconv.Itoa(lot.Iteration)
	fmt.Println("hash lot ", hash)
	return utils.GetHexadecimalHash(hash)

	//TODO: python has a loop. Is a loop also required here?
	//for i := lot.Iteration; i > 0; i-- {
	//	h.Write([]byte(hash))
	//	hash += string(h.Sum(nil))
	//}
	//h.Write([]byte("hello world\n"))
	//fmt.Printf("%x", h.Sum(nil))
}
