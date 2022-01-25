package blockchain

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"encoding/json"
)

type Blockchain struct {
	Blocks       []Block
	AccountModel AccountModel
	ProofOfStake proofOfStake.ProofOfStake
}

var blockchain = new(Blockchain)

func GetBlockChain() *Blockchain {
	return blockchain
}
func NewBlockchain() {
	genesisBlock := createGenesisBlock()
	var blocks []Block
	blocks = append(blocks, genesisBlock)
	proofOfStake.NewProofOfStake()

	pos := proofOfStake.GetProofOfStake()

	blockchain = &Blockchain{
		Blocks:       blocks,
		AccountModel: *accountModel,
		ProofOfStake: *pos,
	}
}

func BlockchainToJson() string {
	blockchainJson, err := json.Marshal(blockchain)
	if err != nil {
		panic("ERROR")
	}
	return string(blockchainJson)
}

func getLatestBlockHeight() int {
	return len(blockchain.Blocks) - 1
}

func addBlock(block Block) {
	if isValidBlockCount(block) && isValidPreviousBlockHash(block) {
		executeTransactions(block.Transactions)
		blockchain.Blocks = append(blockchain.Blocks, block)
	}
}

func executeTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		executeTransaction(transaction)
	}
}

func executeTransaction(transaction Transaction) {
	sender := transaction.SenderPublicKey
	receiver := transaction.ReceiverPublicKey
	amount := transaction.Amount

	if transaction.TransactionType == STAKE {
		if sender == receiver {
			proofOfStake.UpdateStake(sender, amount)
			updateAccountModelBalance(sender, -amount)
		}
	} else {
		updateAccountModelBalance(sender, -amount)
		updateAccountModelBalance(receiver, amount)
	}
}

func isValidBlockCount(block Block) bool {
	blockLenght := len(blockchain.Blocks)
	return blockchain.Blocks[blockLenght-1].Height == block.Height-1
}

func isValidPreviousBlockHash(block Block) bool {
	return getLatestPreviousHash() == block.PreviousHash
}

func isValidForger(block Block) bool {
	return block.Forger == getNextForger()
}

func isBlockTransactionsValid(block Block) bool {
	transactions := block.Transactions
	covered_transactions := getCoveredTransactions(transactions)
	return len(transactions) == len(covered_transactions)
}

func getLatestPreviousHash() string {
	blockLenght := len(blockchain.Blocks)
	latestBlock := blockchain.Blocks[blockLenght-1]
	payload := getBlockPayload(latestBlock)

	return utils.GetHexadecimalHash(payload)
}

func getCoveredTransactions(transactions []Transaction) []Transaction {
	var coveredTransactions = make([]Transaction, len(transactions))

	for _, transaction := range transactions {
		if isTransactionCovered(transaction) {
			coveredTransactions = append(coveredTransactions, transaction)
		}
	}
	return coveredTransactions
}

func isTransactionCovered(transaction Transaction) bool {
	if transaction.TransactionType == EXCHANGE {
		return true
	}
	senderBalance := getAccountModelBalance(transaction.SenderPublicKey)
	return senderBalance >= transaction.Amount
}

func accountBalance(publicKey string) int {
	return getAccountModelBalance(publicKey)
}

func getNextForger() string {
	prevBlockHash := getLatestPreviousHash()
	return proofOfStake.PickForger(prevBlockHash)
}

//func createBlock(transactions []Transaction, forgerWallet interface{}) Block {
//	covered_transactions := self.get_covered_transactions(transactions)
//	self.execute_transactions(covered_transactions)
//	block := forgerWallet.create_block(
//		covered_transactions,
//		self.latest_previous_hash(),
//		self.latest_block_height()+1,
//	)
//	self.add_block(block)
//	return block
//}

func isTransactionInBlockchain(transaction Transaction) bool {
	for _, block := range blockchain.Blocks {
		for _, blockTransaction := range block.Transactions {
			if blockTransaction == transaction {
				return true
			}
		}
	}
	return false
}
