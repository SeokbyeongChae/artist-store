package db

import (
	"context"
	"testing"

	"github.com/seokbyeongchae/artist-store/security"
	"github.com/seokbyeongchae/artist-store/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) int64 {
	randomSalt, err := security.GenerateSalt()
	require.NoError(t, err)

	randomEmail := util.RandomString(12) + "@" + util.RandomString(4) + "." + util.RandomString(4)

	arg := CreateAccountParams{
		Salt:     randomSalt,
		Email:    randomEmail,
		Password: security.HashPassword(randomSalt, "password"),
	}

	result, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	affectedRowCount, err := result.RowsAffected()
	require.NoError(t, err)

	require.Equal(t, int64(1), affectedRowCount)

	insertId, err := result.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, insertId)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	return insertId
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountId := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), accountId)

	require.NoError(t, err)
	require.Equal(t, accountId, account.ID)
}
