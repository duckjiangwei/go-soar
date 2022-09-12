package v1

import (
	"go_web/app/requests"
	"go_web/app/services"
	"go_web/pkg/config"
	"go_web/pkg/response"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type SoarController struct {
	BaseAPIController
}

//获取启发规则
func (ctrl *SoarController) RuleList(c *gin.Context) {
	ruleName := "rule.md"
	err := services.GetRuleList(ruleName)
	if err != nil {
		response.BadRequest(c, err, "")
		return
	}
	//ajax 返回
	url := cast.ToString(config.Env("APP_URL")) + "/soar-result/" + ruleName
	data := map[string]string{
		"url": url,
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}

func (ctrl *SoarController) Batch(c *gin.Context) {
	//获取参数
	request := requests.SoarSqlRequest{}
	c.Bind(&request)
	//参数验证
	errs := requests.SoarFile(c.Request)
	if len(errs) > 0 {
		response.Data(c, 200, errs)
		return
	}
	rootPath, _ := os.Getwd()
	//处理文件上传
	sqlPath, errLoad := services.UploadFile(c, rootPath)
	if errLoad != nil {
		response.BadRequest(c, errLoad, "")
		return
	}
	//开始执行分析并保存结果
	randFileName := cast.ToString(time.Now().Unix()) + ".html"
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/" + randFileName

	errAnalyze := services.AnalyzeAndSave(request, rootPath, sqlPath, soarResultPath)
	if errAnalyze != nil {
		response.BadRequest(c, errAnalyze, "")
		return
	}

	//ajax 返回
	url := cast.ToString(config.Env("APP_URL")) + "/soar-result/" + randFileName
	data := map[string]string{
		"url": url,
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}

func (ctrl *SoarController) Single(c *gin.Context) {
	//参数验证
	request := requests.SoarSqlRequest{}
	if ok := requests.Validate(c, &request, requests.SoarSql); !ok {
		return
	}
	rootPath, _ := os.Getwd()
	//生成 sql 文件
	sql := request.Sql
	sqlPath := rootPath + "/" + cast.ToString(config.Env("SOAR_SQL")) + "/" + "sql_" + cast.ToString(time.Now().Unix()) + ".sql"
	os.WriteFile(sqlPath, []byte(string(sql)), 0644)
	//生成结果文件
	randFileName := cast.ToString(time.Now().Unix()) + ".html"
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/" + randFileName
	//执行分析并保存结果
	err := services.AnalyzeAndSave(request, rootPath, sqlPath, soarResultPath)
	if err != nil {
		response.BadRequest(c, err, "")
		return
	}

	//ajax 返回
	url := cast.ToString(config.Env("APP_URL")) + "/soar-result/" + randFileName
	data := map[string]string{
		"url": url,
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}
