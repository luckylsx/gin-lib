package controller

import (
	"gin-lib/app/service"
	resp "gin-lib/app/utils/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// NewlyAddData get add newly data
func NewlyAddData(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	service.GetData(&g)
}

// JSONResp /json response
func JSONResp(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	g.Response(&resp.Response{Code: http.StatusOK, Message: "ok", Data: map[string]string{}})
}

// Exception Exception test
func Exception(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	g.ThrowException(http.StatusInternalServerError)
}

// Index Index request root domain
func Index(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	result := map[string]interface{}{
		"time":      time.Now().Format("2006-01-02 15:04:05"),
		"clientIp":  ctx.ClientIP(),
		"gin_model": gin.Mode(),
		"env":       os.Getenv("APP_ENV"),
	}
	g.Success(result)
}

// UserShow html template show
func UserShow(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	g.View("user/index.html", map[string]string{
		"title":       "admin",
		"description": "this is a paragraph description",
	})
}

// DeletedRecordByID DeletedRecordByID
func DeletedRecordByID(ctx *gin.Context) {
	g := resp.Gin{Ctx: ctx}
	service.DeletedRecordByID(&g)
}
