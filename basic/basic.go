package basic

import (
	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/DemoLiang/go-micro/basic/db"
)

func Init() {
	config.Init()
	if config.GetMysqlConfig().GetEnabled() {
		db.Init()
	}
}
