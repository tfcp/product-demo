package demo

import (
	"gf/library/code"
	"gf/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type RequestHello struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
}

func Hello(c *gin.Context) {
	var reqHello RequestHello
	c.Bind(&reqHello)
	if err := gvalid.CheckStruct(c, reqHello, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	res := map[string]interface{}{
		"result": reqHello.Name,
	}
	utils.Response(c, code.ErrSuccess, res)
}
