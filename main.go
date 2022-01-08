package main

import (
	"gf/bootstrap"
	"gf/library/log"
	"github.com/gogf/gf/os/gcmd"
)

func main() {
	if err := gcmd.BindHandleMap(map[string]func(){
		"server": bootstrap.Run,
		//"consumer": process.Listener,
	}); err != nil {
		log.Logger.Fatal(err)
	}
	if err := gcmd.AutoRun(); err != nil {
		log.Logger.Fatal(err)
	}
}
