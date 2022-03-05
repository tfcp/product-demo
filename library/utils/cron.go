package utils

import (
	"gf/library/log"
	"github.com/gogf/gf/os/gcron"
)

func AddCron(expression, cronName string, cronFunc func()) []string {
	_, err := gcron.Add(expression, func() {
		cronFunc()
	}, cronName)
	if err != nil {
		log.Logger.Errorf("%s CronStart Error: %v", cronName, err)
		return []string{}
	}
	return []string{cronName, expression}
}
