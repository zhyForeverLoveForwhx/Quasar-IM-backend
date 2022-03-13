package api

import (
	"database/sql"
	"net/http"

	db "github.com/Awadabang/Quasar-IM/db/sqlc"
	"github.com/Awadabang/Quasar-IM/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) Register(ctx *gin.Context) {
	//binding the request of register
	var req Request_register
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// make the password to hash
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//create the user in DB
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}
	// insert the user into DB
	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				//the username is not unique
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

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

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := Response_login{
		AccessToken: accessToken,
		Username:    user.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) Verify(c *gin.Context) {
	c.JSON(200, c.GetHeader("Authorization"))
}
