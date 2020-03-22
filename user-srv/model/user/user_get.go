package user

import (
	"github.com/DemoLiang/go-micro/basic/db"
	proto "github.com/DemoLiang/go-micro/user-srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `select user_id,user_name,pwd from user where user_name = ?`

	o := db.GetDB()
	if o==nil{
		log.Error("o is nil")
	}
	ret = &proto.User{}
	log.SetLevel(log.LevelDebug)
	err = o.QueryRow(queryString, userName).Scan(&ret.Pwd, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Error("[QueryUserByName] 查询数据失败：err:%v", err)
		return
	}
	log.Info("%v",ret)
	return
}
