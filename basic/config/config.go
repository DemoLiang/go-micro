package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePredix = "application-"
	etcdConfig              defaultEtcdConfig
	mysqlConfig             defaultMysqlConfig
	profiles                defaultProfiles
	jwtConfig defaultJwtConfig
	redisConfig defaultRedisConfig
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] 配置已经初始化过")
		return
	}

	//加载yml配置
	//先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))

	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)

	//找到application.yml 文件
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	//找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Info("[Init] 加载配置文件path:%s,%s\n", pt+sp+"application.yml", profiles)

	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")
		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filepath := pt + string(filepath.Separator) + defaultConfigFilePredix + strings.TrimSpace(include[i]+".yml")
			log.Info("[Init] 加载配置文件：path:%s\n", filepath)
			sources[i] = file.NewSource(file.WithPath(filepath))
		}
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}
	//赋值
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)
	config.Get(defaultRootPath,"redis").Scan(&redisConfig)
	config.Get(defaultRootPath,"jwt").Scan(&jwtConfig)
	//标记已经初始化
	inited = true
}

func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}

func GetRedisConfig()(ret RedisConfig){
	return redisConfig
}

func GetJwtConfig()(ret JwtConfig){
	return jwtConfig
}