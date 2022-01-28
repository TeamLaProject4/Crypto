package blockchain

import (
	"cryptomunt/utils"
	"errors"
)

type MemoryPool struct {
	transactions []Transaction
}

const TRANSACTION_THRESHOLD = 100

func CreateMemoryPool() *MemoryPool {
	memoryPool := new(MemoryPool)
	memoryPool.transactions = make([]Transaction, 150) //150 ideally should be 100
	return memoryPool
}

func (memoryPool *MemoryPool) IsTransactionInPool(transaction Transaction) bool {
	for _, transactionInPool := range memoryPool.transactions {
		if transactionInPool == transaction {
			return true
		}
	}
	return false
}

func (memoryPool *MemoryPool) AddTransaction(transaction Transaction) {
	if !memoryPool.IsTransactionInPool(transaction) {
		memoryPool.transactions = append(memoryPool.transactions, transaction)
	}
}

func (memoryPool *MemoryPool) AddTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		memoryPool.AddTransaction(transaction)
	}
}

func (memoryPool *MemoryPool) GetTransactionsLength() int {
	return len(memoryPool.transactions)
}

func (memoryPool *MemoryPool) GetTransactionIndex(transaction Transaction) (int, error) {
	for index, transactionInPool := range memoryPool.transactions {
		if transactionInPool == transaction {
			return index, nil
		}
	}
	err := errors.New("Transaction not found in memory pool!")
	return -1, err
	//panic("Transaction not found in memory pool!")
}

func (memoryPool *MemoryPool) RemoveTransaction(transaction Transaction) {

	index, err := memoryPool.GetTransactionIndex(transaction)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	// Remove the element at index
	memoryPool.transactions[index] = memoryPool.transactions[len(memoryPool.transactions)-1] // Copy last element to index i.
	memoryPool.transactions[len(memoryPool.transactions)-1] = *new(Transaction)              // Erase last element (write zero value).
	memoryPool.transactions = memoryPool.transactions[:len(memoryPool.transactions)-1]       // Truncate slice.
}

func (memoryPool *MemoryPool) RemoveTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		memoryPool.RemoveTransaction(transaction)
	}
}

func (memoryPool *MemoryPool) IsTransactionThresholdReached() bool {
	return len(memoryPool.transactions) >= TRANSACTION_THRESHOLD
}
