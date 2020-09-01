package route

// DailyAddCreateJob DailyAddCreateJob
/**
 * define the struct and implement Run interface
 * 定义之后可执行自定义job
 * 方便处理以下情况：
 * 1、遇到错误 捕获错误 继续执行
 * 2、如果前一个脚本未执行完成 不覆盖执行
 */
type DailyAddCreateJob struct {
	dailyAddCreate func()
}

// Run implement Job interface
func (d *DailyAddCreateJob) Run() {
	d.dailyAddCreate()
}
