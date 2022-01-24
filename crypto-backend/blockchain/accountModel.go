package blockchain

type AccountModel struct {
	balances map[string]int
}

var accountModel = new(AccountModel)

func NewAccountModel() {
	//TODO: add later for multiple accountModel instances
	//accountModel = new(AccountModel)
}

func isAccountInBalances(publicKey string) bool {
	_, accountInBalances := accountModel.balances[publicKey]
	return accountInBalances
}

func addAccount(publicKey string) {
	if !isAccountInBalances(publicKey) {
		accountModel.balances[publicKey] = 0
	}
}

func getBalance(publicKey string) int {
	addAccount(publicKey)
	return accountModel.balances[publicKey]
}

func updateBalance(publicKey string, amount int) {
	addAccount(publicKey)
	accountModel.balances[publicKey] += amount
}
