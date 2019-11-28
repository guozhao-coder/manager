package controller

import (
	logs "github.com/cihub/seelog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"smartroom_address/http_handler/bean"
	errCode "smartroom_address/http_handler/error_code"
	"smartroom_address/http_handler/http_process/service"
	rw "smartroom_address/http_handler/http_read_write"
)


func Delete_Address(response http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	logs.Info("删除请求")
	DeleteAddr := new(bean.DelAddr)
	err := rw.ReadBody(request, DeleteAddr)
	if err != nil{
		rw.WriteErrResp(errCode.JSON_UNMARSHAL_ERROR,err.Error(),response)
		return
	}
	err = service.Delete_Address(*DeleteAddr)
	if err != nil{
		rw.WriteErrResp(errCode.DELETE_ERR,err.Error(),response)
		logs.Info(err)
		return
	}
	rw.WriteErrResp(errCode.SUCCESS,"删除成功",response)
	logs.Info("删除成功")
	return
}

func Insert_Address(response http.ResponseWriter, request *http.Request, ps httprouter.Params){
	logs.Info("添加请求")
	InsertAddr := new(bean.InsAddr)
	err := rw.ReadBody(request,InsertAddr)

	if err != nil{
		rw.WriteErrResp(errCode.JSON_UNMARSHAL_ERROR,err.Error(),response)
		return
	}
	//调用service的添加
	err = service.Insert_Address(*InsertAddr)
	if err != nil{
		rw.WriteErrResp(errCode.ADD_ERROR,err.Error(),response)
		logs.Error("新增失败")
		return
	}
	rw.WriteErrResp(errCode.SUCCESS,"新增成功",response)
	logs.Info("新增成功")
	return
}

func Update_Address(response http.ResponseWriter, request *http.Request, ps httprouter.Params)  {
	logs.Info("修改操作")
	UpdateAddr := new(bean.InsAddr)
	err := rw.ReadBody(request, UpdateAddr)
	if err != nil{
		rw.WriteErrResp(errCode.JSON_UNMARSHAL_ERROR,err.Error(),response)
	}
	err = service.Update_Address(*UpdateAddr)
	if err !=nil{
		rw.WriteErrResp(errCode.UPDATE_ERR,err.Error(),response)
		logs.Error(err)
		return
	}
	rw.WriteErrResp(errCode.SUCCESS,"修改成功",response)
	logs.Info("修改成功")
	return
}

func Query_Address(response http.ResponseWriter, request *http.Request, ps httprouter.Params)  {
	logs.Info("查询操作")
	QueryAddr := new(bean.QueAddr)
	err := rw.ReadBody(request,QueryAddr)
	if err != nil{
		rw.WriteErrResp(errCode.JSON_UNMARSHAL_ERROR,err.Error(),response)
		return
	}
	addrs, err := service.Query_Address(*QueryAddr)
	if err != nil{
		logs.Error("查询失败")
		rw.WriteErrResp(errCode.FAILURE,err.Error(),response)
		return
	}
	rw.WriteDataResult(errCode.SUCCESS,"查询成功",addrs,response)
	logs.Info("查询成功")
}
