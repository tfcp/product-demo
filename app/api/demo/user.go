package demo

import (
	"gf/app/internal/service/demo"
	"gf/library/code"
	"gf/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

type RequestUserInfo struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
}

func UserInfoApi(c *gin.Context) {
	var reqHelloInfo RequestHelloInfo
	c.Bind(&reqHelloInfo)
	if err := gvalid.CheckStruct(c, reqHelloInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewHelloService()
	whereCondition := map[string]interface{}{
		"name": reqHelloInfo.Name,
	}
	oneInfo, err := hs.One(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": oneInfo,
	}

	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserList struct {
	Name string `json:"name" form:"name"`
	Sex  int    `json:"sex" form:"sex"`
}

func UserListApi(c *gin.Context) {
	var reqUserList RequestUserList
	c.Bind(&reqUserList)
	if err := gvalid.CheckStruct(c, reqUserList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	whereCondition := map[string]interface{}{}
	if reqUserList.Name != "" {
		whereCondition["name"] = reqUserList.Name
	}
	if reqUserList.Sex != 0 {
		whereCondition["sex"] = reqUserList.Sex
	}
	ListInfo, err := hs.List(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}
