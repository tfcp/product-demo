package code

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrSuccess = &Error{
		Code: 0,
		Msg:  "success",
	}
	ErrSystem = &Error{
		Code: -1,
		Msg:  "system error",
	}
	ErrUnauthorized = &Error{
		Code: -401,
		Msg:  "user Unauthorized",
	}
	ErrParam = &Error{
		Code: 10001,
		Msg:  "params error",
	}

	ErrDB = &Error{
		Code: 10002,
		Msg:  "DB error",
	}
	ErrRedis = &Error{
		Code: 10003,
		Msg:  "Redis error",
	}
)

type Error struct {
	Code int    // 错误码
	Msg  string // 错误信息
	Err  error  // 详细错误信息
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
	ne.Msg = e.Msg
	if e.Err == nil {
		ne.Err = err
	} else {
		ne.Err = fmt.Errorf("%v , %v", e.Err, err)
	}
	ne.Err = errors.Wrap(ne.Err, message)
	return ne
}
