package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "simple_bank/db/sqlc"
)

type CreateAccountParams struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EDR CNY"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountParams
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}
	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// 查询用户
func (s *Server) getAccount(context *gin.Context) {
	id := context.Param("id")
	s.store.GetAccount()
}
