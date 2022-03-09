package main

import (
	"database/sql"
	"demo/api"
	db "demo/db/sqlc"
	"demo/util"
	"log"
)

func main() {
	//viper
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load of config:", err)
	}

	//Connect to MongoDB
	//conf.MongoDB_Conn(config.MongoDBSource)
	//Connect to Redis
	//conf.Redis_Conn(config.RedisAddr, config.RedisDbName, config.RedisPw)

	//TODO: Mock

	conn, err := sql.Open("mysql", config.MysqlDBSource)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("connot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}
