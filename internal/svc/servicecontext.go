package svc

import (
	"poppulse/internal/config"
	"poppulse/internal/middleware"
	"poppulse/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config   config.Config
	Operator rest.Middleware

	DB        sqlx.SqlConn
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		Operator:  middleware.NewOperatorMiddleware().Handle,
		DB:        conn,
		UserModel: model.NewUserModel(conn),
	}
}
