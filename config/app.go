// Package config 站点配置信息
package config

import "hk591_go/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "hk591_go"),

			// 当前环境，用以区分多环境
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
