package demo

import (
	"gf/app/internal/model"
	"github.com/jinzhu/gorm"
)

// gorm文档: https://www.tizi365.com/archives/22.html

type User struct {
	*model.Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Sex          int    `json:"sex"`
	Status       int    `json:"status"`
	Role         int    `json:"role"`
	Pwd          string `json:"pwd"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	CreateAt     string `json:"create_at"`
}

func (this *User) ListUser(where map[string]interface{}) ([]*User, error) {
	var users []*User
	model.Db.LogMode(true)
	session := model.Db
	// like search
	name, ok := where["name"]
	if ok {
		session = session.Where("name LIKE ?", name.(string)+"%")
		delete(where, "name")
	}
	err := session.Where(where).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users, nil
		}
		return users, err
	}
	return users, nil
}

func (this *User) OneUser(where map[string]interface{}) (User, error) {
	var user User
	err := model.Db.Where(where).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func (this *User) DeleteUser(where map[string]interface{}) error {
	var user User
	err := model.Db.Where(where).Delete(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// ps: 当用结构体更新的时候，当结构体的值是""或者0，false等，就什么也不会更新。
//func (this *User) UpdateUser(user User, update map[string]interface{}) error {
func (this *User) UpdateUser(user User) error {
	// 更新单个值
	//err := model.Db.Model(&user).Update(update).Error
	// 更新模型
	err := model.Db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *User) CreateUser(user User) error {
	err := model.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
