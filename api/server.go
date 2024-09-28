package api

import (
	"fmt"
	"log"
	"simple_bank/middleware"

	"simple_bank/config"

	"simple_bank/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "simple_bank/db/sqlc"
)

type Server struct {
	config    *config.Config
	store     db.Store
	tokenMake token.Maker
	router    *gin.Engine
}

func NewServer(config *config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error creating token maker: %w", err)
	}

	server := &Server{
		config:    config,
		store:     store,
		tokenMake: tokenMaker,
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.Cors())

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := validate.RegisterValidation("currency", validCurrency)
		if err != nil {
			log.Fatalf("error registering validation: %v", err)
		}
	}

	// 创建单个用户
	router.PUT("/users", server.CreateUser)
	// 查询单个用户
	router.GET("/users", server.GetUser)

	// 用户登录
	router.POST("/users/login", server.loginUser)

	// 创建单个账户
	router.PUT("/accounts", server.createAccount)
	// 获取单个账户信息
	router.GET("/accounts/:id", server.getAccount)
	// 获取账户列表信息
	router.GET("/accounts", server.listAccount)

	// 创建转账记录
	router.PUT("/transfers", server.createTransfer)

	server.router = router

	return server, nil
}

// Start 启动
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
