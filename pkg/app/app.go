// Package app 应用信息
package app

import (
	"hk591_go/pkg/config"
)

func IsLocal() bool {
	return config.Get("app.env") == "dev"
}

func IsProduction() bool {
	return config.Get("app.env") == "online"
}

func IsTesting() bool {
	return config.Get("app.env") == "debug"
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
