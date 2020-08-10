package util

import "github.com/lyd2/live/pkg/code"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    code.SUCCESS,
		Message: code.GetMsg(code.SUCCESS),
		Data:    data,
	}
}

func Error(c int) *Response {
	return &Response{
		Code:    c,
		Message: code.GetMsg(c),
		Data:    nil,
	}
}

func ErrorMsg(c int, msg string) *Response {
	return &Response{
		Code:    c,
		Message: msg,
		Data:    nil,
	}
}

func ServerError() *Response {
	return &Response{
		Code:    code.ERROR,
		Message: code.GetMsg(code.ERROR),
		Data:    nil,
	}
}
