package errCode

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"net/http"
	"smartroom_address/util"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h HttpError) Error() string {
	strFormat := `
	error code: %d
	error message %s
`
	return fmt.Sprintf(strFormat, h.Code, h.Message)
}

func New(code int, message string) error {
	return &HttpError{code, message}
}

//如果 err != nil 返回true
func CheckError(err error, response http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	seelog.Errorf("发生错误:%s", err.Error())
	httpError, ok := err.(HttpError)
	if ok {
		writeErrResp(httpError.Code, httpError.Message, response)
	} else {
		writeErrResp(SERVER_ERR, "服务器繁忙", response)
	}
	return true
}
func writeErrResp(code int, errMsg string, response http.ResponseWriter) {

	rm := make(map[string]interface{})
	rm["code"] = code
	rm["message"] = errMsg
	rm["responseTime"] = util.GetUnix13NowTime()

	bs, err := json.Marshal(rm)
	if err != nil {
		seelog.Error("writeErrResp json.Marshal error:" + err.Error())
		return
	}
	//	logs.Debug(string(bs))
	write(response, bs)
}
func write(response http.ResponseWriter, data []byte) {
	response.Header().Add("Access-Control-Allow-Origin", "*")
	response.Write(data)
}
