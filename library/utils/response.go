package utils

import (
	"gf/library/code"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//func Response(Ctx *gin.Context, code int, msg string, data interface{}) {
func Response(Ctx *gin.Context, ErrorMsg *code.Error, data interface{}) {
	Ctx.JSON(200, Res{
		Code:    ErrorMsg.Code,
		Message: ErrorMsg.Msg,
		Data:    data,
	})
	return
}
