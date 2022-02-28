package main

import (
	//"demo/cache"
	"demo/conf"
	"demo/router"
)

func main() {
	//cache.Init()
	conf.Init() //数据库连接初始化
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
