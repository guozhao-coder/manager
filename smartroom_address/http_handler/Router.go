package http_handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"smartroom_address/http_handler/http_process/controller"
)

func newRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("智慧机房地点管理"))
	})
	//地点管理
	router_address(router)
	return router
}

func router_address(router *httprouter.Router) {

	//查询地点
	router.GET("/smartroom/leaf/address/getaddress", controller.Query_Address)

	//删除地点
	router.DELETE("/smartroom/leaf/address/delete", controller.Delete_Address)

	//新增地点

	router.POST("/smartroom/leaf/address/add", controller.Insert_Address)

	//编辑地点
	router.PUT("/smartroom/leaf/address/update", controller.Update_Address)

}
