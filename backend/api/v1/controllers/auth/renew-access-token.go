package auth

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/constants"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/response"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (controller *AuthController) RenewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusBadRequest, err))
		return
	}

	refreshPayload, err := controller.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, err))
		return
	}

	session, err := controller.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(response.BuildErrorResponse(http.StatusNotFound, err))
			return
		}

		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
		return
	}

	if time.Now().After(session.ExpireAt) {
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrUnauthorized))
		return
	}

	if session.IsBlocked {
		// TODO
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrUnauthorized))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		// TODO
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrUnauthorized))
		return
	}

	if session.UserAgent != ctx.Request.UserAgent() {
		// TODO
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrUnauthorized))
		return
	}

	if session.ClientIp != ctx.ClientIP() {
		// TODO
		ctx.JSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrUnauthorized))
		return
	}

	accessToken, accountPayload, err := controller.tokenMaker.CreateToken(session.AccountID, controller.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
		return
	}

	res := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accountPayload.ExpiresAt.Time,
	}

	ctx.JSON(response.BuildSuccessResponse[renewAccessTokenResponse](http.StatusOK, res))
}
