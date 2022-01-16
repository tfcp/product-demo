package demo

import (
	"gf/app/internal/model"
	"time"
)

type User struct {
	*model.Model
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Sex          int       `json:"sex"`
	Role         string    `json:"role"`
	Pwd          string    `json:"pwd"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	CreateAt     time.Time `json:"create_at"`
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
		return users, err
	}
	return users, nil
}

func (this *User) OneUser(where map[string]interface{}) (User, error) {
	var user User
	err := model.Db.Where(where).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
