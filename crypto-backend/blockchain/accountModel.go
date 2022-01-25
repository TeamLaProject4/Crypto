package blockchain

type AccountModel struct {
	Balances map[string]int
}

var accountModel = new(AccountModel)

func NewAccountModel() {
	//TODO: add later for multiple AccountModel instances
	//AccountModel = new(AccountModel)
}

func isAccountInBalances(publicKey string) bool {
	_, accountInBalances := accountModel.Balances[publicKey]
	return accountInBalances
}

func addAccount(publicKey string) {
	if !isAccountInBalances(publicKey) {
		accountModel.Balances[publicKey] = 0
	}
}

func getAccountModelBalance(publicKey string) int {
	addAccount(publicKey)
	return accountModel.Balances[publicKey]
}

func updateAccountModelBalance(publicKey string, amount int) {
	addAccount(publicKey)
	accountModel.Balances[publicKey] += amount
}
