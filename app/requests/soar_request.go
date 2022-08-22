// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SoarRequest struct {
	Sql string `form:"sql,omitempty" valid:"sql"`
}

func SoarSql(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"sql": []string{"required", "size:5000"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"sql": []string{
			"required:sql参数为必须",
			"size:sql太长了",
		},
	}

	return validate(data, rules, messages)
}
