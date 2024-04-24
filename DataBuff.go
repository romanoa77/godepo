package main

import (
	"github.com/gin-gonic/gin"
)

type fstat struct {
	Id        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

type BufStat struct {
	n_itm int `json:"n_itm"`
}

var fpath = "data"

func main() {

	var FCache []fstat

	router := gin.Default()

	router.POST("/sendF", postFile(FCache))

	router.Run("localhost:8080")

}

// postAlbums adds an album from JSON received in the request body.
func postFile(Table []fstat) gin.HandlerFunc {
	var newFile []fstat
	copy(newFile, Table)

	return func(ctx *gin.Context) {

	}
	/*
	   // Call BindJSON to bind the received JSON to
	   // newAlbum.

	   	if err := c.BindJSON(&newFile); err != nil {
	   		return
	   	}

	   // Add the new album to the slice.
	   FTable = append(FTable, newFile)

	   c.IndentedJSON(http.StatusCreated, newFile)
	*/
}
