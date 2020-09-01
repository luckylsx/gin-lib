package route

import (
	"context"
	"gin-lib/app/utils/logging"
	"gin-lib/commend"
	"gin-lib/conf/setting"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/gin-gonic/gin"
)

// StartAndGraceShutdown 开启http服务 关闭时优雅的关闭服务
func StartAndGraceShutdown(r *gin.Engine) {
	srv := &http.Server{
		Addr:    setting.Conf().Server.HTTPPort,
		Handler: r,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen occurred error : %s\n", err)
		}
	}()
	log.Println("http server started")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error : %s\n", err)
	}
	log.Println("server existed")
}

// HTTPServer 定义路由运行 http 服务
func HTTPServer() {
	// start http server
	gin.DisableConsoleColor()
	// 将请求日志记录到日志文件
	currentDate := time.Now().Format("2006-01-02")
	fileName := "runtime/log/log-" + currentDate + ".log"
	f, err := os.Create(fileName)
	if err != nil {
		logging.Error("log create  failed : ", err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.Default()
	// 引入路由定义
	Router(r)
	// 启动http服务 并优雅的关闭 ctrl+c 之后 会完成当前已接收到的请求任务之后 关闭服务
	StartAndGraceShutdown(r)
}

// CronServer 运行 cron 计划任务服务
func CronServer() {
	// 定时任务
	c := cron.New(cron.WithSeconds())
	// DelayIfStillRunning
	// DelayIfStillRunning serializes jobs, delaying subsequent runs until the previous one is complete
	// 如果上个任务未执行完成 则延迟执行当前任务直到上个任务执行完成
	_, _ = c.AddJob("@daily", cron.NewChain(
		cron.Recover(cron.DefaultLogger),             // 遇到错误 捕获错误 继续执行
		cron.DelayIfStillRunning(cron.DefaultLogger), // 不覆盖执行
	).Then(&DailyAddCreateJob{dailyAddCreate: commend.DateNewlyCreate}))
	c.Start()

	// 阻塞主goroutine 防止 main 函数退出
	select {}
}
