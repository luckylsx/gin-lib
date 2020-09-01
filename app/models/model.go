package models

import (
	"errors"
	"fmt"
	"gin-lib/app/pkg/system"
	"gin-lib/app/utils/logging"
	"math"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gin-lib/conf"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Connect connect database and get object of connection
func Connect(vo ConVo) (*gorm.DB, error) {
	var err error
	configMap, errs := conf.GetDbConfig(vo.Driver)
	if errs != nil {
		logging.Error("database config get failed: " + errs.Error())
		return nil, errs
	}
	config, ok := ConnectReadOrWrite(configMap, vo)
	if !ok {
		configNotExist := vo.SQLType + " database config get failed"
		logging.Info(configNotExist)
		return nil, errors.New(configNotExist)
	}
	mysqlObj := config["user"] + ":" + config["password"] + "@(" + config["host"] + ":" + config["port"] + ")/" + config["dbname"] + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", mysqlObj)

	if err != nil {
		logging.Error("database connect failed :" + err.Error())
		return nil, err
	}
	// 使用表单数形式
	db.SingularTable(true)

	// 自定义回调 可处理 创建 更新 和删除等自定义处理相关操作
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// 使用连接池
	// 设置数据库最大打开连接数
	db.DB().SetMaxOpenConns(system.LsCPU() + 1)
	// 设置空闲连接池中最大连接数
	db.DB().SetMaxIdleConns(int(math.Ceil(float64(system.LsCPU()) / 2)))

	// 启用Logger，显示详细日志
	db.LogMode(true)
	// db.SetLogger(gorm.LogFormatter())
	db.SetLogger(logging.SQLLog())

	return db, nil
}

// ConnectReadOrWrite 连接数据库
func ConnectReadOrWrite(config conf.MyMap, vo ConVo) (map[string]string, bool) {
	if vo.SQLType == "" {
		vo.SQLType = "read"
	}
	if config, ok := config[vo.SQLType]; ok {
		return config, true
	}
	return nil, false
}

// Close close the database
func Close() {
	if db != nil {
		_ = db.Close()
	}
}

// updateTimeStampForCreateCallback will set `UpdatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", time.Now())
	}
}

// deleteCallback will set `DeletedOn` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("IsDeleted")
		deletedOnUpdated, _ := scope.FieldByName("UpdatedAt")
		if !scope.Search.Unscoped && hasDeletedOnField {
			sql := fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(IsDeleted),
				scope.Quote(deletedOnUpdated.DBName),
				scope.AddToVars(time.Now()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)
			// fmt.Println(sql) 调试
			scope.Raw(sql).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
