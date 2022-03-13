package api

import (
	"database/sql"
	"demo/util"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (server *Server) Login(ctx *gin.Context) {
	var req Request_login
		if err := ctx.ShouldBindJSON(&req); err != nil {
		//TODO: apifox update the error 400
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			//TODO: apifox update the error 404
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		//TODO: apifox update the error 500
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		//TODO: apifox update the error 401
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
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
	var Token string
	c.Bind(Token)
	if Token != "nil" {
		c.JSON(200, nil)
	}
}

func (server *Server) Getusername(){

}