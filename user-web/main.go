package main

import (
	"fmt"
	"net/http"

	"github.com/DemoLiang/go-micro/basic"
	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	"github.com/DemoLiang/go-micro/user-web/handler"
	"github.com/micro/go-micro/web"
	"github.com/gin-gonic/gin"
)

func main() {

	//初始化配置
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("go.micro.lgm.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		})); err != nil {
		log.Fatal(err)
	}

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))

	router:=gin.New()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})
	router.POST("/user/login",handler.LoginV2)

	service.Handle("/",router)
	// register call handler
	//service.HandleFunc("/user/login", handler.Login)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%v:%v", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
