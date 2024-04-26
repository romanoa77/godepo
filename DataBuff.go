package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type fstat struct {
	Id        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

type bufStat struct {
	N_itm int `json:"n_itm"`
}

type gw struct {
	H_data float64 `json:"h"`
	T_data float64 `json:"t"`
}

var fpath = "data/"

func main() {

	var FCache []fstat
	var Stats bufStat

	Stats.N_itm = 0

	router := gin.Default()

	router.POST("/sendF", postFile(FCache, &Stats))

	router.Run("localhost:8080")

}

// postAlbums adds an album from JSON received in the request body.
func postFile(Table []fstat, Stpt *bufStat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//First update service status

		Stpt.N_itm += 1
		var newFile fstat
		newFile.Id = Stpt.N_itm
		tm := time.Now()

		newFile.Timestamp = tm.String()

		Table = append(Table, newFile)

		var Gwbuf gw
		// Call BindJSON to bind the received JSON to
		// newAlbum.

		if err := ctx.BindJSON(&Gwbuf); err != nil {
			return
		}

		b, _ := json.Marshal(Gwbuf)
		filename := strconv.Itoa(Stpt.N_itm) + tm.String()
		os.WriteFile(fpath+filename, b, 0644)

		ctx.IndentedJSON(http.StatusCreated, Gwbuf)

	}

}
