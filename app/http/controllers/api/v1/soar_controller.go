package v1

import (
	"go_web/app/requests"
	"go_web/pkg/config"
	"go_web/pkg/response"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type SoarController struct {
	BaseAPIController
}

func (ctrl *SoarController) File(c *gin.Context) {
	rootPath, _ := os.Getwd()
	//获取待分析的 sql 文件
	sqlPath := rootPath + "/" + cast.ToString(config.Env("SOAR_SQL")) + "/test.sql"
	//开始执行分析并保存结果
	randFileName := cast.ToString(time.Now().Unix()) + ".html"
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/" + randFileName
	err := AnalyzeAndSave(rootPath, sqlPath, soarResultPath)
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
	err := AnalyzeAndSave(rootPath, sqlPath, soarResultPath)
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

//获取执行文件路径
func GetSoarPath(rootPath string) string {
	var soarPath string
	switch os := runtime.GOOS; os {
	case "windows":
		soarPath = rootPath + "/soar/soar.windows-amd64"
	case "linux":
		soarPath = rootPath + "/soar/soar.windows-amd64"
	default:
		//其他系统
	}
	return soarPath
}

func AnalyzeAndSave(rootPath string, sqlPath string, soarResultPath string) error {
	//获取执行文件的路径
	soarPath := GetSoarPath(rootPath)
	//开始执行分析
	cmd := exec.Command(soarPath, "-query", sqlPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	//结果保存
	os.WriteFile(soarResultPath, []byte(string(out)), 0644)
	//清理残余的临时库表
	defer exec.Command(soarPath, "--cleanup-test-database")

	return nil
}
