package api

import (
	"cryptomunt/blockchain"
	"cryptomunt/networking"
	"cryptomunt/utils"
	"cryptomunt/wallet"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type seedphrase struct {
	ID       string `json:"id"`
	Mnemonic string `json:"mnemonic"`
	Secret   string `json:"secret"`
}

type Mnemonic struct {
	Mnemonic string `json:"mnemonic"`
}

var seedphrases = []seedphrase{}

// getAlbums responds with the list of all albums as JSON.
func getMnemonic(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.IndentedJSON(http.StatusOK, wallet.GenerateMnemonic())
}

func setupResponse(c *gin.Context) {
	var w = c.Writer

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

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
	c.JSON(200, cryptoNode.MemoryPool)
}
func test(c *gin.Context, cryptoNode networking.CryptoNode) {
	test := cryptoNode.Wallet.CreateTransaction("henk", 20, blockchain.TRANSFER)
	cryptoNode.MemoryPool.AddTransaction(test)

	memoryPool := cryptoNode.MemoryPool
	c.JSON(200, memoryPool)
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

func StartApi(cryptoNode networking.CryptoNode) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	//routes for frontend communication
	router.GET("/frontend/balance", func(context *gin.Context) {
		getAccountBalance(context, cryptoNode)
	})
	router.GET("/frontend/transactions", func(context *gin.Context) {
		getAccountTransactions(context, cryptoNode)
	})
	router.GET("/frontend/getMnemonic", getMnemonic)
	router.POST("/frontend/confirmMnemonic", func(context *gin.Context) {
		confirmMnemonic(context, cryptoNode)
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
	router.GET("/test", func(context *gin.Context) {
		test(context, cryptoNode)
	})

	nodeIpAddr := cryptoNode.GetOwnIpAddr()
	utils.Logger.Infof("Rest API %s", nodeIpAddr)
	err := router.Run(cryptoNode.GetOwnIpAddr())
	if err != nil {
		utils.Logger.Fatal("Failed to start rest api", err)
		return
	}
}
