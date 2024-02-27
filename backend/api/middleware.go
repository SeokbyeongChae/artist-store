package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/artist-store/security"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

var (
	errHeaderIsNotProvided = errors.New("authorization header is not provided")
	errInvalidHeaderFormat = errors.New("invalid authorization header format")
)

func authModdleware(tokenMaker security.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errHeaderIsNotProvided))
			return
		}

		fileds := strings.Fields(authorizationHeader)
		if len(fileds) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errInvalidHeaderFormat))
			return
		}

		authorizationType := strings.ToLower(fileds[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(errInvalidHeaderFormat))
			return
		}

		accessToken := fileds[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
