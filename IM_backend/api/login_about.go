package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) Login(c *gin.Context) {
	// var user_searched User //保存数据库的数据&user
	// var user User          //保存传递过来的数据
	// c.Bind(&user)

	// if result.Error != nil {
	// 	c.JSON(404, nil)
	// } else {
	// 	if user_searched.Password != user.Password {
	// 		c.JSON(400, nil)
	// 		return
	// 	}
	// 	response := Response_login{Username: user_searched.Username, Token: "token"}
	// 	c.JSON(200, response)
	// }
}

func (server *Server) Verify(c *gin.Context) {
	c.JSON(200, c.GetHeader("Authorization"))
}
