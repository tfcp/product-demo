package tools

import (
	"fmt"
	"gf/tools/ormgen"
)

var (
	// 生成model的路径(可以自定义 但是需要确保路径文件存在)
	path = "./model/model.go"
	// mysql地址
	//dsn = "user:pwd@tcp(localhost:3306)/database?charset=utf8"
	dsn = "med_read:WFgpz8#QxRQg@tcp(118.31.236.23:3306)/med_msg?charset=utf8"
	//dsn = "meddev:akdjfkaj2399I@tcp(114.55.3.97:3307)/med_spread?charset=utf8"
)

// quick generate orm model
func OrmGenTool() {
	t2t := ormgen.NewTable2Struct()
	// table 需要的table名称()
	err := t2t.
		// table 需要转换的table名称(不指定则默认全部表转换)
		Table("chat_group").
		//Table("doctorUser").
		// 包名
		PackageName("model").
		SavePath(path).
		// 是否需要json标签
		EnableJsonTag(true).
		// 是否生成对应的method方法(goFrame)
		MakeGfMethod(true).
		Dsn(dsn).
		Run()
	if err != nil {
		fmt.Println(err)
	}
}
