package api

import (
	"github.com/gin-gonic/gin"
	db "simple_bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// routes:
	router.PUT("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)

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
