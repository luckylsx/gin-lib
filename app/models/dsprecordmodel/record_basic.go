package dsprecordmodel

import (
	"gin-lib/conf/setting"
	"time"
)

// GetTableName return the table name
func GetTableName() string {
	env := setting.GetEnv()
	var suffix string
	if env == "pro" {
		suffix = time.Now().Format("2006")
	} else {
		suffix = "basic"
	}
	return "record_" + suffix
}
