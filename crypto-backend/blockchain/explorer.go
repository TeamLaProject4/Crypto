package blockchain

import (
	"sync"
)

func (blockchain *Blockchain) GetBlocksFromRange(start int, end int) []Block {
	return blockchain.Blocks[start:end]
}

func (blockchain *Blockchain) IsTransactionInBlockchain(transaction Transaction) bool {
	for _, block := range blockchain.Blocks {
		for _, blockTransaction := range block.Transactions {
			if blockTransaction == transaction {
				return true
			}
		}
	}
	return false
}

func (blockchain *Blockchain) GetAllAccountTransactions(publicKey string) []Transaction {
	var wg sync.WaitGroup
	blocksTransactions := make(chan []Transaction)
	for index, block := range blockchain.Blocks {
		wg.Add(1)
		go func(block Block, index int) {
			defer wg.Done()
			go getAccountTransactionsFromBlock(block, blocksTransactions, publicKey, index)
		}(block, index)
	}
	wg.Wait()

	transactions := *new([]Transaction)
	blockBalancesIndex := 0
	for transaction := range blocksTransactions {
		transactions = append(transactions, transaction...)

		if len(blockchain.Blocks)-1 == blockBalancesIndex {
			close(blocksTransactions)
		}
		blockBalancesIndex++
	}

	return transactions
}

func getAccountTransactionsFromBlock(block Block, transactions chan []Transaction, publicKey string, index int) {
	transactionsFromBlock := *new([]Transaction)
	for _, transaction := range block.Transactions {
		if transaction.SenderPublicKey == publicKey || transaction.ReceiverPublicKey == publicKey {
			transactionsFromBlock = append(transactionsFromBlock, transaction)
		}
	}

	transactions <- transactionsFromBlock
}
