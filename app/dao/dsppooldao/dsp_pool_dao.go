package dsppooldao

import (
	"gin-lib/app/models"
	"gin-lib/app/pkg/system"
	"gin-lib/app/utils/logging"
	"time"
)

// GetCompanyApplyNewly return the count of company newly add
func GetCompanyApplyNewly(tableName string) int {
	db, err := models.Connect(models.ConVo{Driver: "dsp-pool"})
	if err != nil {
		logging.Error("db connect failed: ", err.Error(), "method : ", system.RunFuncName())
		return 0
	}
	var count int
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	start, end := yesterday+" 00:00:00", yesterday+" 23:59:59"
	db.Table(tableName).Where("created_at >= ?", start).
		Where("created_at <= ?", end).
		Where("parent_id = ?", 0).
		Where("is_deleted = ?", 0).
		Count(&count)
	defer db.Close()
	return count
}
