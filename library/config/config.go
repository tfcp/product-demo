package config

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

var v struct {
	Service struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"service"`
}

func Setup() {
	path := "config/config.toml"
	c := config.New(config.WithSource(file.NewSource(path), ))
	fmt.Println(c.Load())
}
