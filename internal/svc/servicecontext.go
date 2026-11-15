package svc

import (
	"poppulse/internal/config"
	"poppulse/internal/middleware"
	"poppulse/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config   config.Config
	Operator rest.Middleware

	DB        sqlx.SqlConn
	UserModel *model.UserModel
	Redis     *redis.Redis // 可能为 nil：配置未启用时不得使用
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ctx := &ServiceContext{
		Config:    c,
		Operator:  middleware.NewOperatorMiddleware().Handle,
		DB:        conn,
		UserModel: model.NewUserModel(conn),
	}
	if c.Redis.Host != "" {
		ctx.Redis = redis.MustNewRedis(redis.RedisConf{
			Host: c.Redis.Host,
			Type: c.Redis.Type,
			Pass: c.Redis.Pass,
		})
	}
	return ctx
}
