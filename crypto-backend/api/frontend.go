package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/wallet"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAccountTransactions(c *gin.Context, cryptoNode networking.CryptoNode) {
	trans := constructTransactions()
	c.JSON(200, trans)
	// queryParameters := c.Request.URL.Query()
	// publicKey := queryParameters["publicKey"]
	// if publicKey != nil {
	// 	// balance := cryptoNode.Blockchain.GetAllAccountTransactions(publicKey[0])
	// 	return
	// }
	// c.JSON(419, gin.H{
	// 	"start": "ERROR: no parameters 'start' and/or 'end' found",
	// })
}

func getAccountBalance(c *gin.Context, cryptoNode networking.CryptoNode) {
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
func getMnemonic(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.IndentedJSON(http.StatusOK, wallet.GenerateMnemonic())
}
func confirmMnemonic(c *gin.Context, node networking.CryptoNode) {
	var req = c.Request
	var mnemonic Mnemonic
	setupResponse(c)

	if (*req).Method == "OPTIONS" {
		return
	}
	if err := c.BindJSON(&mnemonic); err != nil {
		return
	}

	// Add the new album to the slice.
	fmt.Println(mnemonic.Mnemonic)

	masterKey := wallet.NewMasterKey(mnemonic.Mnemonic)
	ecdsaKey := wallet.ConvertBip32ToECDSA(masterKey)
	node.Wallet = wallet.CreateWallet(ecdsaKey)
	wallet.WriteKeyToFile(wallet.PRIVATE_KEY_PATH, ecdsaKey)

	c.IndentedJSON(http.StatusCreated, "key files created")
}

func createTransaction(c *gin.Context, node networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	publicKey := queryParameters["publicKey"]
	amount := queryParameters["amount"]
	constructTransactions()

	if publicKey != nil && amount != nil {
		amountInt, err := strconv.Atoi(amount[0])
		if err != nil {
			c.JSON(419, "Amount is not an integer")
		}

		node.Wallet.CreateTransaction(publicKey[0], amountInt, blockchain.TRANSFER)
		c.JSON(200, "added transaction to memoryPool")
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters publicKey or amount found",
	})
}

/**
PLACEHOLDERS TILL THE REAL THING IS FIXED
TODO: REMOVE
*/
func constructTransaction(pk string, rk string, amount int) blockchain.Transaction {
	return blockchain.Transaction{
		SenderPublicKey:   pk,
		ReceiverPublicKey: rk,
		Amount:            amount,
		Type:              blockchain.TRANSFER,
	}
}
func constructTransactions() []blockchain.Transaction {
	transaction1 := constructTransaction("lars", "jeroen", 20)
	transaction2 := constructTransaction("johan", "jeroen", 10)
	transaction3 := constructTransaction("martijn", "lars", 32)
	transaction4 := constructTransaction("martijn", "henk", 32)
	return []blockchain.Transaction{transaction1, transaction2, transaction3, transaction4}
}
func constructBlock(prevHash string) blockchain.Block {
	block := new(blockchain.Block)
	block.Transactions = constructTransactions()
	block.PreviousHash = prevHash
	block.Forger = "forger"
	block.Height = 1
	return *block
}

func TestWhenSetBalancesFromBlockchainThenBalanceHasCorrectAmount() {
	chain := blockchain.CreateBlockchain()
	block1 := constructBlock(chain.LatestPreviousHash())
	chain.Blocks = []blockchain.Block{block1}

}

func TestWhenGettingAllAccountLarsTransactionsThenIsSix() {
	chain := blockchain.CreateBlockchain()
	block1 := constructBlock(chain.LatestPreviousHash())
	block2 := constructBlock(chain.LatestPreviousHash())
	block3 := constructBlock(chain.LatestPreviousHash())
	chain.Blocks = []blockchain.Block{block1, block2, block3}

}

//TODO: REMOVE CODE ABOVE WHEN DONE!!
