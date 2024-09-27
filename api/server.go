package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "simple_bank/db/sqlc"
	"simple_bank/pkg"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := validate.RegisterValidation("currency", pkg.ValidateCurrency)
		if err != nil {
			log.Fatalf("error registering validation: %v", err)
		}
	}

	// 创建单个用户
	router.PUT("/users", server.CreateUser)

	// 创建单个账户
	router.PUT("/accounts", server.createAccount)
	// 获取单个账户信息
	router.GET("/accounts/:id", server.getAccount)
	// 获取账户列表信息
	router.GET("/accounts", server.listAccount)

	// 创建转账记录
	router.PUT("/transfers", server.createTransfer)

	server.router = router

	return server
}

// Start 启动
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
