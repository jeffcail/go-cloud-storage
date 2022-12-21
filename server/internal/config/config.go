package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DbDsn string
	}
	Redis struct {
		Addr string
		Pass int
		Db   int
	}
}
