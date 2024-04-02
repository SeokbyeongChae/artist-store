package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/artist-store/api/v1/constants"
	"github.com/seokbyeongchae/artist-store/api/v1/response"
	"github.com/seokbyeongchae/artist-store/security"
)

func AuthModdleware(tokenMaker security.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(constants.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrHeaderIsNotProvided))
			return
		}

		fileds := strings.Fields(authorizationHeader)
		if len(fileds) < 2 {
			ctx.AbortWithStatusJSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrInvalidHeaderFormat))
			return
		}

		authorizationType := strings.ToLower(fileds[0])
		if authorizationType != constants.AuthorizationTypeBearer {
			ctx.AbortWithStatusJSON(response.BuildErrorResponse(http.StatusUnauthorized, constants.ErrInvalidHeaderFormat))
			return
		}

		accessToken := fileds[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(response.BuildErrorResponse(http.StatusUnauthorized, err))
			return
		}

		ctx.Set(constants.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
