package api

import (
	"cryptomunt/networking"
	"cryptomunt/structs"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getBlockHeight(c *gin.Context, cryptoNode *networking.CryptoNode) {
	c.JSON(200, len(cryptoNode.Blockchain.Blocks))
}

//return blockchain blocks with start and end index
func getBlocks(c *gin.Context, cryptoNode *networking.CryptoNode, apiRequest chan structs.ApiCallMessage, apiResponse chan structs.ApiCallMessage) {
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

func getMemoryPool(c *gin.Context, cryptoNode *networking.CryptoNode) {
	c.JSON(200, cryptoNode.MemoryPool)
}
