package proofOfStake

import (
	"cryptomunt/utils"
	"fmt"
	"math/big"
)

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

func getOffset(lot Lot, seedInt big.Int) big.Int {
	lotHash := lot.Hash()
	lotHashInt := utils.GetBigIntHash(lotHash)

	var difference = new(big.Int)
	difference.Sub(&lotHashInt, &seedInt)

	return utils.GetAbsolutBigInt(*difference)
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

type ProofOfStake struct {
	Stakers map[string]int
}

const MAX_256_INT_VALUE = "10000000000000000000000000000000000000000000000000000000000000000"

//func GetProofOfStake() *ProofOfStake {
//	return proofOfStake
//}

func NewProofOfStake() ProofOfStake {
	return ProofOfStake{Stakers: map[string]int{}}
}

func (proofOfStake *ProofOfStake) PrintStakers() {
	fmt.Println(proofOfStake.Stakers)
}

// SetGenesisNodeStake TODO: unit test
//TODO: remove hardcoded path and add more generic solution
// TODO: after others stake, then genesis can no longer be forger, if everyone unstaked, then genesis can be forger
func (proofOfStake *ProofOfStake) SetGenesisNodeStake() {
	genesisPublicKey := utils.GetFileContents("../keys/genesisPublicKey.pem")
	proofOfStake.Stakers[genesisPublicKey] = 1
}

func (proofOfStake *ProofOfStake) IsAccountInStakers(publicKey string) bool {
	_, accountInStakers := proofOfStake.Stakers[publicKey]
	return accountInStakers
}

func (proofOfStake *ProofOfStake) AddAccountToStakers(publicKey string) {
	if !proofOfStake.IsAccountInStakers(publicKey) {
		proofOfStake.Stakers[publicKey] = 0
	}
}

func (proofOfStake *ProofOfStake) UpdateStake(publicKey string, stake int) {
	proofOfStake.AddAccountToStakers(publicKey)
	proofOfStake.Stakers[publicKey] += stake
}

func (proofOfStake *ProofOfStake) GetStake(publicKey string) int {
	//TODO: addAccount to Stakers not logical here, remove?
	proofOfStake.AddAccountToStakers(publicKey)
	return proofOfStake.Stakers[publicKey]
}

func (proofOfStake *ProofOfStake) GenerateLots(seed string) []Lot {
	var lots []Lot
	for stakePublicKey, stakeAmount := range proofOfStake.Stakers {
		lots = append(lots, generateStakerLots(stakePublicKey, stakeAmount, seed)...)
	}
	return lots
}

func (proofOfStake *ProofOfStake) PickForger(previousBlockHash string) string {
	lots := proofOfStake.GenerateLots(previousBlockHash)
	winner := pickWinner(lots, previousBlockHash)
	return winner.PublicKey
}
