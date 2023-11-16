package db

import (
	"context"
	"database/sql"
	"echo-simple-bank/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createUserTest(t *testing.T) User {
	account := createRandomAccount(t)
	arg := CreateUserParams{
		Uuid: account.Uuid,
		FirstName: sql.NullString{
			String: utils.RandomName(6),
			Valid:  true,
		},
		LastName: sql.NullString{
			String: utils.RandomName(6),
			Valid:  true,
		},
		Birthday: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Address: sql.NullString{
			String: utils.RandomName(6),
			Valid:  true,
		},
		Phone: sql.NullString{
			String: utils.RandomName(6),
			Valid:  true,
		},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Uuid, user.Uuid)
	require.Equal(t, arg.LastName.String, user.LastName.String)
	require.Equal(t, arg.FirstName.String, user.FirstName.String)
	require.WithinDuration(t, arg.Birthday.Time.UTC(), user.Birthday.Time.UTC(), 7*time.Hour)
	require.Equal(t, arg.Address.String, user.Address.String)
	require.Equal(t, arg.Phone.String, user.Phone.String)
	require.NotNil(t, user.CreatedAt)
	require.NotNil(t, user.UpdatedAt)
	require.NotZero(t, user.ID)

	return user
}

func TestQueries_CreateUser(t *testing.T) {
	createUserTest(t)
}

func TestQueries_GetOneUser(t *testing.T) {
	user1 := createUserTest(t)

	user2, err := testQueries.GetOneUser(context.Background(), user1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Address, user2.Address)
	require.Equal(t, user1.Phone, user2.Phone)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestQueries_UpdateUser(t *testing.T) {
	user1 := createUserTest(t)

	arg := UpdateUserParams{
		Uuid: user1.Uuid,
		FirstName: sql.NullString{
			String: utils.RandomName(10),
			Valid:  true,
		},
		LastName: sql.NullString{
			String: utils.RandomName(10),
			Valid:  true,
		},
		Birthday: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Address: sql.NullString{
			String: utils.RandomName(10),
			Valid:  true,
		},
		Phone: sql.NullString{
			String: utils.RandomName(10),
			Valid:  true,
		},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.NotEqual(t, user1.LastName, user2.LastName)
	require.NotEqual(t, user1.FirstName, user2.FirstName)
	require.NotEqual(t, user1.Birthday, user2.Birthday)
	require.NotEqual(t, user1.Address, user2.Address)
	require.NotEqual(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.NotEqual(t, user1.UpdatedAt, user2.UpdatedAt)
}

func TestQueries_SoftDeleteUser(t *testing.T) {
	user := createUserTest(t)

	err := testQueries.SoftDeleteUser(context.Background(), user.Uuid)
	require.NoError(t, err)

	userDel, err := testQueries.GetOneUser(context.Background(), user.Uuid)
	require.Error(t, err)
	require.Empty(t, userDel)
}

func TestQueries_ListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createUserTest(t)
	}

	arg := ListUsersParams{
		FirstName: sql.NullString{
			String: "%%",
			Valid:  true,
		},
		Offset: 5,
		Limit:  5,
	}

	accounts, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, int32(len(accounts)), arg.Limit)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
