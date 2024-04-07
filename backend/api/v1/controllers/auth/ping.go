package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/constants"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/response"
	"github.com/seokbyeongchae/go-vue-auth-example/security"
)

type PingResponse struct {
	Result int64 `json:"result"`
}

func (controller *AuthController) Ping(ctx *gin.Context) {
	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*security.Payload)

	res := PingResponse{
		Result: authPayload.AccountId,
	}

	ctx.JSON(response.BuildSuccessResponse[PingResponse](http.StatusOK, res))
}
