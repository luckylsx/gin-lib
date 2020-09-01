package dsprecordmodel

import "gin-lib/app/models"

// DateNewlyAdd date_newly_add 每日新增数据model
type DateNewlyAdd struct {
	models.Model
	StaffNewlyAdd           int
	CompanyNewlyAdd         int
	StartEvaluations        int
	StartEvaluationStaff    int
	FinishEvaluation        int
	FinishEvaluationStaff   int
	DownloadEvaluationCount int
	PreviewEvaluationCount  int
	CreateDate              string
	// CreatedAt               string
}
