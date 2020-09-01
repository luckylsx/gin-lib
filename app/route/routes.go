package route

import (
	"gin-lib/app/controller"
	"gin-lib/app/middleware/logger"

	"github.com/gin-gonic/gin"
)

// Router route define
func Router(r *gin.Engine) {
	// 定义中间件
	r.Use(logger.SetUp())

	// 定义路由
	r.GET("/", controller.Index)
	r.GET("/error", controller.Exception)

	// 路由组
	group := r.Group("data")
	{
		group.GET("/newly", controller.NewlyAddData)
		group.GET("/json", controller.JSONResp)
		group.GET("/delete", controller.DeletedRecordByID)
	}

	// html 模版
	r.LoadHTMLGlob("storage/views/**/*")
	r.GET("/user", controller.UserShow)
}
