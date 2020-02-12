package rw

import (
	bean2 "curd/http_handler/bean"
	"encoding/json"
	"net/http"

	logs "github.com/cihub/seelog"

	"smartroom_address/util"
)

//输出错误信息相关的响应
func WriteErrResp(code int, msg string, response http.ResponseWriter) {

	br := &bean2.BaseResponse{
		Code:         code,
		Message:      msg,
		ResponseTime: util.GetUnix13NowTime(),
	}
	bs, err := json.Marshal(br)
	if err != nil {
		logs.Error("writeErrResp json.Marshal error:" + err.Error())
		return
	}
	//	logs.Debug(string(bs))
	write(response, bs)
}

//输出正常的响应
func WriteResult(result interface{}, response http.ResponseWriter) {

	bs, err := json.Marshal(result)
	if err != nil {
		logs.Error("writeResult json.Marshal error:" + err.Error())
		WriteErrResp(99, err.Error(), response)
		return
	}

	logs.Debug(string(bs))
	write(response, bs)
}

func WriteDataResult(code int, msg string, data interface{}, response http.ResponseWriter) {
	dataResult := bean2.DataResult{
		BaseResponse: &bean2.BaseResponse{
			Code:         code,
			Message:      msg,
			ResponseTime: util.GetUnix13NowTime(),
		},
		Result: data,
	}
	bs, err := json.Marshal(dataResult)
	if err != nil {
		logs.Error("writeResult json.Marshal error:" + err.Error())
		WriteErrResp(99, err.Error(), response)
		return
	}

	logs.Debug(string(bs))
	write(response, bs)
}

func write(response http.ResponseWriter, data []byte) {
	response.Header().Add("Access-Control-Allow-Origin", "*")

	response.Header().Add("Access-Control-Request-Method", "*")
	response.Header().Add("Access-Control-Allow-Headers", "User-Agent,X-Requested-With,Cache-Control,Content-Type,"+
		"Access-Token,Authorization,range-token")
	response.Header().Add("Content-Type", "application/json")

	//	response.Header().Add("Access-Control-Allow-Headers", "User-Agent,X-Requested-With,Cache-Control,Content-Type,Access-Token,Authorization")
	response.Write(data)
	//	fmt.Fprintln(response, data)
}
