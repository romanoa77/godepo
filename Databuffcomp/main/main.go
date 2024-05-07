package main

import (
	"encoding/json"
	"net/http"

	"base.url/class/fbufstat"
	"base.url/class/fwrite"
	"github.com/gin-gonic/gin"
)

/*
type fstat struct {
	Id        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

*/

type gw struct {
	H_data float64 `json:"h"`
	T_data float64 `json:"t"`
}

const baseadm = "adm/"
const baseadmn = "StatDesc.json"

func main() {

	StatDesc := fbufstat.NewObj(fwrite.UnFtoStrm(baseadm, baseadmn))

	router := gin.Default()

	router.GET("/stat", getStat(StatDesc.GetObj()))

	router.POST("/sendF", postFile(&StatDesc))

	router.Run("localhost:8080")

	/*
		StatDesc := fbufstat.New(0, 0)
		var buff []byte

		var dataset []fstat

		datast := []fstat{{Id: 0, Timestamp: "aaa"}, {Id: 1, Timestamp: "bbb"}, {Id: 2, Timestamp: "ccc"}}

		// Get a greeting message and print it.

		err := fwrite.PrintStToF("data/", "foo0", 0644, datast)

		buff, err = fwrite.FtoStrm("data/", "foo0")

		if err != nil {
			fmt.Println(err)

		} else {

			fmt.Println(string(buff))
			json.Unmarshal(buff, &dataset)
			fmt.Println(dataset)

			StatDesc.UpdateCnt()
			StatDesc.UpdateSize(128)
			fwrite.PrintStrmToF("adm/", "StatDesc.json", 0644, StatDesc.GetJSONObj())
			simplelogger.LogWriteFile("log/", "LogStream.json", 0, 128)
		}

	*/
}

func getStat(Table any) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		ctx.IndentedJSON(http.StatusOK, Table)

	}

}

func postFile(Stpt *fbufstat.Bufstat) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var Gwbuf gw

		// Call BindJSON to bind the received JSON to
		// newAlbum.

		if err := ctx.BindJSON(&Gwbuf); err != nil {
			return
		}

		b, _ := json.Marshal(Gwbuf)

		//filename := strconv.Itoa(Stpt.N_itm) + tm.String()
		//os.WriteFile(fpath+filename, b, 0644)

		ctx.IndentedJSON(http.StatusCreated, Gwbuf)

	}

}
