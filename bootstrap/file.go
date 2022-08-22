package bootstrap

import (
	"go_web/pkg/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func SetupFile(router *gin.Engine) {
	rootPath, _ := os.Getwd()
	//分析结果
	router.Static("/soar-result", rootPath+"/"+cast.ToString(config.Env("SOAR_RESULT")))
}
