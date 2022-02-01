package blockchain

import (
	"sync"
)

// AccountModel cache for getting balances account
type AccountModel struct {
	Balances map[string]int
}

func CreateAccountModel() AccountModel {
	//TODO: extra multiple AccountModel instances?
	accountModel := new(AccountModel)
	accountModel.Balances = make(map[string]int)
	return *accountModel
}

func (accountModel *AccountModel) SetBalancesFromBlockChain(blockchain Blockchain) {
	//clear old balances
	accountModel.Balances = make(map[string]int)

	//calculate balances per block
	//use multiple go routines(threads) to increase speed
	var wg sync.WaitGroup
	blocksBalances := make(chan map[string]int)
	for _, block := range blockchain.Blocks {
		wg.Add(1)
		go func(block Block) {
			defer wg.Done()
			go getBalancesFromBlock(block, blocksBalances)
		}(block)
	}
	wg.Wait()

	//combine blocksBalances to one balance
	blockBalancesIndex := 0
	for blockBalances := range blocksBalances {
		for balancePubKey, balanceAmount := range blockBalances {
			accountModel.UpdateBalance(balancePubKey, balanceAmount)
		}
		if len(blockchain.Blocks)-1 == blockBalancesIndex {
			close(blocksBalances)
		}

		blockBalancesIndex++
	}
}

func getBalancesFromBlock(block Block, balances chan map[string]int) {
	balancesPerBlock := make(map[string]int)
	for _, transaction := range block.Transactions {
		balancesPerBlock[transaction.SenderPublicKeyString] -= transaction.Amount
		balancesPerBlock[transaction.ReceiverPublicKey] += transaction.Amount
	}
	balances <- balancesPerBlock
}

func (accountModel *AccountModel) IsAccountInBalances(publicKey string) bool {
	//TODO: check blockchain if account exists?
	_, accountInBalances := accountModel.Balances[publicKey]
	return accountInBalances
}

func (accountModel *AccountModel) AddAccount(publicKey string) {
	if !accountModel.IsAccountInBalances(publicKey) {
		accountModel.Balances[publicKey] = 0
	}
}

func (accountModel *AccountModel) GetBalance(publicKey string) int {
	accountModel.AddAccount(publicKey)
	return accountModel.Balances[publicKey]
}

func (accountModel *AccountModel) UpdateBalance(publicKey string, amount int) {
	accountModel.AddAccount(publicKey)
	accountModel.Balances[publicKey] += amount
}
