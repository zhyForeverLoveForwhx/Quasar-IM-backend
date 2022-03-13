package api

import (

	"github.com/gin-gonic/gin"
)

//返回错误信息 ErrorResponse
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

