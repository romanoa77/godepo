package main

import (
	"encoding/json"
	"fmt"

	"base.url/class/fbufstat"
	"base.url/class/fwrite"
)

type fstat struct {
	Id        int    `json:"id"`
	Timestamp string `json:"timestamp"`
}

func main() {

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

	}
}
