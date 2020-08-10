package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lyd2/live/pkg/code"
)

var Validate = validator.New()

// 验证所有字段
func Validator(s interface{}) *Response {

	// 验证结构体 s
	if err := Validate.Struct(s); err != nil {
		//fmt.Println(err.(validator.ValidationErrors))
		// 尝试转为 validator.ValidationErrors 类型
		paramErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrorMsg(code.ERROR, err.Error())
		}

		// 转换成功则返回第一条错误消息
		return ErrorMsg(code.ERROR, fmt.Sprintf("%v", paramErr[0]))
	}

	return nil
}

// 忽略 fields
func ValidatorExcept(s interface{}, fields ...string) *Response {

	// 验证结构体 s
	if err := Validate.StructExcept(s, fields...); err != nil {
		//fmt.Println(err.(validator.ValidationErrors))
		// 尝试转为 validator.ValidationErrors 类型
		paramErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrorMsg(code.ERROR, err.Error())
		}

		// 转换成功则返回第一条错误消息
		return ErrorMsg(code.ERROR, fmt.Sprintf("%v", paramErr[0]))
	}

	return nil
}

// 只验证 fields
func ValidatorPartial(s interface{}, fields ...string) *Response {

	// 验证结构体 s
	if err := Validate.StructPartial(s, fields...); err != nil {
		//fmt.Println(err.(validator.ValidationErrors))
		// 尝试转为 validator.ValidationErrors 类型
		paramErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return ErrorMsg(code.ERROR, err.Error())
		}

		// 转换成功则返回第一条错误消息
		return ErrorMsg(code.ERROR, fmt.Sprintf("%v", paramErr[0]))
	}

	return nil
}