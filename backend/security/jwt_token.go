package security

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

var (
	errExpiredToken = errors.New("token is expired")
	errInvalidToken = errors.New("token is invalid")
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(accountId int64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(accountId, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	switch {
	case jwtToken.Valid:
		payload, ok := jwtToken.Claims.(*Payload)
		if !ok {
			return nil, errInvalidToken
		}

		return payload, nil

	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		return nil, errExpiredToken
	// case errors.Is(err, jwt.ErrTokenMalformed):
	// 	fmt.Println("That's not even a token")
	// case errors.Is(err, jwt.ErrTokenSignatureInvalid):
	// 	// Invalid signature
	// 	fmt.Println("Invalid signature")
	default:
		return nil, errInvalidToken
	}
}
