package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()
	//Recovery中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌
	//中间件会写入500 这个中间件是十分有必要的
	//Logger日志
	r.Use(gin.Recovery(), gin.Logger())
	v1 := r.Group("/")
	{
		v1.GET("Ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return r
}
