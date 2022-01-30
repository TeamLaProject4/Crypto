package blockchain

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"encoding/json"
	//"github.com/btcsuite/btcd/blockchain"
)

type Blockchain struct {
	Blocks       []Block                    `json:"blocks"`
	AccountModel *AccountModel              `json:"-"`
	proofOfStake *proofOfStake.ProofOfStake `json:"-"`
}

//func GetBlockChain() *Blockchain {
//	return blockchain
//}
func CreateBlockchain() Blockchain {
	genesisBlock := CreateGenesisBlock()
	var blocks []Block
	blocks = append(blocks, genesisBlock)
	pos := proofOfStake.NewProofOfStake()

	accountModel := CreateAccountModel()

	return Blockchain{
		Blocks:       blocks,
		AccountModel: &accountModel,
		proofOfStake: &pos,
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
			blockchain.proofOfStake.UpdateStake(sender, amount)
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

func (blockchain *Blockchain) isValidForger(block Block) bool {
	return block.Forger == blockchain.getNextForger()
}

func (blockchain *Blockchain) isBlockTransactionsValid(block Block) bool {
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
	return blockchain.proofOfStake.PickForger(prevBlockHash)
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
