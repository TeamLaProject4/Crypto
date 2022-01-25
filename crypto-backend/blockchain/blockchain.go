package blockchain

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"encoding/json"
	//"github.com/btcsuite/btcd/blockchain"
)

type Blockchain struct {
	Blocks       []Block
	AccountModel *AccountModel
	ProofOfStake *proofOfStake.ProofOfStake
}

//func GetBlockChain() *Blockchain {
//	return blockchain
//}
func NewBlockchain() Blockchain {
	genesisBlock := createGenesisBlock()
	var blocks []Block
	blocks = append(blocks, genesisBlock)
	pos := proofOfStake.NewProofOfStake()

	accountModel := NewAccountModel()

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

func (blockchain *Blockchain) getLatestBlockHeight() int {
	return len(blockchain.Blocks) - 1
}

func (blockchain *Blockchain) addBlock(block Block) {
	if blockchain.isValidBlockHeight(block) && blockchain.isValidPreviousBlockHash(block) {
		blockchain.executeTransactions(block.Transactions)
		blockchain.Blocks = append(blockchain.Blocks, block)
	}
}

func (blockchain *Blockchain) executeTransactions(transactions []Transaction) {
	for _, transaction := range transactions {
		blockchain.executeTransaction(transaction)
	}
}

func (blockchain *Blockchain) executeTransaction(transaction Transaction) {
	sender := transaction.SenderPublicKey
	receiver := transaction.ReceiverPublicKey
	amount := transaction.Amount

	if transaction.TransactionType == STAKE {
		if sender == receiver {

			blockchain.ProofOfStake.UpdateStake(sender, amount)
			blockchain.AccountModel.updateBalance(sender, -amount)
		}
	} else {
		blockchain.AccountModel.updateBalance(sender, -amount)
		blockchain.AccountModel.updateBalance(receiver, amount)
	}
}

func (blockchain *Blockchain) isValidBlockHeight(block Block) bool {
	blockLength := len(blockchain.Blocks)
	return blockchain.Blocks[blockLength-1].Height == block.Height-1
}

func (blockchain *Blockchain) isValidPreviousBlockHash(block Block) bool {
	return blockchain.getLatestPreviousHash() == block.PreviousHash
}

func (blockchain *Blockchain) isValidForger(block Block) bool {
	return block.Forger == blockchain.getNextForger()
}

func (blockchain *Blockchain) isBlockTransactionsValid(block Block) bool {
	transactions := block.Transactions
	coveredTransactions := blockchain.getCoveredTransactions(transactions)
	return len(transactions) == len(coveredTransactions)
}

func (blockchain *Blockchain) getLatestPreviousHash() string {
	blockLenght := len(blockchain.Blocks)
	latestBlock := blockchain.Blocks[blockLenght-1]
	payload := latestBlock.getPayload()

	return utils.GetHexadecimalHash(payload)
}

func (blockchain *Blockchain) getCoveredTransactions(transactions []Transaction) []Transaction {
	var coveredTransactions = make([]Transaction, len(transactions))

	for _, transaction := range transactions {
		if blockchain.isTransactionCovered(transaction) {
			coveredTransactions = append(coveredTransactions, transaction)
		}
	}
	return coveredTransactions
}

func (blockchain *Blockchain) isTransactionCovered(transaction Transaction) bool {
	if transaction.TransactionType == EXCHANGE {
		return true
	}
	senderBalance := blockchain.AccountModel.getBalance(transaction.SenderPublicKey)
	return senderBalance >= transaction.Amount
}

func (blockchain *Blockchain) accountBalance(publicKey string) int {
	return blockchain.AccountModel.getBalance(publicKey)
}

func (blockchain *Blockchain) getNextForger() string {
	prevBlockHash := blockchain.getLatestPreviousHash()
	return blockchain.ProofOfStake.PickForger(prevBlockHash)
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

func (blockchain *Blockchain) isTransactionInBlockchain(transaction Transaction) bool {
	for _, block := range blockchain.Blocks {
		for _, blockTransaction := range block.Transactions {
			if blockTransaction == transaction {
				return true
			}
		}
	}
	return false
}
