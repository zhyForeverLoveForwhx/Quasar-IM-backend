package main

import (
	//"demo/cache"
	"demo/conf"
	"demo/router"
	"fmt"
)

func main() {
	fmt.Print("Hello World\n")
	//cache.Init()
	conf.Init() //数据库连接初始化
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
