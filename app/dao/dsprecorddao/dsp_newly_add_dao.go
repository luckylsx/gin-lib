package dsprecorddao

import (
	"gin-lib/app/models"
	"gin-lib/app/models/dsprecordmodel"
	"log"
	"time"
)

// Params Params
type Params struct {
	TableName          string
	DbSelect           string
	WhereActionIDKey   string
	WhereActionIDValue int8
}

// InsertOne 将每日统计数据入库
func InsertOne(data dsprecordmodel.DateNewlyAdd) error {
	db, err := models.Connect(models.ConVo{Driver: "dsp-pool-record", SQLType: "write"})
	if err != nil {
		log.Println("staff number get failed ", err.Error())
		return err
	}
	db.Create(&data)
	return nil
}

// RecordCount record_basic 记录表统计
func RecordCount(params *Params) int {
	db, err := models.Connect(models.ConVo{Driver: "dsp-pool-record"})
	if err != nil {
		log.Println("staff number get failed ", err.Error())
		return 0
	}
	var count int
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	start, end := yesterday+" 00:00:00", yesterday+" 23:59:59"

	db = db.Table(params.TableName)
	if params.DbSelect != "" {
		db.Select(params.DbSelect)
	}
	if params.WhereActionIDKey != "" && params.WhereActionIDValue != 0 {
		db.Where(params.WhereActionIDKey, params.WhereActionIDValue)
	}
	db.Where("log_time >= ?", start).
		Where("log_time <= ?", end).
		Count(&count)
	defer db.Close()
	return count
}

// DeletedRecode DeletedRecode
func DeletedRecode(id int) int {
	var dateNewlyAdd dsprecordmodel.DateNewlyAdd
	db, err := models.Connect(models.ConVo{Driver: "dsp-pool-record"})
	if err != nil {
		log.Println("staff number get failed ", err.Error())
		return 0
	}

	dateNewlyAdd.ID = id
	db.Where("is_deleted = ?", models.NotDeleted).Delete(&dateNewlyAdd)
	return int(db.RowsAffected)
}
