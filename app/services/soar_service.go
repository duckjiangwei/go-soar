package services

import (
	"fmt"
	"go_web/app/requests"
	"go_web/pkg/config"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

//使用 soar 分析并保存结果
func AnalyzeAndSave(request requests.SoarSqlRequest, rootPath string, sqlPath string, soarResultPath string) error {
	//获取执行文件的路径
	soarPath := GetSoarPath(rootPath)
	//拼装dns信息
	dns := GetDnsInfo(request)
	//开始执行分析
	cmd := exec.Command(soarPath, "-test-dsn", dns, "-query", sqlPath)
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
		//mac
		soarPath = rootPath + "/soar/soar.darwin-amd64"
	}
	return soarPath
}

//处理文件上传
func UploadFile(c *gin.Context, rootPath string) (string, error) {
	file, errLoad := c.FormFile("sql_file")
	if errLoad != nil {
		return "", errLoad
	}
	ret := make(map[string]string)
	ret["file_name"] = cast.ToString(time.Now().Unix()) + "_" + file.Filename
	ret["file_name_origin"] = file.Filename
	sqlPath := rootPath + "/" + cast.ToString(config.Env("SOAR_SQL")) + "/" + ret["file_name"]

	err := c.SaveUploadedFile(file, sqlPath)
	if err != nil {
		return "", err
	}
	return sqlPath, nil
}

//获取dns信息
func GetDnsInfo(request requests.SoarSqlRequest) string {
	//取默认dns信息
	username := cast.ToString(config.Env("DB_USERNAME"))
	password := cast.ToString(config.Env("DB_PASSWORD"))
	host := cast.ToString(config.Env("DB_HOST"))
	port := cast.ToString(config.Env("DB_PORT"))
	db := cast.ToString(config.Env("DB_DATABASE"))
	//覆盖默认dns信息
	if request.Username != "" {
		username = request.Username
	}
	if request.Password != "" {
		password = request.Password
	}
	if request.Host != "" {
		host = request.Host
	}
	if request.Port != "" {
		port = request.Port
	}

	if request.Db != "" {
		db = request.Db
	}
	dns := username + ":" + password + "@" + host + ":" + port + "/" + db
	fmt.Println(dns)
	return dns
}
