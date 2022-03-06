package demo

import (
	"gf/library/utils"
)

// (Seconds Minutes Hours Day Month Week)
// @hourly @daily @weekly @monthly @yearly
// @every <duration>: @every 1h30m10s 每隔1小时30分钟10秒
// 2 * * * * * 		: 每分钟第2秒执行
// */5 * * * * *	: 每分钟第2秒执行
// 0 * * * * *		: 每分钟执行

func CronDemo() [][]string {
	var stdoutContents [][]string
	stdoutContents = append(stdoutContents,
		// cron demo
		utils.AddCron("*/1 * * * * *", "HelloDemoCron1", HelloDemoCron),
		utils.AddCron("*/2 * * * * *", "HelloDemoCron2", HelloDemoCron),
	)
	return stdoutContents
}
