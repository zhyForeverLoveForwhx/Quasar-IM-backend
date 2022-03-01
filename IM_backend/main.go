package main

import (
	"demo/conf"
	"demo/router"
	"demo/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load of config:", err)
	}
	//Connect to Mysql
	conf.Mysql_Conn(config.MysqlDBSource)
	//Connect to MongoDB
	conf.MongoDB_Conn(config.MongoDBSource)
	conf.Redis_Conn(config.RedisAddr, config.RedisDbName, config.RedisPw)
	//conf.Init() //数据库连接初始化
	r := router.NewRouter()
	_ = r.Run(config.HttpPort)
}
