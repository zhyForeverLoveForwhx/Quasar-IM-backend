package api

import (
	db "demo/db/sqlc"
	"demo/middleware"
	"demo/token"
	"demo/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// v.RegisterValidation("currency", validCurrency)
	// v.RegisterValidation("type", validType)
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	v1 := router.Group("/").Use(middleware.Cors())
	{
		v1.POST("login", server.Login)
		v1.POST("verify", server.Verify)
		v1.GET("get_conv", server.Get_conv)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
