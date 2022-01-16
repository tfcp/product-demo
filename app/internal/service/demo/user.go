package demo

import (
	"gf/app/internal/model/demo"
	"gf/library/log"
	"gf/library/utils"
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
