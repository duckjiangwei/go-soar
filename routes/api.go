// Package routes 注册路由
package routes

import (
	controllers "hk591_go/app/http/controllers/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1 的路由组
	v1 := r.Group("/v1")
	{
		/* ------------------ 用户 ------------------ */
		usersGroup := v1.Group("/users")
		uc := new(controllers.UsersController)
		usersGroup.GET("", uc.Index)
	}
}
