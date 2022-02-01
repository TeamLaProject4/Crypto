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
func getOwnPublicKey(c *gin.Context, cryptoNode *networking.CryptoNode) {
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
	utils.Logger.Info("in conf mnuemonic")
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

type TransactionParams struct {
	RecieverPublicKey string
	Amount            string
	TransactionType   string
}

func createTransaction(c *gin.Context, node *networking.CryptoNode) {
	var req = c.Request
	var params TransactionParams
	setupResponse(c)

	if (*req).Method == "OPTIONS" {
		return
	}
	if err := c.BindJSON(&params); err != nil {
		return
	}

	recieverPublicKey := params.RecieverPublicKey
	amount := params.Amount
	transactionType := params.TransactionType

	if recieverPublicKey != "" && amount != "" && transactionType != "" {
		amountInt, err := strconv.Atoi(amount)
		recieverPublicKeyString := recieverPublicKey
		if err != nil {
			c.JSON(419, "Amount is not an integer")
			return
		}
		//set trans type
		var transType blockchain.TransactionType
		if transactionType == "transfer" {
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
