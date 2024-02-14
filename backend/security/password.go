package security

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
)

const (
	SALT_BYTES = 128
)

func GenerateSalt() []byte {
	salt := make([]byte, SALT_BYTES)
	if _, err := rand.Read(salt[:]); err != nil {
		panic(err)
	}

	return salt
}

func GeneratgeHashedPassword(salt []byte, password string) []byte {
	hasher := hmac.New(sha256.New, salt)
	hasher.Write([]byte(password))

	// return bytes.NewBuffer(hasher.Sum(nil)).String()
	return hasher.Sum(nil)
}
