package http_handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"smartroom_address/http_handler/http_process/controller"
)

func newRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	})
	//地点管理
	router_address(router)
	return router
}

func router_address(router *httprouter.Router) {

	router.GET("/smartroom/leaf/address/getaddress", controller.Query_Address)

	router.DELETE("/smartroom/leaf/address/delete", controller.Delete_Address)

	router.POST("/smartroom/leaf/address/add", controller.Insert_Address)

	router.PUT("/smartroom/leaf/address/update", controller.Update_Address)

}