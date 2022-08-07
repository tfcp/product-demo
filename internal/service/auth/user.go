package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"tfpro/internal/model/auth"
	"tfpro/library/log"
	"tfpro/library/utils"
)

type UserService struct {
	userModel *auth.User
}

func NewUserService() (s *UserService) {
	h := &UserService{}
	h.userModel = &auth.User{}
	return h
}

func (this *UserService) List(where map[string]interface{}, page, size int) ([]*auth.User, error) {
	list, err := this.userModel.ListUser(where, page, size)
	if err != nil {
		log.Logger.Errorf("UserService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *UserService) Count(where map[string]interface{}) (int, error) {
	count, err := this.userModel.CountUser(where)
	if err != nil {
		log.Logger.Errorf("UserService CountError: %v", err)
		return 0, err
	}
	return count, nil
}

func (this *UserService) One(where map[string]interface{}) (*auth.User, error) {
	one, err := this.userModel.OneUser(where)
	if err != nil {
		log.Logger.Errorf("UserService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}

func (this *UserService) Info(token string) (map[string]interface{}, error) {
	info, err := utils.ParseToken(token)
	if err != nil {
		log.Logger.Errorf("UserService InfoError: %v", err)
		return nil, err
	}
	resInfo := map[string]interface{}{
		"avatar":       info.Avatar,
		"introduction": info.Introduction,
		"name":         info.Name,
		"roles":        info.Roles,
	}
	return resInfo, nil
}

func (this *UserService) Delete(id int) error {
	whereDelete := map[string]interface{}{
		"id": id,
	}
	err := this.userModel.DeleteUser(whereDelete)
	if err != nil {
		log.Logger.Errorf("UserService DeleteError: %v", err)
		return err
	}
	return nil
}

func (this *UserService) ChangeStatus(id, status int) error {
	whereOne := map[string]interface{}{
		"id": id,
	}
	user, err := this.userModel.OneUser(whereOne)
	if err != nil {
		log.Logger.Errorf("UserService ChangeStatus OneError: %v", err)
		return err
	}
	if err := this.userModel.UpdateUser(user, map[string]interface{}{"status": status}); err != nil {
		log.Logger.Errorf("UserService ChangeStatus UpdateUserError: %v", err)
		return err
	}
	return nil
}

func (this *UserService) UserExists(id int, name string) bool {
	if id > 0 {
		return false
	}
	whereUser := map[string]interface{}{
		"username": name,
	}
	user, err := this.userModel.OneUser(whereUser)
	if err != nil {
		log.Logger.Errorf("UserExists OneUserError: %v", err)
		return true
	}
	if user.ID != 0 {
		return true
	}
	return false
}

func (this *UserService) ChangePwd(id int, oldPwd, newPwd1, newPwd2 string) error {
	var err error
	whereOld := map[string]interface{}{
		"id": id,
	}
	oldUser, err := this.userModel.OneUser(whereOld)
	if err != nil {
		log.Logger.Errorf("UserService ChangePwd OneUserError: %v", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(oldUser.Pwd), []byte(oldPwd)) //加密处理
	if err != nil {
		return errors.New(fmt.Sprintf("旧密码不正确: %v", err))
	}
	if newPwd1 != newPwd2 {
		return errors.New(fmt.Sprintf("两次密码输入不一样: %v", err))
	}
	upMap := map[string]interface{}{
		"password": newPwd1,
	}
	if err := this.userModel.UpdateUser(oldUser, upMap); err != nil {
		log.Logger.Errorf("UserService ChangePwd UpdateUserError: %v", err)
		return err
	}
	return err
}

func (this *UserService) Save(id, role int, name, avatar, introduction, pwd string) error {
	var err error
	newUser := auth.User{
		Name:         name,
		Role:         role,
		Pwd:          pwd,
		Avatar:       avatar,
		Introduction: introduction,
	}
	if id > 0 {
		// update
		whereUp := map[string]interface{}{
			"id": id,
		}
		userUp, err := this.userModel.OneUser(whereUp)
		if err != nil {
			log.Logger.Errorf("UserService Save OneUserError: %v", err)
			return err
		}
		upMap := map[string]interface{}{}
		if role > 0 {
			upMap["role"] = role
		}
		if name != "" {
			upMap["username"] = name
		}
		if avatar != "" {
			upMap["avatar"] = avatar
		}
		if pwd != "" {
			upMap["password"] = pwd
		}
		if introduction != "" {
			upMap["introduction"] = introduction
		}
		if err = this.userModel.UpdateUser(userUp, upMap); err != nil {
			log.Logger.Errorf("UserService Save UpdateUserError: %v", err)
			return err
		}
		return nil
	}
	if err = this.userModel.CreateUser(newUser); err != nil {
		log.Logger.Errorf("UserService Save CreateUserError: %v", err)
		return nil
	}
	return nil
}
