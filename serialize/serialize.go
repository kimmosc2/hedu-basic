package serialize

import (
	"hedu-basic/util/ecode"
)

// 标准响应体
type Response struct {
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	ECode int         `json:"ecode"`
}

// 通用序列化响应
func BuildAutoResponse(code int, data interface{}) *Response {
	return &Response{
		Data:  data,
		Msg:   ecode.Ecode[code],
		ECode: code,
	}
}

// 错误序列化响应
func BuildErrorResponse(code int) *Response {
	return &Response{
		Data:  nil,
		Msg:   ecode.Ecode[code],
		ECode: code,
	}
}
