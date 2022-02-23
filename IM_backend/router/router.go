package router

import (
	"demo/api"
	"demo/model"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

var user api.User

func NewRouter() *gin.Engine {
	r := gin.Default()
	//Recovery中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌
	//中间件会写入500 这个中间件是十分有必要的
	//Logger日志
	r.Use(gin.Recovery(), gin.Logger())
	v1 := r.Group("/")
	{
		v1.POST("login", func(c *gin.Context) {
			c.Bind(&user)
			result := model.DB.Table("users").First(&user)
			logging.Info(result.RowsAffected)
			c.JSON(200, "success")
		})
	}
	return r
}
