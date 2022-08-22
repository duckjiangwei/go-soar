package services

import (
	"go_web/pkg/config"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

//使用 soar 分析并保存结果
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

//获取执行文件路径
func GetSoarPath(rootPath string) string {
	var soarPath string
	switch os := runtime.GOOS; os {
	case "windows":
		soarPath = rootPath + "/soar/soar.windows-amd64"
	case "linux":
		soarPath = rootPath + "/soar/soar.linux-amd64"
	default:
		//其他系统
	}
	return soarPath
}

//处理文件上传
func UploadFile(c *gin.Context, rootPath string) (string, error) {
	file, errLoad := c.FormFile("file")
	if errLoad != nil {
		return "", errLoad
	}
	ret := make(map[string]string)
	ret["file_name"] = cast.ToString(time.Now().Unix()) + "_" + file.Filename
	ret["file_name_origin"] = file.Filename
	sqlPath := rootPath + "/" + cast.ToString(config.Env("SOAR_SQL")) + ret["file_name"]

	err := c.SaveUploadedFile(file, sqlPath)
	if err != nil {
		return "", err
	}
	return sqlPath, nil
}
