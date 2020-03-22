package handler

import (
	"context"
	"log"

	"github.com/DemoLiang/go-micro/user-srv/model/user"
	proto "github.com/DemoLiang/go-micro/user-srv/proto/user"
)

type Service struct {
}

var (
	userService user.Service
)

func Init() {
	var err error
	userService, err = user.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}
	rsp.User = user
	rsp.Success = true
	return nil
}
