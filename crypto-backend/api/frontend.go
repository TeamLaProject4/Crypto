package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/wallet"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getAccountTransactions(c *gin.Context, cryptoNode networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	publicKey := queryParameters["publicKey"]
	if publicKey != nil {
		balance := cryptoNode.Blockchain.GetAllAccountTransactions(publicKey[0])
		c.JSON(200, balance)
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters 'start' and/or 'end' found",
	})
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

	//wallet.ConvertMnemonicToKeys(mnemonic.Mnemonic, "secret")
	node.Wallet.SetWalletKeyAndWritePrivateKeyFile(mnemonic.Mnemonic)

	//    fmt.Println(requestBody.Mnemonic)
	c.IndentedJSON(http.StatusCreated, "key files created")
}

func createTransaction(c *gin.Context, node networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	publicKey := queryParameters["publicKey"]
	amount := queryParameters["amount"]

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
