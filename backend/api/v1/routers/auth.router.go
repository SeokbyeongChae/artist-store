package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/controllers/auth"
	"github.com/seokbyeongchae/go-vue-auth-example/api/v1/middlewares"
	"github.com/seokbyeongchae/go-vue-auth-example/security"
)

type AuthRouter struct {
	controller *auth.AuthController
	tokenMaker security.Maker
}

func New(authController *auth.AuthController, tokenMaker security.Maker) Router {
	return &AuthRouter{authController, tokenMaker}
}

func (router *AuthRouter) RegisterRouter(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/auth")
	group.POST("/signup", router.controller.Signup)
	group.POST("/signin", router.controller.Signin)
	group.POST("/renew_access_token", router.controller.RenewAccessToken)

	authGoup := group.Group("/").Use(middlewares.AuthModdleware(router.tokenMaker))
	authGoup.GET("/ping", router.controller.Ping)
}
