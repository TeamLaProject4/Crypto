package blockchain

import "sync"

// GetBlocksFromRange, if end > len(blocks) then return all blocks
func (blockchain *Blockchain) GetBlocksFromRange(start int, end int) []Block {
	if len(blockchain.Blocks) <= 0 {
		return nil
	}
	if end > len(blockchain.Blocks) {
		return blockchain.Blocks[start:len(blockchain.Blocks)]
	}

	blocks := blockchain.Blocks[start:end]
	return blocks
}

//func CreateBlock(Transactions []Transaction, forgerWallet interface{}) Block {
//	covered_transactions := self.get_covered_transactions(Transactions)
//	self.execute_transactions(covered_transactions)
//	block := forgerWallet.create_block(
//		covered_transactions,
//		self.latest_previous_hash(),
//		self.latest_block_height()+1,
//	)
//	self.add_block(block)
//	return block
//}

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
			go getAccountTransactionsFromBlock(block, blocksTransactions, publicKey)
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

func getAccountTransactionsFromBlock(block Block, transactions chan []Transaction, publicKey string) {
	transactionsFromBlock := *new([]Transaction)
	for _, transaction := range block.Transactions {
		if transaction.SenderPublicKey == publicKey || transaction.ReceiverPublicKey == publicKey {
			transactionsFromBlock = append(transactionsFromBlock, transaction)
		}
	}

	transactions <- transactionsFromBlock
}
