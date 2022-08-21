// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type DealRecordRequest struct {
	Id      int `form:"id,omitempty"`
	PerPage int `form:"page_size,omitempty" valid:"per_page"`
}

func DealRecord(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"id":       []string{"min:1"},
		"per_page": []string{"numeric_between:2,100"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"id": []string{
			"min:ID异常",
		},
		"per_page": []string{
			"numeric_between:每页数量在2~100",
		},
	}

	return validate(data, rules, messages)
}
