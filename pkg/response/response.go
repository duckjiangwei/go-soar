// Package response 响应处理工具
package response

import (
	"hk591_go/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用
func Success(c *gin.Context) {
	JSON(c, gin.H{
		"code":    200,
		"message": "操作成功！",
	})
}

//Error 响应 200 和预设 操作失败 的 Json 数据
//执行某个『没有具体返回数据』的『变更』操作成功后调用
func Error(c *gin.Context, code int) {
	JSON(c, gin.H{
		"code":    code,
		"message": "操作失败",
	})
}

// http 200 和带 data 键的 JSON 数据
func Data(c *gin.Context, code int, data interface{}) {
	JSON(c, gin.H{
		"code": code,
		"data": data,
	})
}

// http 响应 200，code 400
func BadRequest(c *gin.Context, err error, msg string) {
	logger.LogIf(err)
	JSON(c, gin.H{
		"code":    http.StatusBadRequest,
		"message": msg,
		"error":   err.Error(),
	})
}

//响应表单验证错误
func ValidationError(c *gin.Context, errors map[string][]string) {
	JSON(c, gin.H{
		"code":  http.StatusUnprocessableEntity,
		"error": errors,
	})
}

func Abort404(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": msg,
	})
}
