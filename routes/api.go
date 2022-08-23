// Package routes 注册路由
package routes

import (
	controllers "go_web/app/http/controllers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	//静态页面
	r.LoadHTMLGlob("html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "go-soar",
		})
	})
	r.GET("/file", func(c *gin.Context) {
		c.HTML(http.StatusOK, "file.html", gin.H{
			"title": "go-soar",
		})
	})
	//v1 的路由组
	v1 := r.Group("/v1")
	{
		soarGroup := v1.Group("/soar")
		s := new(controllers.SoarController)
		soarGroup.POST("/sql", s.Sql)
		soarGroup.POST("/file", s.File)
	}
}
