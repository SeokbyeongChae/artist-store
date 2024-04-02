package auth

import (
	db "github.com/seokbyeongchae/artist-store/db/sqlc"
	"github.com/seokbyeongchae/artist-store/security"
	"github.com/seokbyeongchae/artist-store/util"
)

type AuthController struct {
	store      *db.Store
	config     util.Config
	tokenMaker security.Maker
}

func New(store *db.Store, config util.Config, tokenMaker security.Maker) *AuthController {
	return &AuthController{store, config, tokenMaker}
}
