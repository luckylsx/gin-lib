package dspbusinesspooldao

import (
	"gin-lib/app/models"
	"gin-lib/app/pkg/system"
	"gin-lib/app/utils/logging"
	"time"
)

// Params Params
type Params struct {
	TableName        string
	DbSelect         string
	WhereStatusKey   string
	WhereStatusValue int8
}

// DspBusinessNewlyCount dsp-business-pool 新增数据统计封装 dspBusinessNewlyCount
func DspBusinessNewlyCount(params *Params) int {
	db, err := models.Connect(models.ConVo{Driver: "dsp-business-pool"})
	if err != nil {
		logging.Info("db connect failed : error ", err.Error(), " method:", system.RunFuncName())
		return 0
	}
	var count int
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	start, end := yesterday+" 00:00:00", yesterday+" 23:59:59"

	db = db.Table(params.TableName)
	if params.DbSelect != "" {
		db.Select(params.DbSelect)
	}
	if params.WhereStatusKey != "" && params.WhereStatusValue != 0 {
		db.Where(params.WhereStatusKey, params.WhereStatusValue)
	}
	db.Where("created_at >= ?", start).
		Where("created_at <= ?", end).
		Where("is_deleted = ?", 0).
		Count(&count)
	defer models.Close()
	return count
}
