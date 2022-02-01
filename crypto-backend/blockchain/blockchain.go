package blockchain

import (
	"cryptomunt/proofOfStake"
	"cryptomunt/utils"
	"encoding/json"
	"errors"
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
	amountWithoutFee := CalculateInitialAmount(amount)

	if transaction.Type == STAKE {
		if sender == receiver {
			blockchain.ProofOfStake.UpdateStake(sender, amountWithoutFee)
			blockchain.AccountModel.UpdateBalance(sender, -amount)
		}
	} else if transaction.Type == REWARD {
		blockchain.AccountModel.UpdateBalance(receiver, amount)
	} else {
		blockchain.AccountModel.UpdateBalance(sender, -amount)
		blockchain.AccountModel.UpdateBalance(receiver, amountWithoutFee)
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

func (blockchain *Blockchain) IsBlockRewardTransactionValid(block Block) bool {
	if !blockHasOneRewardTransaction(block) {
		return false
	}

	rewardTx, _ := getRewardTransactionFromBlock(block)
	return rewardTransactionHasCorrectReceiver(block, rewardTx) && rewardTransactionHasCorrectAmount(block, rewardTx)
}

func blockHasOneRewardTransaction(block Block) bool {
	rewardTxFound := false

	for _, transaction := range block.Transactions {
		if transaction.Type == REWARD && rewardTxFound {
			return false
		}
		if transaction.Type == REWARD {
			rewardTxFound = true
		}
	}

	return rewardTxFound
}

func getRewardTransactionFromBlock(block Block) (Transaction, error) {
	for _, transaction := range block.Transactions {
		if transaction.Type == REWARD {
			return transaction, nil
		}
	}
	return *new(Transaction), errors.New("No reward transaction in block")
}

func rewardTransactionHasCorrectReceiver(block Block, rewardTx Transaction) bool {
	return block.Forger == rewardTx.ReceiverPublicKey
}

func rewardTransactionHasCorrectAmount(block Block, rewardTx Transaction) bool {
	transactionsLength := len(block.Transactions)
	totalReward := CalculateTotalReward(block.Transactions[:transactionsLength-1])
	return rewardTx.Amount == totalReward
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
	if transaction.Type == EXCHANGE || transaction.Type == REWARD {
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
