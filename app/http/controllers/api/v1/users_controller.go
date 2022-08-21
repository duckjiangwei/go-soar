package v1

import (
	"hk591_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data := map[string]string{
		"name": "a",
		"sex":  "man",
	}
	response.Data(c, 200, gin.H{
		"data": data,
	})
}
