// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "hk591_go/app/http/controllers/api/v1"
	"hk591_go/app/models/user"
	"hk591_go/app/requests"
	"hk591_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.Data(c, 200, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}
