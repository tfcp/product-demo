package mysqlcompare

import (
	"log"
	"tfpro/tools/mysqlcompare/cmd"
)

func Tools() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
