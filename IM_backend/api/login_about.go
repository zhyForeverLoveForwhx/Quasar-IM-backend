package api

import (
	"demo/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user_searched User //保存数据库的数据&user
	var user User          //保存传递过来的数据
	c.Bind(&user)
	result := model.DB.Where("username = ?", user.Username).Table("users").First(&user_searched)
	if result.Error != nil {
		c.JSON(404, nil)
	} else {
		if user_searched.Password != user.Password {
			c.JSON(400, nil)
			return
		}
		response := Response_login{Username: user_searched.Username, Token: "token"}
		c.JSON(200, response)
	}

}

func Verify(c *gin.Context) {
	var Token string
	c.Bind(Token)
	if Token != "nil" {
		c.JSON(200, nil)
	}
}
