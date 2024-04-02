package routers

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRouter(routerGroup *gin.RouterGroup)
}
