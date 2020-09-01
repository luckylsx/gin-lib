package main

import (
	"flag"
	"gin-lib/app/route"
)

func main() {
	var isCron bool
	// 接收flag 参数 确定运行 方式 默认 运行 cron 服务
	// 指定 -isCron=true 时 运行 cron 计划任务 服务
	flag.BoolVar(&isCron, "isCron", false, "is start http server")
	flag.Parse()
	if !isCron {
		// 启动http 服务
		route.HTTPServer()
	} else {
		// 启动计划任务
		route.CronServer()
	}
}
