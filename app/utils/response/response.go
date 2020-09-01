package response

import (
	error "gin-lib/app/exception/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin Gin
type Gin struct {
	Ctx *gin.Context
}

// Response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Response Response
func (g *Gin) Response(resp *Response) interface{} {
	g.Ctx.JSON(http.StatusOK, Response{
		Code:    resp.Code,
		Message: error.GetError(resp.Code),
		Data:    resp.Data,
	})
	return nil
}

// Success  success response
func (g *Gin) Success(data interface{}) interface{} {
	g.Ctx.JSON(http.StatusOK, Response{
		Code:    0,
		Message: error.GetError(0),
		Data:    data,
	})
	return nil
}

// func (g *Gin) ThrowException(code int) {
func (g *Gin) Error(code int) interface{} {
	g.Ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: error.GetError(code),
		Data:    map[string]interface{}{},
	})
	return nil
}

// ThrowException ThrowException
func (g *Gin) ThrowException(code int) interface{} {
	g.Ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: error.GetError(code),
		Data:    map[string]interface{}{},
	})
	return nil
}

// View show the html template view
func (g *Gin) View(template string, data interface{}) interface{} {
	g.Ctx.HTML(http.StatusOK, template, data)
	return nil
}
