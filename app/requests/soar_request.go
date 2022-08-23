// Package requests 处理请求数据和表单验证
package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SoarSqlRequest struct {
	Sql string `form:"sql,omitempty" valid:"sql"`
}
type SoarFileRequest struct {
	File interface{} `form:"file" valid:"file"`
}

func SoarSql(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"sql": []string{"required"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"sql": []string{
			"required:sql参数为必须",
		},
	}

	return validate(data, rules, messages)
}

func SoarFile(r *http.Request) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"file:sql_file": []string{"required", "ext:sql", "size:10000"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"file:sql_file": []string{
			"required:请选择上传的文件",
			"ext:请选择.sql文件",
			"size:文件太大",
		},
	}

	opts := govalidator.Options{
		Request:  r,     // request object
		Rules:    rules, // rules map,
		Messages: messages,
	}

	v := govalidator.New(opts)
	e := v.Validate()
	return e
}
