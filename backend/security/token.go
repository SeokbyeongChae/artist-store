package security

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(accountId int64, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	Id        uuid.UUID `json:"id"`
	AccountId int64     `json:"account_id"`
	IssuedAt  time.Time `json:"issued_at"`
	jwt.RegisteredClaims
}

func NewPayload(accountId int64, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		tokenId,
		accountId,
		time.Now(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	return payload, nil
}
