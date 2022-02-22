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
	conf.Init()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
