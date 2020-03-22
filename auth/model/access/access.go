package access

import (
	"fmt"
	"sync"
	r "github.com/go-redis/redis"
	"github.com/DemoLiang/go-micro/basic/redis"
)

var (
	s *service
	ca *r.Client
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	//生成token
	MakeAccessToken(subject *Subject)(ret string,err error)

	//获取token
	GetCachedAccessToken(subject *Subject)(ret string,err error)

	//清除用户token
	DelUserAccessToken(token string)(er error)
}

func GetService()(Service,error){
	if s==nil{
		return nil,fmt.Errorf("[GetService] GetService 为初始化")
	}
	return s,nil
}

func Init(){
	m.Lock()
	defer m.Unlock()
	if s!=nil{
		return
	}
	ca = redis.GetRedis()
	s = &service{}
}
