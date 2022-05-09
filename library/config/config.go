package config

import (
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"tfpro/library/environment"
)

var (
	TfConf             config.Config
	TfConfDriverApollo = "apollo"
	TfConfDriverNacos  = "nacos"
	TfConfDriverConsul = "consul"
)

func Setup(args ...interface{}) error {
	// path为config目录, 走yaml配置文件, 比较好适配apollo等配置中心
	path := "config"
	if len(args) == 0 {
		TfConf = config.New(config.WithSource(file.NewSource(path)))
	} else {
		switch args[0] {
		case "":
			TfConf = config.New(config.WithSource(file.NewSource(path)))
			break
		case TfConfDriverApollo:
			TfConf = config.New(
				config.WithSource(
					apollo.NewSource(
						apollo.WithAppID(environment.AppID),
						apollo.WithCluster("dev"),
						//apollo.WithEndpoint("http://apollo-server:8080"),
						apollo.WithEndpoint(environment.ApolloAddr),
						apollo.WithNamespace("config.yaml"),
						apollo.WithEnableBackup(),
						apollo.WithSecret("ad75b33c77ae4b9c9626d969c44f41ee"),
					),
				),
			)
			break
		case TfConfDriverNacos:
			// TODO
			break
		case TfConfDriverConsul:
			// TODO
			break
		}
	}
	return TfConf.Load()
}
