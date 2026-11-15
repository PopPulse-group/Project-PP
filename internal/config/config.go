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

	// MySQL 数据源，例：user:pass@tcp(127.0.0.1:3306)/fresh?charset=utf8mb4&parseTime=true&loc=Local
	Mysql struct {
		DataSource string
	}

	// Redis（商品缓存等读取加速用，不承载库存事实）；Host 为空表示未启用
	Redis struct {
		Host string `json:",optional"`
		Type string `json:",default=node"` // node | cluster
		Pass string `json:",optional"`
	}

	// 订单支付期限（分钟），默认 15
	PayTimeoutMinutes int `json:",default=15"`

	// 模拟能力开关：模拟支付/履约推进必须显式启用，生产环境必须为 false
	Mock struct {
		Enabled bool
	}

	// AI 助手：provider=mock|real；真实模型接入需配置 api key（通过环境变量注入）
	Ai struct {
		Provider      string `json:",default=mock"`
		ApiKey        string `json:",optional"`
		Model         string `json:",optional"`
		TimeoutMs     int64  `json:",default=10000"`
		MaxConcurrent int64  `json:",default=5"` // AI 并发信号量，隔离于其他接口
	}
}
