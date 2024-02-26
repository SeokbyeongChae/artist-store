package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	my "github.com/go-mysql/errors"
	db "github.com/seokbyeongchae/artist-store/db/sqlc"
	"github.com/seokbyeongchae/artist-store/security"
)

var (
	errNotFound = errors.New("cannot find account")
)

type signupRequest struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type signupResponse struct {
	Result bool `json:"result"`
}

func (server *Server) signup(ctx *gin.Context) {
	var req signupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	salt, saltErr := security.GenerateSalt()
	if saltErr != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(saltErr))
		return
	}

	password := security.HashPassword(salt, req.Password)

	arg := db.CreateAccountParams{
		Salt:     salt,
		Email:    req.Email,
		Password: password,
	}

	_, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if ok, myErr := my.Error(err); ok {
			switch myErr {
			case my.ErrDupeKey:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := signupResponse{
		Result: true,
	}

	ctx.JSON(http.StatusOK, res)
}

type signinRequest struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type signinResponse struct {
	AccessToken string `json:"access_token"`
}

func (server *Server) signin(ctx *gin.Context) {
	var req signinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if isValidPassword := security.CheckPassword(account.Password, account.Salt, req.Password); !isValidPassword {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errNotFound))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(account.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := signinResponse{
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, res)
}
