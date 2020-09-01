package main

import (
	"gin-lib/app/service"
	"reflect"
	"testing"
)

func BenchmarkGetNewlyAddData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.GetNewlyAddData()
	}
	b.StopTimer()
}

// 统计数据单元测试
func TestNewlyAddData(t *testing.T) {
	maps := map[string]int{
		"StaffNewlyAdd":           0,
		"CompanyNewlyAdd":         0,
		"StartEvaluations":        0,
		"StartEvaluationStaff":    0,
		"FinishEvaluation":        0,
		"FinishEvaluationStaff":   0,
		"DownloadEvaluationCount": 0,
		"PreviewEvaluationCount":  0,
		//"CreateDate":              time.Now().Format("2006-01-02"),
	}
	data := service.GetNewlyAddData()
	for index, v := range maps {
		actual := reflect.ValueOf(&data).Elem().FieldByName(index)
		if int(actual.Int()) != v {
			t.Errorf("the statistic data has error , expected data is :%v, but got data is %v", v, actual)
		}
	}
}
