package main

import (
	"gf/bootstrap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
)

func main() {
	if err := gcmd.BindHandleMap(map[string]func(){
		"server": bootstrap.Run,
		//"consumer": process.Listener,
	}); err != nil {
		g.Log().Fatal(err)
	}
	if err := gcmd.AutoRun(); err != nil {
		g.Log().Fatal(err)
	}
}
