package main

import (
	"fmt"

	"github.com/DemoLiang/go-micro/basic"
	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/DemoLiang/go-micro/user-srv/handler"
	"github.com/DemoLiang/go-micro/user-srv/model"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	user "github.com/DemoLiang/go-micro/user-srv/proto/user"
)

func main() {

	// 初始化配置,数据库配置等配置信息
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.lgm.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			return
		}),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
