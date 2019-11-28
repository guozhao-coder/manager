package util

import (
	"fmt"

	logs "github.com/cihub/seelog"
)

//func LoadLogConfig() {
func init() {
	logger, err := logs.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	logs.ReplaceLogger(logger)
}