package auth

import (
	db "github.com/seokbyeongchae/go-vue-auth-example/db/sqlc"
	"github.com/seokbyeongchae/go-vue-auth-example/security"
	"github.com/seokbyeongchae/go-vue-auth-example/util"
)

type AuthController struct {
	store      *db.Store
	config     util.Config
	tokenMaker security.Maker
}

func New(store *db.Store, config util.Config, tokenMaker security.Maker) *AuthController {
	return &AuthController{store, config, tokenMaker}
}
