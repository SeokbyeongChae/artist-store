package db

import (
	"context"
	"testing"

	"github.com/seokbyeongchae/artist-store/security"
	"github.com/seokbyeongchae/artist-store/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	randomSalt := security.GenerateSalt()
	randomEmail := util.RandomString(12) + "@" + util.RandomString(4) + "." + util.RandomString(4)

	arg := CreateAccountParams{
		Salt:     randomSalt,
		Email:    randomEmail,
		Password: security.GeneratgeHashedPassword(randomSalt, "password"),
	}

	result, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
