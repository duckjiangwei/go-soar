package v1

import (
	"fmt"
	"go_web/pkg/config"
	"go_web/pkg/response"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type SoarController struct {
	BaseAPIController
}

func (ctrl *SoarController) Index(c *gin.Context) {
	rootPath, _ := os.Getwd()
	soarPath := rootPath + "/soar/soar.windows-amd64"
	sqlPath := rootPath + "/soar/test.sql"
	cmd := exec.Command(soarPath, "-query", sqlPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		response.BadRequest(c, err, "")
	}
	soarResultPath := rootPath + "/" + cast.ToString(config.Env("SOAR_RESULT")) + "/index.html"

	fmt.Println(soarResultPath)
	os.WriteFile(soarResultPath, []byte(string(out)), 0644)
	url := cast.ToString(config.Env("APP_URL")) + "/soar-result/index.html"

	data := map[string]string{
		"url": url,
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}
