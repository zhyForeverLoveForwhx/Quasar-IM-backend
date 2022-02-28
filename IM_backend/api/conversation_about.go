package api

import "github.com/gin-gonic/gin"

func Get_conv(c *gin.Context) {
	var Token string
	c.Bind(Token)
	if Token != "nil" {
		c.JSON(200, nil)
	}
}
