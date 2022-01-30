package blockchain

const FEE_PERCENTAGE = 0.004

func CalculateFee(amount int) int {
	fee := float64(amount) * FEE_PERCENTAGE
	if fee < 0 {
		return 1
	}
	return int(fee)
}

func CalculateInitialAmount(amountIncludingFee int) int {
	initialAmount := float64(amountIncludingFee) / 1.0 + FEE_PERCENTAGE
	return int(initialAmount)
}