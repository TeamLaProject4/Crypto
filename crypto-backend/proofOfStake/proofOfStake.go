package proofOfStake

import (
	"cryptomunt/utils"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

type ProofOfStake struct {
	stakers map[string]int
}

const MAX_256_INT_VALUE = "10000000000000000000000000000000000000000000000000000000000000000"

var proofOfStake = new(ProofOfStake)

func GetProofOfStake() *ProofOfStake {
	return proofOfStake
}

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

func isOffsetSmaller(lot Lot, seedInt big.Int, leastOffset big.Int) bool {
	offset := getOffset(lot, seedInt)
	compareOffset := offset.Cmp(&leastOffset)

	return compareOffset == -1
}

func getOffset(lot Lot, seedInt big.Int) big.Int {
	lotHash := GetLotHash(lot)
	lotHashInt := utils.GetBigIntHash(lotHash)

	var difference = new(big.Int)
	difference.Sub(&lotHashInt, &seedInt)

	return utils.GetAbsolutBigInt(*difference)
}

func pickWinner(lots []Lot, seed string) Lot {
	winner := lots[0]
	leastOffset := utils.HexToBigInt(MAX_256_INT_VALUE)
	seedInt := utils.GetBigIntHash(seed)

	for _, lot := range lots {
		if isOffsetSmaller(lot, seedInt, leastOffset) {
			leastOffset = getOffset(lot, seedInt)
			winner = lot
		}
	}
	return winner
}

func PickForger(previousBlockHash string) string {
	lots := GenerateLots(previousBlockHash)
	winner := pickWinner(lots, previousBlockHash)
	return winner.PublicKey
}
