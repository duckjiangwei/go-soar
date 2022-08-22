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

func (ctrl *SoarController) File(c *gin.Context) {
	rootPath, _ := os.Getwd()
	//处理文件上传
	sqlPath, errLoad := services.UploadFile(c, rootPath)
	if errLoad != nil {
		response.BadRequest(c, errLoad, "")
	}

	//开始执行分析并保存结果
	randFileName := cast.ToString(time.Now().Unix()) + ".html"
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/" + randFileName

	errAnalyze := services.AnalyzeAndSave(rootPath, sqlPath, soarResultPath)
	if errAnalyze != nil {
		response.BadRequest(c, errAnalyze, "")
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

func (ctrl *SoarController) Sql(c *gin.Context) {
	//参数验证
	request := requests.SoarRequest{}
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
	err := services.AnalyzeAndSave(rootPath, sqlPath, soarResultPath)
	if err != nil {
		response.BadRequest(c, err, "")
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
