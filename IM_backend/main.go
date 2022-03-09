package main

import (
	"demo/conf"
	"demo/server"
	"demo/util"
	"log"
)

func main() {
	//viper
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load of config:", err)
	}
	//Connect to Mysql
	conf.Mysql_Conn(config.MysqlDBSource)
	//Connect to MongoDB
	conf.MongoDB_Conn(config.MongoDBSource)
	//Connect to Redis
	conf.Redis_Conn(config.RedisAddr, config.RedisDbName, config.RedisPw)

	//TODO: Mock

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("connot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}
