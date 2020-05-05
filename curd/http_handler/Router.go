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

	//在brant上新增功能
	router.POST("/smartroom/leaf/address/add", controller.Insert_brant)

	//郭朝新接到一个需求，在本地创建一个自己的分支//发现问题并修改
	router.POST("/smartroom/leaf/address/add", controller.Insert-guozhao-update)
	//开发完毕提交代码//首先提交到自己的远程分支

	//吴靓接到一个任务，并创建一个自己的分支
	router.POST("/smartroom/leaf/address/add", controller.Insert-wuliang)

	//郭朝接到一个新需求，又写一个接口
	router.POST("/smartroom/leaf/address/add", controller.Insert-guozhao2-update)
}
