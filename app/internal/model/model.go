package model

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func init() {
	var err error
	Db, err = gorm.Open(g.Config().GetString("database.demo.type"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		g.Config().GetString("database.demo.user"),
		g.Config().GetString("database.demo.pass"),
		g.Config().GetString("database.demo.host"),
		g.Config().GetString("database.demo.name")))
	Db.LogMode(g.Config().GetBool("database.demo.log"))
	if err != nil {
		fmt.Printf("models.Setup err: %v", err)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return g.Config().GetString("database.demo.prefix") + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	fmt.Print("db init success")
}
