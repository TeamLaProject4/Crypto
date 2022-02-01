package api

import (
	"cryptomunt/networking"
	"cryptomunt/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func setupResponse(c *gin.Context) {
	var w = c.Writer

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func StartApi(cryptoNode *networking.CryptoNode) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	//routes for frontend communication
	router.GET("/frontend/publickey", func(context *gin.Context) {
		getOwnPublicKey(context, cryptoNode)
	})

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

	router.POST("/frontend/transaction", func(context *gin.Context) {
		createTransaction(context, cryptoNode)
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
