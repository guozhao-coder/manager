package main

import (
	logs "github.com/cihub/seelog"
	ht "smartroom_address/http_handler"
)

const HTTP_PORT  = 9999

func main() {
	defer logs.Flush()
	defer func() {
		if err := recover(); err != nil {
			logs.Critical(err)
		}
	}()


	ht.InitServlet(HTTP_PORT)
}


