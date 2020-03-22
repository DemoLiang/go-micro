package handler

import (
	"context"
	"strconv"

	"github.com/DemoLiang/go-micro/auth/model/access"
	"github.com/prometheus/common/log"
	auth "github.com/DemoLiang/go-micro/auth/proto/auth"
)

var (
	accessService access.Service
)

func Init(){
	var err error
	accessService,err=access.GetService()
	if err!=nil{
		log.Fatal("[Init] 初始化handler错误",err.Error())
		return
	}
}

type Service struct {
}


func (s *Service)MakeAccessToken(ctx context.Context,req *auth.Request,rsp *auth.Response)(err error){
	log.Info("[MakeAccessToken] 收到创建token请求")

	token,err:=accessService.MakeAccessToken(&access.Subject{
		Id:strconv.FormatUint(req.UserId,10),
		Name:req.UserName,
	})
	if err!=nil{
		rsp.Error = &auth.Error{
			Detail:err.Error(),
		}
		log.Error("[MakeAccessToken] token生成失败",err.Error())
		return err
	}
	rsp.Token = token
	return nil
}

func (s *Service)DelUserAccessToken(ctx context.Context,req *auth.Request,rsp *auth.Response)(err error){
	log.Info("[DelUserAccessToken]清除用户token")
	err=accessService.DelUserAccessToken(req.Token)
	if err!=nil{
		rsp.Error = &auth.Error{
			Detail:err.Error(),
		}
		log.Info("[DelUserAccessToken]清除用户Token失败",err.Error())
		return err
	}
	return nil
}