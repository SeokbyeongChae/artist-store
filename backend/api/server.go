package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/seokbyeongchae/artist-store/db/sqlc"
	"github.com/seokbyeongchae/artist-store/security"
	"github.com/seokbyeongchae/artist-store/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker security.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := security.NewJWTMaker(config.JwtTokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// v.RegisterValidation("validator이름", 함수)
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/ping", ping)

	router.POST("/auth/signup", server.signup)
	router.POST("/auth/signin", server.signin)

	authRoutes := router.Group("/").Use(authModdleware(server.tokenMaker))
	authRoutes.GET("auth/ping", authPing)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
