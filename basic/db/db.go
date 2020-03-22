package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/micro/go-micro/util/log"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error
	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Error(err)
		return
	}
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	log.Error("%v",config.GetMysqlConfig().GetEnabled())
	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
