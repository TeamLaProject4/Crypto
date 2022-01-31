package main

import (
	"cryptomunt/wallet"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
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

type Mnemonic struct {
	Mnemonic string `json:"mnemonic"`
}

// // albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

var seedphrases = []seedphrase{}

// getAlbums responds with the list of all albums as JSON.
func getMnemonic(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.IndentedJSON(http.StatusOK, wallet.GenerateMnemonic())
}

// postAlbums adds an album from JSON received in the request body.
func confirmMnemonic(c *gin.Context) {

}

func setupResponse(c *gin.Context) {
	var w = c.Writer

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(c *gin.Context) {
	var req = c.Request

	setupResponse(c)
	if (*req).Method == "OPTIONS" {
		return
	}

	var mnemonic Mnemonic

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&mnemonic); err != nil {
		return
	}

	// Add the new album to the slice.
	fmt.Println(mnemonic.Mnemonic)

	wallet.ConvertMnemonicToKeys(mnemonic.Mnemonic, "secret")

	//    fmt.Println(requestBody.Mnemonic)
	c.IndentedJSON(http.StatusCreated, "key files created")
}

func main_router() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/getMnemonic", getMnemonic)
	router.POST("/confirmMnemonic", indexHandler)
	//	router.POST("/confirmMnemonic", confirmMnemonic)

	router.Run("localhost:8080")
}
