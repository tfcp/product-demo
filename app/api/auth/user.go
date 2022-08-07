package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"tfpro/internal/middleware/jwt"
	"tfpro/internal/service/auth"
	"tfpro/library/code"
	"tfpro/library/utils"
)

type RequestUserList struct {
	Name     string `json:"name" form:"name"`
	Role     int    `json:"role" form:"role"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// @Summary swagger文档示例2
// @Tags 示例
// @Param name query string true "name " required
// @Param sex query string true "sex " required
// @Param page query string true "page " required
// @Param page_size query string true "page_size " required
// @Success 200 {object} utils.Res
// @Router /auth/user-list [get]
func UserListApi(c *gin.Context) {
	var reqUserList RequestUserList
	c.Bind(&reqUserList)
	if err := gvalid.CheckStruct(c, reqUserList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := auth.NewUserService()
	whereCondition := map[string]interface{}{}
	if reqUserList.Name != "" {
		whereCondition["username"] = reqUserList.Name
	}
	if reqUserList.Role != 0 {
		whereCondition["role"] = reqUserList.Role
	}
	ListInfo, err := hs.List(whereCondition, reqUserList.Page, reqUserList.PageSize)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserCount struct {
	Name string `json:"name" form:"name"`
	Role int    `json:"role" form:"role"`
}

func UserCountApi(c *gin.Context) {
	var reqUserCount RequestUserCount
	c.Bind(&reqUserCount)
	if err := gvalid.CheckStruct(c, reqUserCount, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := auth.NewUserService()
	whereCondition := map[string]interface{}{}
	if reqUserCount.Name != "" {
		whereCondition["username"] = reqUserCount.Name
	}
	if reqUserCount.Role != 0 {
		whereCondition["role"] = reqUserCount.Role
	}
	ListInfo, err := hs.Count(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestUserInfo struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func UserInfoApi(c *gin.Context) {
	var reqUserInfo RequestUserInfo
	c.Bind(&reqUserInfo)
	if err := gvalid.CheckStruct(c, reqUserInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := auth.NewUserService()
	whereCondition := map[string]interface{}{}
	whereCondition["id"] = reqUserInfo.Id
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
	hs := auth.NewUserService()
	err := hs.Delete(reqUserDelete.Id)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserChange struct {
	Id     int `json:"id" form:"id" valid:"id      @required|integer|min:1#id不能为空"`
	Status int `json:"status" form:"status" valid:"status      @required|integer|min:1#status不能为空"`
}

func UserChangeStatusApi(c *gin.Context) {
	var reqUserChange RequestUserChange
	c.Bind(&reqUserChange)
	if err := gvalid.CheckStruct(c, reqUserChange, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := auth.NewUserService()
	err := hs.ChangeStatus(reqUserChange.Id, reqUserChange.Status)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserSave struct {
	Id           int    `json:"id" form:"id"`
	Role         int    `json:"role" form:"role" valid:"role      @required|integer|min:1#role不能为空"`
	Name         string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Pwd          string `json:"pwd" form:"pwd" valid:"pwd      @required#pwd不能为空"`
	Avatar       string `json:"avatar" form:"avatar"`
	Introduction string `json:"introduction" form:"introduction"`
}

func UserSaveApi(c *gin.Context) {
	var reqUserSave RequestUserSave
	c.Bind(&reqUserSave)
	if err := gvalid.CheckStruct(c, reqUserSave, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	us := auth.NewUserService()
	if us.UserExists(reqUserSave.Id, reqUserSave.Name) {
		utils.Response(c, code.ErrExist, nil)
		return
	}
	err := us.Save(reqUserSave.Id,
		reqUserSave.Role,
		reqUserSave.Name,
		reqUserSave.Avatar,
		reqUserSave.Introduction,
		reqUserSave.Pwd,
	)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestUserPwdChange struct {
	OldPwd  string `json:"old_pwd" form:"old_pwd" valid:"old_pwd      @required#old_pwd不能为空"`
	NewPwd1 string `json:"new_pwd_1" form:"new_pwd_1" valid:"new_pwd_1      @required#new_pwd_1不能为空"`
	NewPwd2 string `json:"new_pwd_2" form:"new_pwd_2" valid:"new_pwd_2      @required#new_pwd_2不能为空"`
}

func UserPwdChangeApi(c *gin.Context) {
	var reqUserPwdChange RequestUserPwdChange
	c.Bind(&reqUserPwdChange)
	if err := gvalid.CheckStruct(c, reqUserPwdChange, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	us := auth.NewUserService()
	authInfo := jwt.GetUserInfo(c)
	if err := us.ChangePwd(
		authInfo.Id,
		reqUserPwdChange.OldPwd,
		reqUserPwdChange.NewPwd1,
		reqUserPwdChange.NewPwd2,
	); err != nil {
		utils.Response(c, code.ErrUserUpdate, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, nil)
}
