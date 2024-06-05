package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ekefan/panda_url_shortner/util"
)


func randomEmail() (email string) {
	emailAddress, _:= util.RandomShortCode(4)
	email = fmt.Sprintf("%s@gmail.com", emailAddress)
	return
}

func createRandomUser(t *testing.T) {
	name, err := util.RandomShortCode(6)
	require.NoError(t, err)
	require.NotEmpty(t, name)
	hash, err:= util.RandomShortCode(8)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	args := CreateUserArgs{
		Name: name,
		Email: randomEmail(),
		HashedPassword: hash,
	}
	user, err := ts.CreateUser(args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.Name)
	require.NotEmpty(t, user.Password)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.Name, user.Name)
	require.Equal(t, args.HashedPassword, user.Password)
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}


