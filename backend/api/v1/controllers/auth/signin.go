package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/constants"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/response"
	db "github.com/seokbyeongchae/go-vue-auth-example/db/sqlc"
	"github.com/seokbyeongchae/go-vue-auth-example/security"
)

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

func (controller *AuthController) Signin(ctx *gin.Context) {
	var req signinRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusBadRequest, err))
		return
	}

	account, err := controller.store.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(response.BuildErrorResponse(http.StatusNotFound, err))
			return
		}

		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
		return
	}

	if isValidPassword := security.CheckPassword(account.Password, account.Salt, req.Password); !isValidPassword {
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrNotFound))
		return
	}

	accessToken, accountPayload, err := controller.tokenMaker.CreateToken(account.ID, controller.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
		return
	}

	refreshToken, refreshPayload, err := controller.tokenMaker.CreateToken(account.ID, controller.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
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

	controller.store.CreateSession(ctx, arg)

	res := signinResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accountPayload.ExpiresAt.Time,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt.Time,
	}

	ctx.JSON(response.BuildSuccessResponse[signinResponse](http.StatusOK, res))
}
