package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/gin-gonic/gin"

	db "simple_bank/db/sqlc"
)

func (s *Server) createAccount(ctx *gin.Context) {
	type createAccountRequest struct {
		Owner    string `json:"owner" binding:"required"`
		Currency string `json:"currency" binding:"required,currency"`
	}
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Printf("postgres sql err message is '%s' \n", pgErr.Message)
			fmt.Printf("postgres sql err code is '%s' \n", pgErr.Code)

			switch pgErr.Code {
			case "23503":
				ctx.JSON(http.StatusForbidden, gin.H{
					"error": pgErr.Error(),
				})
				return
			case "23505":
				ctx.JSON(http.StatusForbidden, gin.H{
					"error": pgErr.Error(),
				})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// 查询用户
func (s *Server) getAccount(ctx *gin.Context) {
	type getAccountRequest struct {
		ID int64 `form:"id" uri:"id" binding:"required,gte=1"`
	}
	var req getAccountRequest
	// 绑定id到结构体
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		// 数据库不存在的情况
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		// 服务器内部错误
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (s *Server) listAccount(ctx *gin.Context) {
	type listAccountRequest struct {
		PageID   uint32 `form:"page_id" binding:"required,gte=1"`
		PageSize uint32 `form:"page_size" binding:"required,gte=5,lte=20"`
	}
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAccountsParams{
		Limit:  int64(req.PageSize),
		Offset: int64((req.PageID - 1) * req.PageSize),
	}
	accounts, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}
