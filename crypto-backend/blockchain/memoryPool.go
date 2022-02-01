package blockchain

import (
	"cryptomunt/utils"
	"errors"
)

type MemoryPool struct {
	Transactions []Transaction
}

//const TRANSACTION_THRESHOLD = 99
const TRANSACTION_THRESHOLD = 3

func CreateMemoryPool() MemoryPool {
	memoryPool := new(MemoryPool)
	memoryPool.Transactions = make([]Transaction, 0)
	return *memoryPool
}

func (memoryPool *MemoryPool) IsTransactionInPool(transaction Transaction) bool {
	for _, transactionInPool := range memoryPool.Transactions {
		if transactionInPool == transaction {
			return true
		}
	}
	return false
}

func (memoryPool *MemoryPool) AddTransaction(transaction Transaction) {
	if !memoryPool.IsTransactionInPool(transaction) {
		memoryPool.Transactions = append(memoryPool.Transactions, transaction)
	}
}

func (memoryPool *MemoryPool) AddTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		memoryPool.AddTransaction(transaction)
	}
}

func (memoryPool *MemoryPool) GetTransactionsLength() int {
	return len(memoryPool.Transactions)
}

func (memoryPool *MemoryPool) GetTransactionIndex(transaction Transaction) (int, error) {
	utils.Logger.Info("transaction", transaction.Id)
	for index, transactionInPool := range memoryPool.Transactions {
		utils.Logger.Info("transaction in pool", transactionInPool.Id)
		if transactionInPool.Id == transaction.Id || transactionInPool.Timestamp == transaction.Timestamp {
			return index, nil
		}
	}
	err := errors.New("transaction not found in memory pool")
	return -1, err
}

func (memoryPool *MemoryPool) RemoveTransaction(transaction Transaction) {
	index, err := memoryPool.GetTransactionIndex(transaction)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	memoryPool.Transactions = append(memoryPool.Transactions[:index], memoryPool.Transactions[index+1:]...)

	// Remove the element at index
	//memoryPool.Transactions[index] = memoryPool.Transactions[len(memoryPool.Transactions)-1] // Copy last element to index i.
	//memoryPool.Transactions[len(memoryPool.Transactions)-1] = *new(Transaction)              // Erase last element (write zero value).
	//memoryPool.Transactions = memoryPool.Transactions[:len(memoryPool.Transactions)-1]       // Truncate slice.
}

func (memoryPool *MemoryPool) RemoveTransactions(transactions []Transaction) {
	//if memoryPool.areTransactionsInMemoryPool(transactions) {
	for _, transaction := range transactions {
		memoryPool.RemoveTransaction(transaction)
	}
	//}
}

func (memoryPool *MemoryPool) areTransactionsInMemoryPool(transactions []Transaction) bool {
	for _, transactionFromForgedBlock := range transactions {
		if !memoryPool.containsTransaction(transactionFromForgedBlock) {
			return false
		}
	}

	return true
}
func (memoryPool *MemoryPool) containsTransaction(transaction Transaction) bool {
	for _, transactionFromMemPool := range memoryPool.Transactions {
		if transaction == transactionFromMemPool {
			return true
		}
	}
	return false
}

func (memoryPool *MemoryPool) IsTransactionThresholdReached() bool {
	return len(memoryPool.Transactions) >= TRANSACTION_THRESHOLD
}
