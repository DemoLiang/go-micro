package user

import (
	"fmt"
	proto "github.com/DemoLiang/go-micro/user-srv/proto/user"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	QueryUserByName(username string) (req *proto.User, err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}
