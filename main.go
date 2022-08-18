package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 初始化 Gin 实例
	// r := gin.Default()

	// // 注册一个路由
	// r.GET("/", func(c *gin.Context) {
	// str := ""
	// cmd := exec.Command("soar/soar.windows-amd64", "-print-config")

	cmd := exec.Command("soar/soar.windows-amd64", "-query", "soar/test.sql", ">", "soar/index.html")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	// 以 JSON 格式响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"Hello": "World!",
	// })
	// })

	// 运行服务
	// r.Run()
}
