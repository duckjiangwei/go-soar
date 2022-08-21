// Package routes 注册路由
package routes

import (
	controllers "go_web/app/http/controllers/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1 的路由组
	v1 := r.Group("/v1")
	{
		/* ------------------ 用户 ------------------ */
		soarGroup := v1.Group("/soar")
		s := new(controllers.SoarController)
		soarGroup.GET("", s.Index)
	}
}
