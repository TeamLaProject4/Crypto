package blockchain

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"encoding/json"
	"sync"
)

type Blockchain struct {
	Blocks       []Block                    `json:"blocks"`
	AccountModel *AccountModel              `json:"-"`
	ProofOfStake *proofOfStake.ProofOfStake `json:"-"`
}

func CreateBlockchain() Blockchain {
	genesisBlock := CreateGenesisBlock()
	var blocks []Block
	blocks = append(blocks, genesisBlock)
	pos := proofOfStake.NewProofOfStake()

	accountModel := CreateAccountModel()

	return Blockchain{
		Blocks:       blocks,
		AccountModel: &accountModel,
		ProofOfStake: &pos,
	}
}

func (blockchain *Blockchain) ToJson() string {
	blockchainJson, err := json.Marshal(blockchain)
	if err != nil {
		panic("ERROR")
	}
	return string(blockchainJson)
}

func (blockchain *Blockchain) LatestBlockHeight() int {
	return len(blockchain.Blocks) - 1
}

func (blockchain *Blockchain) AddBlock(block Block) {
	if blockchain.IsValidBlockHeight(block) && blockchain.IsValidPreviousBlockHash(block) {
		blockchain.ExecuteTransactions(block.Transactions)
		blockchain.Blocks = append(blockchain.Blocks, block)
	}
}

func (blockchain *Blockchain) ExecuteTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		blockchain.executeTransaction(transaction)
	}
}

func (blockchain *Blockchain) executeTransaction(transaction Transaction) {
	sender := transaction.SenderPublicKey
	receiver := transaction.ReceiverPublicKey
	amount := transaction.Amount

	if transaction.Type == STAKE {
		if sender == receiver {
			blockchain.ProofOfStake.UpdateStake(sender, amount)
			blockchain.AccountModel.UpdateBalance(sender, -amount)
		}
	} else {
		blockchain.AccountModel.UpdateBalance(sender, -amount)
		blockchain.AccountModel.UpdateBalance(receiver, amount)
	}
}

func (blockchain *Blockchain) IsValidBlockHeight(block Block) bool {
	blockLength := len(blockchain.Blocks)
	return blockchain.Blocks[blockLength-1].Height == block.Height-1
}

func (blockchain *Blockchain) IsValidPreviousBlockHash(block Block) bool {
	return blockchain.LatestPreviousHash() == block.PreviousHash
}

func (blockchain *Blockchain) IsValidForger(block Block) bool {
	return block.Forger == blockchain.getNextForger()
}

func (blockchain *Blockchain) IsBlockTransactionsValid(block Block) bool {
	transactions := block.Transactions
	coveredTransactions := blockchain.GetCoveredTransactions(transactions)
	return len(transactions) == len(coveredTransactions)
}

func (blockchain *Blockchain) LatestPreviousHash() string {
	blockLenght := len(blockchain.Blocks)
	latestBlock := blockchain.Blocks[blockLenght-1]
	payload := latestBlock.Payload()

	return utils.GetHexadecimalHash(payload)
}

func (blockchain *Blockchain) GetCoveredTransactions(transactions []Transaction) []Transaction {
	var coveredTransactions = make([]Transaction, len(transactions))

	for _, transaction := range transactions {
		if blockchain.IsTransactionCovered(transaction) {
			coveredTransactions = append(coveredTransactions, transaction)
		}
	}
	return coveredTransactions
}

func (blockchain *Blockchain) IsTransactionCovered(transaction Transaction) bool {
	if transaction.Type == EXCHANGE {
		return true
	}
	senderBalance := blockchain.AccountModel.GetBalance(transaction.SenderPublicKey)
	return senderBalance >= transaction.Amount
}

func (blockchain *Blockchain) GetAccountBalance(publicKey string) int {
	return blockchain.AccountModel.GetBalance(publicKey)
}

func (blockchain *Blockchain) getNextForger() string {
	prevBlockHash := blockchain.LatestPreviousHash()
	return blockchain.ProofOfStake.PickForger(prevBlockHash)
}

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
		//fmt.Printf("Goroutine: %d  trans: %v  \n", index, transaction)
		if transaction.SenderPublicKey == publicKey || transaction.ReceiverPublicKey == publicKey {
			transactionsFromBlock = append(transactionsFromBlock, transaction)
		}
	}

	transactions <- transactionsFromBlock
}
