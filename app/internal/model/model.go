package model

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

type Model struct {
	ID       int    `gorm:"primary_key" json:"id"`
	CreateAt string `json:"create_at"`
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
	Db.DB().SetConnMaxLifetime(time.Hour)
	//Db.Callback().Create().After("gorm:create_at", updateTimeStampForCreateCallback)
	//Db.Callback().Create().After("gorm:create_at", updateTimeStampForCreateCallback)
	//Db.Callback().Update().After("updateTimeStampForCreateCallback")

	fmt.Print("db init success")
}

// 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	//if !scope.HasError() {
	nowTime := time.Now().Format("2006-01-01 03:04:00")
	fmt.Println("s:", nowTime)
	if createTimeField, ok := scope.FieldByName("CreateAt"); ok {
		fmt.Println("a", nowTime)
		//if createTimeField.IsBlank {
		fmt.Println(3344)
		createTimeField.Set(nowTime)
		//}
	}
	//}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_at"); !ok {
		scope.SetColumn("UpdatedTime", time.Now().Unix())
	}
}
