package code

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrSuccess = &Error{
		Code:    0,
		Message: "success",
	}
	ErrSystem = &Error{
		Code:    -1,
		Message: "system error",
	}
	ErrUnauthorized = &Error{
		Code:    -401,
		Message: "鉴权失败",
	}
	ErrAuthTimeout = &Error{
		Code:    -402,
		Message: "token过期",
	}
	ErrAccount = &Error{
		Code:    -403,
		Message: "账号异常",
	}
	ErrPwd = &Error{
		Code:    -404,
		Message: "账号密码错误",
	}
	ErrUnable = &Error{
		Code:    -405,
		Message: "你的账号已被禁用, 请联系管理员激活",
	}
	ErrExist = &Error{
		Code:    -406,
		Message: "该名称已存在",
	}
	ErrUserUpdate = &Error{
		Code:    -407,
		Message: "更新失败",
	}
	ErrParam = &Error{
		Code:    10001,
		Message: "请求参数错误",
	}
	ErrDB = &Error{
		Code:    10002,
		Message: "DB error",
	}
	ErrRedis = &Error{
		Code:    10003,
		Message: "Redis error",
	}
)

type Error struct {
	Code    int    // 错误码
	Message string // 错误信息
	Err     error  // 详细错误信息
}

func (e *Error) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e *Error) String() string {
	return e.Error()
}

/* 基于error包裹
 * err : 错误信息
 * wrap: 包裹信息
 */
func (e *Error) Wrap(err error, message string) *Error {
	ne := &Error{}
	ne.Code = e.Code
	ne.Message = e.Message
	if e.Err == nil {
		ne.Err = err
	} else {
		ne.Err = fmt.Errorf("%v , %v", e.Err, err)
	}
	ne.Err = errors.Wrap(ne.Err, message)
	return ne
}
