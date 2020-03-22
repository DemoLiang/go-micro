package config

import (
	"strings"
)

type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}


type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}

type defaultRedisConfig struct {
	Enabled bool `json:"enabled"`
	Conn string `json:"conn"`
	Password string `json:"password"`
	DBNum int `json:"db_num"`
	Timeout int `json:"timeout"`
	Sentinel redisSentinel `json:"sentinel"`
}


type redisSentinel struct {
	Enabled bool `json:"enabled"`
	Master string `json:"master"`
	Nodes string `json:"nodes"`
	nodes []string `json:"nodes"`
}

func (r defaultRedisConfig)GetEnabled()bool{
	return r.Enabled
}

func (r defaultRedisConfig)GetConn()string{
	return r.Conn
}

func (r defaultRedisConfig)GetPassword()string{
	return r.Password
}

//获取数据库分区序号
func (r defaultRedisConfig)GetDBNum()int{
	return r.DBNum
}

func (r defaultRedisConfig)GetSentinelConfig()RedisSentinelConfig{
	return r.Sentinel
}

func (s redisSentinel)GetEnabled()bool{
	return s.Enabled
}
func (s redisSentinel)GetMaster()string{
	return s.Master
}

func (s redisSentinel)GetNodes()[]string{
	if len(s.Nodes)!=0{
		for _,v:=range strings.Split(s.Nodes,","){
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes,v)
		}
	}
	return s.nodes
}
