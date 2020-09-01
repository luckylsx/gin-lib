package service

import (
	"encoding/json"
	"gin-lib/app/dao/dspbusinesspooldao"
	"gin-lib/app/dao/dsppooldao"
	"gin-lib/app/dao/dsprecorddao"
	"gin-lib/app/models/dspbusinesspoolmodel"
	"gin-lib/app/models/dsppoolmodel"
	"gin-lib/app/models/dsprecordmodel"
	"gin-lib/app/pkg/system"
	"gin-lib/app/utils/logging"
	resp "gin-lib/app/utils/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetNewlyAddData GetNewlyAddData
func GetNewlyAddData() dsprecordmodel.DateNewlyAdd {
	return dsprecordmodel.DateNewlyAdd{
		StaffNewlyAdd:           StaffNumbersNewly(),
		CompanyNewlyAdd:         dsppooldao.GetCompanyApplyNewly(dsppoolmodel.DspOrganizeTable),
		StartEvaluations:        StartEvaluationsNewly(),
		StartEvaluationStaff:    StartEvaluationsStaffNewly(),
		FinishEvaluation:        FinishEvaluationsNewly(),
		FinishEvaluationStaff:   FinishEvaluationStaffNewly(),
		DownloadEvaluationCount: DownloadEvaluationReportCount(),
		PreviewEvaluationCount:  PreviewEvaluationReportCount(),
		CreateDate:              time.Now().Format("2006-01-02"),
		//CreatedAt:               time.Now().Format("2006-01-02 15:04:05"),
	}
}

// StaffNumbersNewly return the count of staff number newly add
func StaffNumbersNewly() int {
	return dspbusinesspooldao.DspBusinessNewlyCount(&dspbusinesspooldao.Params{
		TableName: dspbusinesspoolmodel.StaffResumeRelationTable,
	})
}

// StartEvaluationsNewly StartEvaluationsNewly
func StartEvaluationsNewly() int {
	return dspbusinesspooldao.DspBusinessNewlyCount(&dspbusinesspooldao.Params{
		TableName: dspbusinesspoolmodel.DspEvaluationStaffRelationTable,
	})
}

// StartEvaluationsStaffNewly StartEvaluationsStaffNewly
func StartEvaluationsStaffNewly() int {
	return dspbusinesspooldao.DspBusinessNewlyCount(
		&dspbusinesspooldao.Params{
			TableName: dspbusinesspoolmodel.DspEvaluationStaffRelationTable,
			DbSelect:  "count(distinct(staff_id))",
		})
}

// FinishEvaluationsNewly FinishEvaluationsNewly
func FinishEvaluationsNewly() int {
	return dspbusinesspooldao.DspBusinessNewlyCount(
		&dspbusinesspooldao.Params{
			TableName:        dspbusinesspoolmodel.DspEvaluationStaffRelationTable,
			WhereStatusKey:   "evaluation_status = ?",
			WhereStatusValue: 2,
		})
}

// FinishEvaluationStaffNewly FinishEvaluationStaffNewly
func FinishEvaluationStaffNewly() int {
	return dspbusinesspooldao.DspBusinessNewlyCount(
		&dspbusinesspooldao.Params{
			TableName:        dspbusinesspoolmodel.DspEvaluationStaffRelationTable,
			DbSelect:         "count(distinct(staff_id))",
			WhereStatusKey:   "evaluation_status = ?",
			WhereStatusValue: 2,
		})
}

// DownloadEvaluationReportCount DownloadEvaluationReportCount
func DownloadEvaluationReportCount() int {
	return dsprecorddao.RecordCount(&dsprecorddao.Params{
		TableName:          dsprecordmodel.GetTableName(),
		DbSelect:           "count(*)",
		WhereActionIDKey:   "action_id = ?",
		WhereActionIDValue: 11,
	})
}

// PreviewEvaluationReportCount PreviewEvaluationReportCount
func PreviewEvaluationReportCount() int {
	return dsprecorddao.RecordCount(&dsprecorddao.Params{
		TableName:          dsprecordmodel.GetTableName(),
		DbSelect:           "count(*)",
		WhereActionIDKey:   "action_id = ?",
		WhereActionIDValue: 17,
	})
}

// DateNewlyAddCreate get newly add data and insert database
func DateNewlyAddCreate() bool {
	logging.Info("data newly add data handle starting..., method: ", system.RunFuncName())

	dateNewlyAdd := GetNewlyAddData()
	if err := dsprecorddao.InsertOne(dateNewlyAdd); err != nil {
		logging.Info("newly create failed :", err.Error(), "method: ", system.RunFuncName())
		return false
	}
	logging.Info("data newly add data handle ending... method: ", system.RunFuncName())
	return true
}

// GetData get data
func GetData(g *resp.Gin) interface{} {
	response := make(map[string]interface{})
	respData := GetNewlyAddData()
	respBytes, _ := json.Marshal(respData)
	err := json.Unmarshal(respBytes, &response)
	if err != nil {
		return g.ThrowException(http.StatusInternalServerError)
	}
	response["time"] = time.Now().Format("2006-01-02 15:04:05")
	response["clientIp"] = g.Ctx.ClientIP()
	response["gin_model"] = gin.Mode()
	response["env"] = os.Getenv("APP_ENV")
	return g.Success(response)
}

// DeletedRecordByID DeletedRecordByID
func DeletedRecordByID(g *resp.Gin) interface{} {
	id := g.Ctx.Query("id")
	if id == "" {
		return g.ThrowException(404)
	}
	intID, _ := strconv.Atoi(id)
	return g.Success(dsprecorddao.DeletedRecode(intID))
}
