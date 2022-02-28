package router

import (
	"demo/api"
	"demo/model"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

var user api.User  //保存数据库的数据&user
var user2 api.User //保存传递过来的数据

func NewRouter() *gin.Engine {
	r := gin.Default()
	//Recovery中间件会恢复(recovers) 任何恐慌(panics) 如果存在恐慌
	//中间件会写入500 这个中间件是十分有必要的
	//Logger日志
	r.Use(gin.Recovery(), gin.Logger())
	v1 := r.Group("/")
	{
		v1.POST("login", func(c *gin.Context) {
			c.Bind(&user2)
			logging.Info(user2)
			result := model.DB.Where("username = ?", user2.Username).Table("users").First(&user)
			logging.Info(user)
			if result.Error != nil {
				c.JSON(404, "this username does not exist")
			} else {
				if user.Password != user2.Password {
					c.JSON(400, "Wrong with this username and Password")
					return
				}
				c.JSON(200, "login success")
			}

		})
	}
	return r
}
