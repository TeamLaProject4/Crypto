package blockchain

type AccountModel struct {
	Balances map[string]int
}

//var accountModel = new(AccountModel)

func CreateAccountModel() AccountModel {
	//TODO: add later for multiple AccountModel instances
	accountModel := new(AccountModel)
	accountModel.Balances = make(map[string]int)
	return *accountModel
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
