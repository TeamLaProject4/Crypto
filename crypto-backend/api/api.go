package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// album represents data about a record album.
// capitalize first letter of new var in struct to escape last line
type seedphrase struct {
	ID       string `json:"id"`
	Mnemonic string `json:"mnemonic"`
	Secret   string `json:"secret"`
}

var seedphrases = []seedphrase{}

//return blockchain blocks with start and end index
func getBlocks(c *gin.Context, cryptoNode networking.CryptoNode) {
	queryParameters := c.Request.URL.Query()
	start := queryParameters["start"]
	end := queryParameters["end"]

	if start != nil && end != nil {
		startInt, _ := strconv.Atoi(start[0])
		endInt, _ := strconv.Atoi(end[0])
		blocks := cryptoNode.Blockchain.GetBlocksFromRange(startInt, endInt)

		c.JSON(200, blocks)
		return
	}

	c.JSON(419, gin.H{
		"start": "ERROR: no parameters 'start' and/or 'end' found",
	})
}

func getBlockHeight(c *gin.Context, cryptoNode networking.CryptoNode) {
	c.JSON(200, len(cryptoNode.Blockchain.Blocks))
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
		"start": "ERROR: no parameters 'start' and/or 'end' found",
	})
}

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

func getMemoryPool(c *gin.Context, cryptoNode networking.CryptoNode) {
	//memoryPool := cryptoNode.MemoryPool
	c.JSON(200, "TEST")
	//c.JSON(200, memoryPool)
}
func test(c *gin.Context, cryptoNode networking.CryptoNode) {
	test := cryptoNode.Wallet.CreateTransaction("henk", 20, blockchain.TRANSFER)
	cryptoNode.MemoryPool.AddTransaction(test)

	memoryPool := cryptoNode.MemoryPool
	c.JSON(200, memoryPool)
}

func StartApi(cryptoNode networking.CryptoNode) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//routes for frontend communication
	router.GET("/frontend/balance", func(context *gin.Context) {
		getAccountBalance(context, cryptoNode)
	})
	router.GET("/frontend/transactions", func(context *gin.Context) {
		getAccountTransactions(context, cryptoNode)
	})

	router.GET("/test", func(context *gin.Context) {
		test(context, cryptoNode)
	})
	//routes for node communication
	router.GET("/blockchain/block-length", func(context *gin.Context) {
		getBlockHeight(context, cryptoNode)
	})

	router.GET("/blockchain/blocks", func(context *gin.Context) {
		getBlocks(context, cryptoNode)
	})
	router.GET("/blockchain/memory-pool", func(context *gin.Context) {
		getMemoryPool(context, cryptoNode)
	})

	nodeIpAddr := cryptoNode.GetOwnIpAddr()
	utils.Logger.Infof("Rest API %s", nodeIpAddr)
	err := router.Run(cryptoNode.GetOwnIpAddr())
	if err != nil {
		utils.Logger.Fatal("Failed to start rest api", err)
		return
	}
}
