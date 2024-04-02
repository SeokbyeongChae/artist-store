package constants

import "errors"

var (
	ErrNotFound            = errors.New("cannot find account")
	ErrUnauthorized        = errors.New("session is unauthorized")
	ErrHeaderIsNotProvided = errors.New("authorization header is not provided")
	ErrInvalidHeaderFormat = errors.New("invalid authorization header format")
)
