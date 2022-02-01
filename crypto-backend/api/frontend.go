package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	//    fmt.Println(requestBody.Mnemonic)
	c.IndentedJSON(http.StatusCreated, "key files created")
}

func createTransaction(c *gin.Context, node *networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	recieverPublicKey := queryParameters["recieverPublicKey"]
	amount := queryParameters["amount"]

	if recieverPublicKey != nil && amount != nil {
		amountInt, err := strconv.Atoi(amount[0])
		if err != nil {
			c.JSON(419, "Amount is not an integer")
			return
		}

		transaction := node.Wallet.CreateTransaction(recieverPublicKey[0], amountInt, blockchain.TRANSFER)
		if !node.IsTransactionValid(transaction) {
			c.JSON(419, "Transaction is not valid")
			return
		}
		utils.Logger.Info("TRANSACTION VALID")

		//add transaction to mem pool
		node.MemoryPool.AddTransaction(transaction)
		//write to topic
		node.WriteToTopic(transaction.ToJson(), networking.TRANSACTION)

		//stake if threshold is reached
		if node.MemoryPool.IsTransactionThresholdReached() {

		}

		c.JSON(200, "added transaction to memoryPool")
		return
	}
	c.JSON(419, gin.H{
		"start": "ERROR: no parameters recieverPublicKey or amount found",
	})
}
