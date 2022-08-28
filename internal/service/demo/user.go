package demo

import (
	"github.com/gogf/gf/util/gconv"
	"tfpro/internal/model/demo"
	"tfpro/library/log"
	"tfpro/library/utils"
)

type UserService struct {
	userModel *demo.User
}

func NewUserService() (s *UserService) {
	h := &UserService{}
	h.userModel = &demo.User{}
	return h
}

func (this *UserService) List(where map[string]interface{}, page, size int) ([]*demo.User, error) {
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
	if err := this.userModel.UpdateUser(user, map[string]interface{}{"status": status}); err != nil {
		log.Logger.Errorf("UserService ChangeStatus UpdateUserError: %v", err)
		return err
	}
	return nil
}

func (this *UserService) Save(userInfo map[string]interface{}) error {
	var (
		err error
	)
	newUser := demo.User{}
	if err := gconv.Struct(userInfo, &newUser); err != nil {
		log.Logger.Errorf("UserService Save StructError: %v", err)
		return err
	}
	userId, ok := userInfo["id"].(int)
	if ok {
		if userId > 0 {
			resourceUp, err := this.userModel.OneUser(map[string]interface{}{
				"id": userId,
			})
			if err != nil {
				log.Logger.Errorf("UserService Save OneError: %v", err)
				return err
			}
			if err = this.userModel.UpdateUser(resourceUp, userInfo); err != nil {
				log.Logger.Errorf("UserService Save UpdateError: %v", err)
				return err
			}
			return err
		}
	}
	err = this.userModel.CreateUser(newUser)
	if err != nil {
		log.Logger.Errorf("UserService Save CreateError: %v", err)
		return err
	}
	return err
}
