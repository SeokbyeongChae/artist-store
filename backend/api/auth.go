package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	my "github.com/go-mysql/errors"
	db "github.com/seokbyeongchae/artist-store/db/sqlc"
	"github.com/seokbyeongchae/artist-store/security"
)

var (
	errNotFound     = errors.New("cannot find account")
	errUnauthorized = errors.New("session is unauthorized")
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
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
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

	accessToken, accountPayload, err := server.tokenMaker.CreateToken(account.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(account.ID, server.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateSessionParams{
		AccountID:    account.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpireAt:     refreshPayload.ExpiresAt.Time,
	}

	server.store.CreateSession(ctx, arg)

	res := signinResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accountPayload.ExpiresAt.Time,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt.Time,
	}

	ctx.JSON(http.StatusOK, res)
}

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpireAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errUnauthorized))
		return
	}

	if session.IsBlocked {
		// TODO
		ctx.JSON(http.StatusUnauthorized, errorResponse(errUnauthorized))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		// TODO
		ctx.JSON(http.StatusUnauthorized, errorResponse(errUnauthorized))
		return
	}

	if session.UserAgent != ctx.Request.UserAgent() {
		// TODO
		ctx.JSON(http.StatusUnauthorized, errorResponse(errUnauthorized))
		return
	}

	if session.ClientIp != ctx.ClientIP() {
		// TODO
		ctx.JSON(http.StatusUnauthorized, errorResponse(errUnauthorized))
		return
	}

	accessToken, accountPayload, err := server.tokenMaker.CreateToken(session.AccountID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accountPayload.ExpiresAt.Time,
	}

	ctx.JSON(http.StatusOK, res)
}

func authPing(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*security.Payload)
	ctx.JSON(http.StatusOK, gin.H{
		"message": authPayload.AccountId,
	})
}
