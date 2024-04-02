package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mysql "github.com/go-mysql/errors"
	"github.com/seokbyeongchae/artist-store/api/v1/response"
	db "github.com/seokbyeongchae/artist-store/db/sqlc"
	"github.com/seokbyeongchae/artist-store/security"
)

type signupRequest struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type signupResponse struct {
	Result bool `json:"result"`
}

func (controller *AuthController) Signup(ctx *gin.Context) {
	var req signupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusBadRequest, err))
		return
	}

	salt, saltErr := security.GenerateSalt()
	if saltErr != nil {
		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, saltErr))
		return
	}

	password := security.HashPassword(salt, req.Password)

	arg := db.CreateAccountParams{
		Salt:     salt,
		Email:    req.Email,
		Password: password,
	}

	_, err := controller.store.CreateAccount(ctx, arg)
	if err != nil {
		if ok, myErr := mysql.Error(err); ok {
			switch myErr {
			case mysql.ErrDupeKey:
				ctx.JSON(response.BuildErrorResponse(http.StatusForbidden, err))
				return
			}
		}

		ctx.JSON(response.BuildErrorResponse(http.StatusInternalServerError, err))
		return
	}

	res := signupResponse{
		Result: true,
	}

	ctx.JSON(response.BuildSuccessResponse[signupResponse](http.StatusOK, res))
}
