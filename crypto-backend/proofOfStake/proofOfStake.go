package proofOfStake

import (
	"cryptomunt/utils"
	"fmt"
	"math/big"
)

const MAX_256_INT_VALUE = "10000000000000000000000000000000000000000000000000000000000000000"

type ProofOfStake struct {
	Stakers map[string]int
	genesisPublicKey string
}

func NewProofOfStake() ProofOfStake {
	pos := new(ProofOfStake)
	pos.Stakers = make(map[string]int)
	pos.genesisPublicKey = "key" // TODO: load key from file/config
	return *pos
}

func (proofOfStake *ProofOfStake) PickForger(previousBlockHash string) string {
	lots := proofOfStake.generateLots(previousBlockHash)
	winner := pickWinner(lots, previousBlockHash)
	return winner.PublicKey
}

func (proofOfStake *ProofOfStake) GetStake(publicKey string) int {
	proofOfStake.AddAccountToStakers(publicKey)
	return proofOfStake.Stakers[publicKey]
}

func (proofOfStake *ProofOfStake) AddAccountToStakers(publicKey string) {
	if !proofOfStake.IsAccountInStakers(publicKey) {
		proofOfStake.Stakers[publicKey] = 0
	}
}

func (proofOfStake *ProofOfStake) IsAccountInStakers(publicKey string) bool {
	_, accountInStakers := proofOfStake.Stakers[publicKey]
	return accountInStakers
}

func (proofOfStake *ProofOfStake) UpdateStake(publicKey string, stake int) {
	proofOfStake.AddAccountToStakers(publicKey)
	proofOfStake.Stakers[publicKey] += stake
}

// func (proofOfStake *ProofOfStake) setGenesisNodeStake() {
// 	proofOfStake.Stakers[proofOfStake.genesisPublicKey] = 1
// }

func (proofOfStake *ProofOfStake) generateLots(seed string) []Lot {
	lots := generateLotsFromStakers(proofOfStake, seed)

	if len(lots) == 0 {
		lots = generateLotFromGenesis(proofOfStake, seed)
		fmt.Println(lots[0])
	}

	return lots
}

func generateLotsFromStakers(proofOfStake *ProofOfStake, seed string) []Lot {
	var lots []Lot
	for stakePublicKey, stakeAmount := range proofOfStake.Stakers {
		lots = append(lots, generateStakerLots(stakePublicKey, stakeAmount, seed)...)
	}
	return lots
}

func generateLotFromGenesis(proofOfStake *ProofOfStake, seed string) []Lot {
	lots := generateStakerLots(proofOfStake.genesisPublicKey, 1 , seed)
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

func isOffsetSmaller(lot Lot, seedInt big.Int, leastOffset big.Int) bool {
	offset := getOffset(lot, seedInt)
	compareOffset := offset.Cmp(&leastOffset)

	return compareOffset == -1
}

func getOffset(lot Lot, seedInt big.Int) big.Int {
	lotHash := lot.Hash()
	lotHashInt := utils.GetBigIntHash(lotHash)

	var difference = new(big.Int)
	difference.Sub(&lotHashInt, &seedInt)

	return utils.GetAbsolutBigInt(*difference)
}