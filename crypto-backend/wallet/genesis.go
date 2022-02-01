package wallet

import (
	"cryptomunt/blockchain"
	"cryptomunt/utils"
	"fmt"
)

func CreateGenesisTransactions() []blockchain.Transaction {
	//key ophalen
	mnemonicBytes := utils.ReadFileBytes("./keys/demo-keys/wallet-mnemonic-genesis.txt")
	mnemonic := string(mnemonicBytes)

	bipKey := NewMasterKey(mnemonic)
	key := ConvertBip32ToECDSA(bipKey)
	wal := CreateWallet(key)

	transactions := make([]blockchain.Transaction, 0)
	for i := 0; i < 10; i++ {
		pubKey := string(utils.ReadFileBytes(fmt.Sprintf("./keys/demo-keys/wallet-pubkey-%d.txt", i)))
		utils.Logger.Info("pubkey", pubKey)

		trans := wal.CreateTransaction(pubKey, 10000, blockchain.TRANSFER)
		transactions = append(transactions, trans)
	}
	return transactions

}
