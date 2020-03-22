package basic

import (
	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/DemoLiang/go-micro/basic/db"
	"github.com/DemoLiang/go-micro/basic/redis"
)

func Init() {
	config.Init()
	if config.GetMysqlConfig().GetEnabled() {
		db.Init()
	}
	if config.GetRedisConfig().GetEnabled(){
		redis.Init()
	}
}
