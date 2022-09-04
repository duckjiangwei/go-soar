// Package requests 处理请求数据和表单验证
package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Dns struct {
	Host     string `form:"host,omitempty" valid:"host"`
	Db       string `form:"db,omitempty" valid:"db"`
	Port     string `form:"port,omitempty" valid:"port"`
	Username string `form:"username,omitempty" valid:"username"`
	Password string `form:"password,omitempty" valid:"password"`
}

type SoarSqlRequest struct {
	Sql string `form:"sql,omitempty" valid:"sql"`
	Dns
}

func SoarSql(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"sql":  []string{"required"},
		"host": []string{"ip"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"sql": []string{
			"required:sql参数为必须",
		},
		"host": []string{
			"ip:请输出正确的ip格式",
		},
	}

	return validate(data, rules, messages)
}

func SoarFile(r *http.Request) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"file:sql_file": []string{"required", "ext:sql", "size:10000000"},
		"host":          []string{"ip"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"file:sql_file": []string{
			"required:请选择上传的文件",
			"ext:请选择.sql文件",
			"size:文件太大",
		},
		"host": []string{
			"ip:请输出正确的ip格式",
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
