package proofOfStake

import (
	"cryptomunt/utils"
	"encoding/hex"
	"fmt"
	"math/big"
)

const MAX_256_INT_VALUE = "10000000000000000000000000000000000000000000000000000000000000000"

type ProofOfStake struct {
	Stakers          map[string]int
	GenesisPublicKey string
}

type NegativeBalanceError struct {
	account        string
	currentBalance int
	withdrawAmount int
	Msg            string
}

func (e *NegativeBalanceError) Error() string {
	return e.Msg
}

func CreateProofOfStake() ProofOfStake {
	pos := new(ProofOfStake)
	pos.Stakers = make(map[string]int)

	pubkeyBytes := utils.ReadFileBytes("./keys/demo-keys/wallet-pubkey-genesis.txt")
	hex.EncodeToString(pubkeyBytes)
	pos.GenesisPublicKey = hex.EncodeToString(pubkeyBytes)
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

func (proofOfStake *ProofOfStake) UpdateStake(publicKey string, stake int) error {
	proofOfStake.AddAccountToStakers(publicKey)
	err := validateNegativeStake(proofOfStake, publicKey, stake)
	if err != nil {
		return err
	}
	proofOfStake.Stakers[publicKey] += stake
	return nil
}

func validateNegativeStake(proofOfStake *ProofOfStake, publicKey string, stake int) error {
	if stake < 0 && !proofOfStake.balanceIsSufficient(publicKey, -stake) {
		currentBalance := proofOfStake.GetStake(publicKey)
		return &NegativeBalanceError{
			account:        publicKey,
			currentBalance: currentBalance,
			withdrawAmount: stake,
			Msg:            fmt.Sprintf("Unstake amount (%d) cannot be greater than balance (%d)", stake, currentBalance),
		}
	}
	return nil
}

func (proofOfStake *ProofOfStake) balanceIsSufficient(publicKey string, withdrawAmount int) bool {
	return proofOfStake.GetStake(publicKey) >= withdrawAmount
}

func (proofOfStake *ProofOfStake) generateLots(seed string) []Lot {
	lots := generateLotsFromStakers(proofOfStake, seed)

	if len(lots) == 0 {
		lots = generateLotFromGenesis(proofOfStake, seed)
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
	lots := generateStakerLots(proofOfStake.GenesisPublicKey, 1, seed)
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
