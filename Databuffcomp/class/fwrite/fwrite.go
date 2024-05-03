package fwrite

/*
The scope of this class is to provide methods for read/write streams
of bytes. Every operation i/o operation should be atomical

i.e. I receive a stream from endpoints. Write a new file locally
i.e Reading single file
*/

import (
	"encoding/json"
	"io/fs"
	"os"
)

// If the func literal starts with an uppercase letter
// the function will be public
func StructToByte(buf any) []byte {

	bt_arr, _ := json.Marshal(buf)

	return bt_arr

}

func PrintStToF(dest string, filename string, mod fs.FileMode, DtStr any) error {

	var code error

	buf := StructToByte(DtStr)

	//0644 read-only mod for other users.
	//0666 THE NUMBER OF THE BEAST! Read and write for all!
	//Remember WriteFile return nil with success.
	code = os.WriteFile(dest+filename, buf, mod)

	return code

}

func PrintStrmToF(dest string, filename string, mod fs.FileMode, DtStr []byte) error {

	//0644 read-only mod for other users.
	//0666 THE NUMBER OF THE BEAST! Read and write for all!
	//Remember WriteFile return nil with success.
	code := os.WriteFile(dest+filename, DtStr, mod)

	return code

}

func FtoStrm(dest string, filename string) ([]byte, error) {

	var code error
	var bt_arr []byte

	bt_arr, code = os.ReadFile(dest + filename)

	return bt_arr, code

}
