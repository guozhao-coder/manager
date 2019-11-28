package rw

import (
	"encoding/json"
	"fmt"
	logs "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
	errCode "smartroom_address/http_handler/error_code"
)

func ReadBody(request *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Println(" ioutil.ReadAll is error........")
		return err
	}
	defer request.Body.Close()
	//fmt.Println(string(data))
	if data == nil|| len(data) == 0 {
		logs.Error(err)
		return errCode.New( errCode.JSON_UNMARSHAL_ERROR,"数据为空")
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		logs.Error(err)
		return errCode.New( errCode.JSON_UNMARSHAL_ERROR,"json解析错误")

	}
	return nil
}
