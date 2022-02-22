package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

// Redis 初始化redis链接
func Init() {
	file, err := ini.Load("D:/go_project/Quasar-IM-backend/IM_backend/conf/conf.ini")
	if err != nil {
		fmt.Println("Redis 配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedisData(file) //读取配置信息
	Redis()             //redis链接
}

//Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) //string to uint64
	client := redis.NewClient(&redis.Options{       //登录Redis
		Addr: RedisAddr,
		//Password: conf.RedisPw,  // 无密码，就这样就好了
		DB: int(db),
	})
	_, err := client.Ping().Result() //验证是否ping通
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}

//读取配置信息
func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
