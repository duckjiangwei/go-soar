// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type EstateIndexRequest struct {
	AreaId     int    `form:"area_id,omitempty"`
	DistrictId string `form:"district_id,omitempty"`
	Keyword    string `form:"keyword,omitempty"`
	PerPage    int    `form:"page_size,omitempty" valid:"per_page"`
	Price      int    `form:"price,omitempty" valid:"price"`
	Age        int    `form:"age,omitempty"`
}

func EstateIndex(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"per_page": []string{"numeric_between:2,100"},
		"price":    []string{"numeric_between:0,6"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"per_page": []string{
			"numeric_between:每页数量在2~100",
		},
		"price": []string{
			"numeric_between:价格区间参数错误",
		},
	}

	return validate(data, rules, messages)
}
