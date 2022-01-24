package proofOfStake

import (
	"cryptomunt/utils"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

type ProofOfStake struct {
	stakers map[string]int
}

var proofOfStake = new(ProofOfStake)

func NewProofOfStake() {
	newProofOfStake := ProofOfStake{stakers: map[string]int{}}
	proofOfStake = &newProofOfStake
}

func PrintStakers() {
	fmt.Println(proofOfStake.stakers)
}

// SetGenesisNodeStake TODO: unit test
//TODO: remove hardcoded path and add more generic solution
// TODO: after others stake, then genesis can no longer be forger, if everyone unstaked, then genesis can be forger
func SetGenesisNodeStake() {
	file, err := os.OpenFile("../keys/genesisPublicKey.pem", os.O_RDONLY, 0o777)
	if err != nil {
		panic(err)
	}
	fileBytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		panic(err2)
	}

	genisisPublicKey := string(fileBytes)
	proofOfStake.stakers[genisisPublicKey] = 1
}

func IsAccountInStakers(publicKey string) bool {
	_, accountInStakers := proofOfStake.stakers[publicKey]
	return accountInStakers
}

func AddAccountToStakers(publicKey string) {
	if !IsAccountInStakers(publicKey) {
		proofOfStake.stakers[publicKey] = 0
	}
}

func UpdateStake(publicKey string, stake int) {
	AddAccountToStakers(publicKey)
	proofOfStake.stakers[publicKey] += stake
}

func GetStake(publicKey string) int {
	//TODO: addAccount to stakers not logical here, remove?
	AddAccountToStakers(publicKey)
	return proofOfStake.stakers[publicKey]
}

func GenerateLots(seed string) []Lot {
	var lots []Lot
	for stakePublicKey, stakeAmount := range proofOfStake.stakers {
		lots = append(lots, generateStakerLots(stakePublicKey, stakeAmount, seed)...)
	}
	return lots
}

func generateStakerLots(stakePublicKey string, stakeAmount int, seed string) []Lot {
	var lots []Lot
	for stake := 0; stake < stakeAmount; stake++ {
		lot := Lot{
			PublicKey:         stakePublicKey,
			Iteration:         stake + 1, //not zero indexed
			PreviousBlockHash: seed,
		}
		lots = append(lots, lot)
	}
	return lots
}

func isOffsetSmaller(lot Lot, seedInt int64, leastOffset int64) bool {
	offset := getOffset(lot, seedInt)
	return offset < leastOffset
}

func getOffset(lot Lot, seedInt int64) int64 {
	lotHash := GetLotHash(lot)
	lotHashInt, err := strconv.ParseInt(lotHash, 16, 64)
	if err != nil {
		return 0
	}
	return int64(math.Abs(float64(lotHashInt - seedInt)))
}

//TODO: is float required for leastOffset?
func pickWinner(lots []Lot, seed string) Lot {
	winner := lots[0]
	var leastOffset int64 = 9223372036854775807 //max 64 integer
	seedHash := utils.GetHash(seed)
	seedInt, err := strconv.ParseInt(seedHash, 16, 64)
	if err != nil {
		panic("Hash conversion error")
	}

	for _, lot := range lots {
		if isOffsetSmaller(lot, seedInt, leastOffset) {
			leastOffset = getOffset(lot, seedInt)
			winner = lot
		}
	}
	return winner
}

func pickForger(previousBlockHash string) string {
	lots := GenerateLots(previousBlockHash)
	winner := pickWinner(lots, previousBlockHash)
	return winner.PublicKey
}
