package main

import (
	"fmt"

	"github.com/DemoLiang/go-micro/auth/handler"
	"github.com/DemoLiang/go-micro/auth/model"
	"github.com/DemoLiang/go-micro/basic"
	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	s "github.com/DemoLiang/go-micro/auth/proto/auth"
)

func main() {
	basic.Init()

	micReg:=etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.lgm.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) {
			model.Init()
			handler.Init()
			return
		}))

	// Register Handler
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options){
	etcdCfg:=config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%v:%v",etcdCfg.GetHost(),etcdCfg.GetPort())}
}
