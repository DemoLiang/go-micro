package access

import (
	"fmt"
	"time"

	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/broker"
	"github.com/prometheus/common/log"
)

var (
	tokenExpiredDate = 3600*24*30*time.Second
	tokenIDKeyPrefix = "token:auth:id:"
	tokenExpiredTopic = "go.micro.lgm.topic.auth.tokenExpired"
)

//subject token持有者
type Subject struct {
	Id string `json:"id"`
	Name string `json:"name"`
}


func (s *service)MakeAccessToken(subject *Subject)(ret string,err error){
	m,err:=s.createTokenClaims(subject)
	if err!=nil{
		return "",fmt.Errorf("[MakeAccessToken] 创建token claim 失败")
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,m)
	ret,err= token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err!=nil{
		return "",fmt.Errorf("[MakeAccessToken]创建token失败",err.Error())
	}
	err = s.saveTokenCache(subject,ret)
	if err!=nil{
		return "",fmt.Errorf("[MakeAccessToken]保存到缓存失败",err.Error())
	}
	return
}

func (s *service)GetCachedAccessToken(subject *Subject)(ret string,err error){
	ret,err=s.getTokenFromCache(subject)
	if err!=nil{
		return "",fmt.Errorf("[GetChachedAccessToken] 从缓存获取token失败")
	}
	return
}

func (s *service)DelUserAccessToken(tk string)(err error){
	claims,err:=s.parseToken(tk)
	if err!=nil{
		return fmt.Errorf("[DelUserAccessToken]错误的Token",err.Error())
	}
	err = s.delTokenFromCache(&Subject{
		Id:claims.Subject,
	})
	if err!=nil{
		return fmt.Errorf("[DelUserAccessToken]清除用户token",err.Error())
	}
	msg:=&broker.Message{
		Body:[]byte(claims.Subject),
	}
	if err=broker.Publish(tokenExpiredTopic,msg);err!=nil{
		log.Info("[pub]发布token删除消息失败",err.Error())
	}else{
		log.Info("[pub]发布token删除消息",string(msg.Body))
	}
	return
}