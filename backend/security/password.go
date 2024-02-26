package security

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"github.com/google/uuid"
)

const (
	SALT_BYTES = 128
)

func GenerateSalt() ([]byte, error) {
	randBytes := make([]byte, SALT_BYTES)
	if _, err := rand.Read(randBytes[:]); err != nil {
		return nil, err
	}

	randUuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	hasher := hmac.New(sha256.New, randBytes)
	hasher.Write([]byte(randUuid.String()))

	return hasher.Sum(nil), nil
}

func HashPassword(salt []byte, password string) []byte {
	hasher := hmac.New(sha256.New, salt)
	hasher.Write([]byte(password))
	return hasher.Sum(nil)
}

func CheckPassword(hashedPassword []byte, salt []byte, password string) bool {
	return bytes.Compare(hashedPassword, HashPassword(salt, password)) == 0
}
