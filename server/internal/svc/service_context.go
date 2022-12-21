package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/jeffcail/cloud-storage/server/internal/config"
	"github.com/jeffcail/cloud-storage/server/internal/middleware"
	"github.com/jeffcail/cloud-storage/server/models"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DbDsn),
		RDB:    models.InitRDB(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
