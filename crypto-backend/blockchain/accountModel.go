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

func (accountModel *AccountModel) isAccountInBalances(publicKey string) bool {
	_, accountInBalances := accountModel.Balances[publicKey]
	return accountInBalances
}

func (accountModel *AccountModel) addAccount(publicKey string) {
	if !accountModel.isAccountInBalances(publicKey) {
		accountModel.Balances[publicKey] = 0
	}
}

func (accountModel *AccountModel) getBalance(publicKey string) int {
	accountModel.addAccount(publicKey)
	return accountModel.Balances[publicKey]
}

func (accountModel *AccountModel) updateBalance(publicKey string, amount int) {
	accountModel.addAccount(publicKey)
	accountModel.Balances[publicKey] += amount
}
