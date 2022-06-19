package process

import (
	"tfpro/app/process/demo"
	"tfpro/library/utils"
)

var (
	processStdoutTitle    = []string{"进程名称", "负责人", "创建时间", "备注"}
	processStdoutContents = [][]string{}
)

// consumer process
func Process() {
	demoStdout := demo.ProcessDemo()
	processStdoutContents = append(processStdoutContents, demoStdout...)

	// stdout cron list
	utils.TableStdout(processStdoutTitle, processStdoutContents)
}
