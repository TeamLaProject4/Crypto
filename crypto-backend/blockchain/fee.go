package blockchain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const FEE_PERCENTAGE = 0.004

func CalculateFee(amount int) int {
	fee := float64(amount) * FEE_PERCENTAGE
	if int(fee) <= 0 {
		return 1
	}
	return int(fee)
}

func CalculateInitialAmount(amountIncludingFee int) int {
	initialAmount := float64(amountIncludingFee) / (1.0 + FEE_PERCENTAGE)
	return int(initialAmount)
}

func CreateRewardTransaction(forger string, transactions []Transaction) Transaction {
	reward := CalculateTotalReward(transactions)
	transaction := new(Transaction)
	transaction.ReceiverPublicKey = forger
	transaction.Amount = reward
	transaction.Type = REWARD
	transaction.Id = uuid.New().String()
	transaction.Timestamp = time.Now().Unix()
	return *transaction
}

func CalculateTotalReward(transactions []Transaction) int {
	reward := 0
	for _, transaction := range transactions {
		fmt.Println("||", CalculateFee(CalculateInitialAmount(transaction.Amount)))
		fee := CalculateFee(CalculateInitialAmount(transaction.Amount))
		reward += fee
	}
	return reward
}
