package api

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"

	"simple_bank/pkg"

	"github.com/gin-gonic/gin"

	db "simple_bank/db/sqlc"
)

func (s *Server) CreateUser(ctx *gin.Context) {
	type CreateUserRequest struct {
		Username string `json:"username" binding:"required"`
		FullName string `json:"fullName" binding:"required"`
		Password string `json:"password" binding:"required,gte=6"`
		Email    string `json:"email" binding:"required,email"`
	}
	type CreateUserResponse struct {
		Username string `json:"username" binding:"required"`
		FullName string `json:"fullName" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := pkg.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		HashedPassword: password,
		Email:          req.Email,
	}

	user, createErr := s.store.CreateUser(ctx, arg)
	if createErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(createErr, &pgErr) {
			switch pgErr.Code {
			case "23505":
				ctx.JSON(http.StatusForbidden, gin.H{
					"message": pgErr.Message,
					"code":    pgErr.Code,
					"body":    "用户名已存在",
				})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := CreateUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
	ctx.JSON(http.StatusOK, rsp)
}
