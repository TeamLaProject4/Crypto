package api

import (
	"cryptomunt/networking"
	"cryptomunt/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album represents data about a record album.
// capitalize first letter of new var in struct to escape last line
type seedphrase struct {
	ID       string `json:"id"`
	Mnemonic string `json:"mnemonic"`
	Secret   string `json:"secret"`
}

// // albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

var seedphrases = []seedphrase{}

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	c.JSON(200, gin.H{
		"start": "ERROR: no parameters 'start' and/or 'end' found",
	})
	return

	var newSeedphrase seedphrase

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newSeedphrase); err != nil {
		return
	}

	// Add the new album to the slice.
	seedphrases = append(seedphrases, newSeedphrase)
	c.IndentedJSON(http.StatusCreated, seedphrases)
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

func StartApi(cryptoNode networking.CryptoNode) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/albums", postAlbums)

	//routes for frontend communication
	router.GET("/frontend/balance", func(context *gin.Context) {
		getAccountBalance(context, cryptoNode)
	})
	router.GET("/frontend/transactions", func(context *gin.Context) {
		getAccountTransactions(context, cryptoNode)
	})

	//routes for node communication
	router.GET("/blockchain/block-length", func(context *gin.Context) {
		getBlockHeight(context, cryptoNode)
	})

	router.GET("/blockchain/blocks", func(context *gin.Context) {
		getBlocks(context, cryptoNode)
	})

	nodeIpAddr := cryptoNode.GetOwnIpAddr()
	utils.Logger.Infof("Rest API %s", nodeIpAddr)
	err := router.Run(cryptoNode.GetOwnIpAddr())
	if err != nil {
		utils.Logger.Fatal("Failed to start rest api", err)
		return
	}
}
