package v1

import (
	"go_web/pkg/response"

	"github.com/gin-gonic/gin"
)

type SoarController struct {
	BaseAPIController
}

func (ctrl *SoarController) Index(c *gin.Context) {
	// cmd := exec.Command("soar/soar.windows-amd64", "-query", "soar/test.sql")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	response.BadRequest(c, err, "")
	// }
	// file := "./soar/index.html"
	// os.WriteFile(file, []byte(string(out)), 0644)
	data := map[string]string{
		"url": "a",
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}
