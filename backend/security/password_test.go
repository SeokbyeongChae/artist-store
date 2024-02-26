package security

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	randomSalt, err := GenerateSalt()
	require.NoError(t, err)

	const password = "testPassword"
	hashedPassword := HashPassword(randomSalt, password)
	require.Equal(t, true, CheckPassword(hashedPassword, randomSalt, password))

	const wrongSssword = "wrongPassword"
	hashedWrongPassword := HashPassword(randomSalt, wrongSssword)
	require.Equal(t, false, CheckPassword(hashedWrongPassword, randomSalt, password))
}
