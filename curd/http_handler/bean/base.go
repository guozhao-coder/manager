package bean

type BaseRequest struct {
	Signature   string `json:"signature"`   //用户签名标记
	RequestTime int64  `json:"requestTime"` //请求时间
}

type BaseResponse struct {
	Code         int    `json:"code"`               //返回码
	Message      string `json:"message,omitempty"`  //错误信息
	ResponseTime int64  `json:"respTime,omitempty"` //响应时间
}

//响应参数
type DataResult struct {
	*BaseResponse
	Result interface{} `json:"result"`
}
