package demo

import (
	"gf/app/internal/model/demo"
	"gf/library/log"
	"gf/library/utils"
	"time"
)

type UserService struct {
	userModel *demo.User
}

func NewUserService() (s *UserService) {
	h := &UserService{}
	h.userModel = &demo.User{}
	return h
}

func (this *UserService) List(where map[string]interface{}) ([]*demo.User, error) {
	list, err := this.userModel.ListUser(where)
	if err != nil {
		log.Logger.Errorf("UserService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *UserService) One(where map[string]interface{}) (*demo.User, error) {
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
	updateMap := map[string]interface{}{
		"status": status,
	}
	if err := this.userModel.UpdateUser(user, updateMap); err != nil {
		return err
	}
	return nil
}

func (this *UserService) Save(id, status int, name, avatar, introduction string) error {
	if id > 0 {
		// update
		return nil
	}
	newUser := demo.User{
		Name:         name,
		Age:          0,
		Sex:          0,
		Status:       status,
		Role:         0,
		Pwd:          "",
		Avatar:       avatar,
		Introduction: introduction,
	}
}
