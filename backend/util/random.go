package util

import (
	"math/rand"
	"strings"
	"time"
)

const characterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(length int) string {
	var stringBuilder strings.Builder
	k := len(characterRunes)

	for i := 0; i < length; i++ {
		c := characterRunes[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}

	return stringBuilder.String()
}
