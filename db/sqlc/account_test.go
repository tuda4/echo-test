package db

import (
	"context"
	"echo-simple-bank/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Uuid:         uuid.NewString(),
		Email:        utils.RandomEmail(10),
		HashPassword: utils.RandomHashPassword(16),
	}

	accountTest, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountTest)

	require.NotZero(t, accountTest.ID)
	require.Equal(t, arg.Uuid, accountTest.Uuid)
	require.Equal(t, arg.Email, accountTest.Email)
	require.Equal(t, arg.HashPassword, accountTest.HashPassword)
	require.NotNil(t, accountTest.CreatedAt)
	require.NotNil(t, accountTest.UpdatedAt)

	return accountTest
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetOneAccount(t *testing.T) {
	account := createRandomAccount(t)

	accountGet, err := testQueries.GetOneAccount(context.Background(), account.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, accountGet)
	require.Equal(t, accountGet.Email, account.Email)
	require.Equal(t, accountGet.CreatedAt, account.CreatedAt)
	require.Equal(t, accountGet.UpdatedAt, account.UpdatedAt)
}

func TestChangePassword(t *testing.T) {
	account := createRandomAccount(t)
	arg := ChangePasswordParams{
		Uuid:         account.Uuid,
		HashPassword: utils.RandomHashPassword(16),
	}

	err := testQueries.ChangePassword(context.Background(), arg)
	require.NoError(t, err)
}

func TestQueries_SoftDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.SoftDeleteAccount(context.Background(), account.Uuid)
	require.NoError(t, err)

	accountGet, err := testQueries.GetOneAccount(context.Background(), account.Uuid)
	require.Error(t, err)
	require.Empty(t, accountGet)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Email:  "%%",
		Offset: 5,
		Limit:  5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(accounts), 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
