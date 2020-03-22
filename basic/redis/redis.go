package redis

import (
	"sync"

	"github.com/DemoLiang/go-micro/basic/config"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
)


var (
	client *redis.Client
	m sync.RWMutex
	inited bool
)

func Init(){
	m.Lock()
	defer m.Unlock()

	if inited{
		log.Error("已经初始化过redis")
		return
	}

	redisConfig:=config.GetRedisConfig()

	//打开才加载
	if redisConfig!=nil&&redisConfig.GetEnabled(){
		log.Info("初始化reids")

		//加载哨兵模式
		if redisConfig.GetSentinelConfig()!=nil&&redisConfig.GetSentinelConfig().GetEnabled(){
			log.Info("初始化redis,哨兵模式...")
			initSentinel(redisConfig)
		}else{
			log.Info("初始化redis,普通模式...")
			initSingle(redisConfig)
		}

		log.Info("初始化redis,检测链接...")
		pong,err:=client.Ping().Result()
		if err!=nil{
			log.Fatal(err.Error())
		}
		log.Info("初始化redis,检测链接ping.",pong)
	}
	inited = true
}


func GetRedis() *redis.Client{
	return client
}


func initSentinel(redisConfig config.RedisConfig){
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs:redisConfig.GetSentinelConfig().GetNodes(),
		DB:redisConfig.GetDBNum(),
		Password:redisConfig.GetPassword(),
	})
}

func initSingle(redisConfig config.RedisConfig){
	log.Info("%v",redisConfig)
	client = redis.NewClient(&redis.Options{
		Addr:redisConfig.GetConn(),
		Password:redisConfig.GetPassword(),
		DB:redisConfig.GetDBNum(),
	})
}