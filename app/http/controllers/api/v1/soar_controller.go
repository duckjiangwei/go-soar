package v1

import (
	"fmt"
	"go_web/pkg/config"
	"go_web/pkg/response"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type SoarController struct {
	BaseAPIController
}

func (ctrl *SoarController) Index(c *gin.Context) {
	rootPath, _ := os.Getwd()
	//获取执行文件的路径
	soarPath := rootPath + "/soar/soar.windows-amd64"
	//获取待分析的 sql 文件
	sqlPath := rootPath + "/" + cast.ToString(config.Env("SOAR_SQL"))
	sqlFileName := sqlPath + "/test.sql"
	fmt.Println(sqlFileName)
	//开始执行分析
	cmd := exec.Command(soarPath, "-query", sqlFileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		response.BadRequest(c, err, "")
	}
	//结果路径
	randFileName := cast.ToString(time.Now().Unix())
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/" + randFileName + ".html"
	//结果保存
	os.WriteFile(soarResultPath, []byte(string(out)), 0644)

	//ajax 返回
	url := cast.ToString(config.Env("APP_URL")) + "/soar-result/" + randFileName + ".html"
	data := map[string]string{
		"url": url,
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}
