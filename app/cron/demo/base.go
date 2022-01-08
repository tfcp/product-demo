package demo

import (
	"gf/library/log"
	"github.com/gogf/gf/os/gcron"
)

// (Seconds Minutes Hours Day Month Week)
// @hourly @daily @weekly @monthly @yearly
// @every <duration>: @every 1h30m10s 每隔1小时30分钟10秒
// 2 * * * * * 		: 每分钟第2秒执行
// */5 * * * * *	: 每分钟第2秒执行
// 0 * * * * *		: 每分钟执行

func CronDemo() {
	_, err := gcron.Add("*/5 * * * * *", func() {
		HelloCron()
	}, "DemoHelloCron")
	if err != nil {
		log.Logger.Errorf("DemoHelloCron Start Error: %v", err)
	}
}
