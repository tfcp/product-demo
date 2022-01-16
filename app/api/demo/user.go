package demo

import (
	"gf/app/internal/service/demo"
	"gf/library/code"
	"gf/library/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
)

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

type RequestUserDetail struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func UserDetailApi(c *gin.Context) {
	var reqUserDetail RequestUserDetail
	c.Bind(&reqUserDetail)
	if err := gvalid.CheckStruct(c, reqUserDetail, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	whereCondition := map[string]interface{}{}
	whereCondition["id"] = reqUserDetail.Id
	userDetail, err := hs.One(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": userDetail,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserDelete struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func UserDeleteApi(c *gin.Context) {
	var reqUserDelete RequestUserDelete
	c.Bind(&reqUserDelete)
	if err := gvalid.CheckStruct(c, reqUserDelete, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := demo.NewUserService()
	err := hs.Delete(reqUserDelete.Id)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserChange struct {
	Id     int `json:"id" form:"id" valid:"id      @required#id不能为空"`
	Status int `json:"status" form:"status" valid:"status      @required#status不能为空"`
}

func UserChangeStatusApi(c *gin.Context) {
	var reqUserUserChange RequestUserChange
	c.Bind(&reqUserUserChange)
	if err := gvalid.CheckStruct(c, reqUserUserChange, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	//hs := demo.NewUserService()
	////err := hs.Delete(reqUserDelete.Id)
	////if err != nil {
	////	utils.Response(c, code.ErrSystem, err.Error())
	////	return
	////}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}
