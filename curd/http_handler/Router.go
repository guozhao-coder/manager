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
	//
	router.GET("/smartroom/leaf/address/getaddress", controller.Query_Address)

	router.DELETE("/smartroom/leaf/address/delete", controller.Delete_Address)

	router.POST("/smartroom/leaf/address/add", controller.Insert_Address)

	router.PUT("/smartroom/leaf/address/update", controller.Update_Address)

	//添加post
	router.Post("/smartroom/leaf/address/update", controller.Update_Address)
	//添加post2
	router.Post("/smartroom/leaf/address/update", controller.Update_Address)
	//添加post3（branch1做的操作）
	router.Post("/smartroom/leaf/address/update", controller.Update_Address)
	//添加post4（branch1做的操作）
	router.Post("/smartroom/leaf/address/update", controller.Update_Address)
	//branch1开发新功能
	router.GET("/smartroom/leaf/address/update", controller.Update_Address)
	//添加post5（branch2做的操作）
	router.Post("/smartroom/leaf/address/update", controller.Update_Address)
	//branch2开发新功能
	router.POST("/smartroom/leaf/address/update", controller.Update_Address)

	//branch2修改人
	router.POST("/smartroom/leaf/address/update", controller.Update_people)
	//branch2修改狗
	router.POST("/smartroom/leaf/address/update", controller.Update_dog)

}
