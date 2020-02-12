package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(fileDir string, data interface{}) interface{} {
	r, err := os.Open(fileDir)
	if err != nil {
		fmt.Println("ReadJsonFile open json file error:" + err.Error())
	}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("ReadJsonFile json decode error:" + err.Error())
	}
	return data
}
