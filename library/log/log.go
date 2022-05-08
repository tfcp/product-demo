package log

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"tfpro/library/config"
)

var Logger *glog.Logger

func Setup() {
	Logger = glog.New()
	path, _ := config.TfConf.Value("log.path").String()
	level, _ := config.TfConf.Value("log.level").String()
	stdout, _ := config.TfConf.Value("log.stdout").String()
	if err := Logger.SetConfigWithMap(g.Map{
		"path":   path,
		"level":  level,
		"stdout": stdout,
	}); err != nil {
		Logger.Fatal(err)
	}
}
