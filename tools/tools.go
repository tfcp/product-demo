package tools

import "gf/tools/ormgen"

// quick generate orm model
func OrmGenTools(tableName, dbName string) {
	ormgen.BaseTool(tableName, dbName)
}
