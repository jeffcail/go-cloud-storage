package models

import (
	"fmt"
	"log"

	"github.com/jeffcail/cloud-storage/server/internal/config"

	"github.com/go-redis/redis/v8"

	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

// Init
func Init(dbDsn string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	engine.ShowSQL(true)
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return engine
}

func InitRDB(c config.Config) *redis.Client {
	pass := fmt.Sprintf("%d", c.Redis.Pass)
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: pass,       // no password set
		DB:       c.Redis.Db, // use default DB
	})
	return rdb
}
