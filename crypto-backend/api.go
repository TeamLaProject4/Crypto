package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func main_router() {
	router := gin.Default()
	// router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
