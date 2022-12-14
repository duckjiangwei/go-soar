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
		soarGroup := v1.Group("/soar")
		s := new(controllers.SoarController)
		//获取启发规则
		soarGroup.GET("/rule", s.RuleList)
		//单条sql优化
		soarGroup.POST("/single", s.Single)
		//批量优化
		soarGroup.POST("/batch", s.Batch)
	}
}
