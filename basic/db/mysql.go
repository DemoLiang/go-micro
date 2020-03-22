package db

import (
	"database/sql"
	"time"

	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/micro/go-micro/util/log"
	_ "github.com/go-sql-driver/mysql"
)

func initMysql() {
	var err error
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetUrl())
	if err != nil {
		log.Error(err)
		panic(err)
	}
	log.Error("mysql db open ok")

	//最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMalIdleConnection())

	mysqlDB.SetConnMaxLifetime(time.Second * config.GetMysqlConfig().GetConnMaxLifetime())

	//激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
