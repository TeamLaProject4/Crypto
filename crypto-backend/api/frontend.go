package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAccountTransactions(c *gin.Context, cryptoNode *networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	publicKey := queryParameters["publicKey"]
	if publicKey != nil {
		accountTransaction := cryptoNode.Blockchain.GetAllAccountTransactions(publicKey[0])
		c.JSON(200, accountTransaction)
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters publicKey",
	})
}

func getAccountBalance(c *gin.Context, cryptoNode *networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	publicKey := queryParameters["publicKey"]
	if publicKey != nil {
		balance := cryptoNode.Blockchain.AccountModel.GetBalance(publicKey[0])
		c.JSON(200, balance)
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters publicKey found",
	})
}
func getOwnPublicKey(c *gin.Context, cryptoNode *networking.CryptoNode) {
	c.JSON(200, cryptoNode.Wallet.GetPublicKeyHex())
}
func getGenesisPublicKey(c *gin.Context, cryptoNode *networking.CryptoNode) {
	c.JSON(200, cryptoNode.Wallet.GetPublicKeyHex())
}

func getMnemonic(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.IndentedJSON(http.StatusOK, wallet.GenerateMnemonic())
}
func confirmMnemonic(c *gin.Context, node *networking.CryptoNode) {
	var req = c.Request
	var mnemonic Mnemonic
	setupResponse(c)

	if (*req).Method == "OPTIONS" {
		return
	}
	if err := c.BindJSON(&mnemonic); err != nil {
		return
	}

	//wallet.ConvertMnemonicToKeys(mnemonic.Mnemonic, "secret")
	masterKey := wallet.NewMasterKey(mnemonic.Mnemonic)
	ecdsaKey := wallet.ConvertBip32ToECDSA(masterKey)

	utils.Logger.Info("key", ecdsaKey)
	node.Wallet = wallet.CreateWallet(ecdsaKey)
	utils.Logger.Info("Node wallet", node.Wallet)

	ecdsaKeyPemEncoded := wallet.EncodePrivateKey(&ecdsaKey)
	wallet.WriteKeyToFile(wallet.PRIVATE_KEY_PATH, ecdsaKeyPemEncoded)

	c.IndentedJSON(http.StatusCreated, "key files created")
}

func createTransaction(c *gin.Context, node *networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	recieverPublicKey := queryParameters["recieverPublicKey"]
	amount := queryParameters["amount"]
	transactionType := queryParameters["transactionType"]

	if recieverPublicKey != nil && amount != nil && transactionType != nil {
		amountInt, err := strconv.Atoi(amount[0])
		recieverPublicKeyString := recieverPublicKey[0]
		if err != nil {
			c.JSON(419, "Amount is not an integer")
			return
		}
		//set trans type
		var transType blockchain.TransactionType
		if transactionType[0] == "transfer" {
			transType = blockchain.TRANSFER
		} else {
			recieverPublicKeyString = node.Blockchain.ProofOfStake.GenesisPublicKey
			transType = blockchain.STAKE
		}

		//create and verify transaction
		transaction := node.Wallet.CreateTransaction(recieverPublicKeyString, amountInt, transType)
		if !node.IsTransactionValid(transaction) {
			c.JSON(419, "Transaction is not valid")
			return
		}

		//memPool, validate & forge if threshold reached
		node.HandleTransaction(transaction)

		//write to topic
		node.WriteToTopic(transaction.ToJson(), networking.TRANSACTION)

		c.JSON(200, "added transaction to memoryPool")
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters recieverPublicKey, amount or transactionType found",
	})
}

/**
PLACEHOLDERS TILL THE REAL THING IS FIXED
TODO: REMOVE
*/
// func constructTransaction(pk string, rk string, amount int) blockchain.Transaction {
// 	return blockchain.Transaction{
// 		SenderPublicKey:   pk,
// 		ReceiverPublicKey: rk,
// 		Amount:            amount,
// 		Type:              blockchain.TRANSFER,
// 	}
// }
// func constructTransactions() []blockchain.Transaction {
// 	transaction1 := constructTransaction("lars", "jeroen", 20)
// 	transaction2 := constructTransaction("johan", "jeroen", 10)
// 	transaction3 := constructTransaction("martijn", "lars", 32)
// 	transaction4 := constructTransaction("martijn", "henk", 32)
// 	return []blockchain.Transaction{transaction1, transaction2, transaction3, transaction4}
// }
// func constructBlock(prevHash string) blockchain.Block {
// 	block := new(blockchain.Block)
// 	block.Transactions = constructTransactions()
// 	block.PreviousHash = prevHash
// 	block.Forger = "forger"
// 	block.Height = 1
// 	return *block
// }

// func TestWhenSetBalancesFromBlockchainThenBalanceHasCorrectAmount() {
// 	chain := blockchain.CreateBlockchain()
// 	block1 := constructBlock(chain.LatestPreviousHash())
// 	chain.Blocks = []blockchain.Block{block1}

// }

// func TestWhenGettingAllAccountLarsTransactionsThenIsSix() {
// 	chain := blockchain.CreateBlockchain()
// 	block1 := constructBlock(chain.LatestPreviousHash())
// 	block2 := constructBlock(chain.LatestPreviousHash())
// 	block3 := constructBlock(chain.LatestPreviousHash())
// 	chain.Blocks = []blockchain.Block{block1, block2, block3}

// }

//TODO: REMOVE CODE ABOVE WHEN DONE!!
