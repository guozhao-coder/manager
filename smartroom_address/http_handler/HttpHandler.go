package http_handler

import (
	"fmt"
	logs "github.com/cihub/seelog"
	"net/http"
)

func InitServlet(port int){
	defer  func() {
		if err := recover(); err != nil {
			logs.Error("initServlet error:", err)
		}
	}()

	router := newRouter()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	err := server.ListenAndServe()
	logs.Debug("listenning port...")
	if err != nil{
		logs.Error("listen server error")
	}
}