package blockchain

type AccountModel struct {
	Balances map[string]int
}

//var accountModel = new(AccountModel)

func NewAccountModel() AccountModel {
	//TODO: add later for multiple AccountModel instances
	//AccountModel = new(AccountModel)
	return AccountModel{Balances: *new(map[string]int)}
}

func (accountModel *AccountModel) IsAccountInBalances(publicKey string) bool {
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
