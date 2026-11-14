package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// JWT 鉴权配置，路由 @server(jwt: Auth) 引用本字段
	Auth struct {
		AccessSecret string // 生产环境通过环境变量注入，不得硬编码
		AccessExpire int64  // 过期秒数
	}

	// MySQL 数据源，例：user:pass@tcp(127.0.0.1:3306)/poppulse?charset=utf8mb4&parseTime=true&loc=Local
	Mysql struct {
		DataSource string
	}

	// 订单支付期限（分钟），默认 15
	PayTimeoutMinutes int `json:",default=15"`

	// 模拟能力开关：模拟支付/退款必须显式启用，生产环境必须为 false
	Mock struct {
		Enabled bool
	}
}
