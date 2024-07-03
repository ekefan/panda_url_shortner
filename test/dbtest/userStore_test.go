package dbtest

import (
	"fmt"
	"testing"

	db "github.com/ekefan/panda_url_shortner/database"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/stretchr/testify/require"
)

func randomEmail() (email string) {
	emailAddress, _ := util.RandomShortCode(4)
	email = fmt.Sprintf("%s@gmail.com", emailAddress)
	return
}

func createRandomUser(t *testing.T) db.USER {
	name, err := util.RandomShortCode(6)
	require.NoError(t, err)
	require.NotEmpty(t, name)
	hash, err := util.RandomShortCode(8)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	args := db.CreateUserArgs{
		Name:           name,
		Email:          randomEmail(),
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
	// fmt.Println(args.HashedPassword, user.Password)
	require.Equal(t, args.HashedPassword, user.Password)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t) //correctUSer

}

func TestGetUser(t *testing.T) {
	usr := createRandomUser(t)

	args := db.GetUserArgs{
		Name: usr.Name,
	}
	user, err := ts.GetUser(args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, usr.Name, user.Name)
	require.Equal(t, usr.Email, user.Email)
	require.Equal(t, usr.Password, user.Password)
	require.Equal(t, usr.CreatedAt, user.CreatedAt)
}

func TestUpdateUser(t *testing.T) {
	usr := createRandomUser(t)
	url := createRandomURL(t, usr)
	name, err := util.RandomShortCode(4)
	require.NoError(t, err)
	require.NotEmpty(t, name)
	require.Equal(t, url.Owner, usr.Name)
	args := db.TxUserArgs{
		UserID:     usr.ID,
		UserName:   usr.Name,
		NameUpdate: name,
	}

	updatedUser, err := ts.TxUpdateUser(args)
	require.NoError(t, err)
	updatedUrl, err := ts.GetURL(db.GetURLArgs{
		ShortCode: url.ShortCode,
	})
	require.NotEmpty(t, updatedUrl)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.Equal(t, updatedUser.Name, name)
	require.Equal(t, updatedUser.Name, updatedUrl.Owner)
}

func TestDeleteUser(t *testing.T) {
	usr := createRandomUser(t)
	urls := []db.URL{}
	for i := 0; i < 5; i++ {
		urls = append(urls, createRandomURL(t, usr))
	}
	err := ts.TxDeleteUser(db.TxUserArgs{
		UserID:   usr.ID,
		UserName: usr.Name,
	})
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		_, err = ts.GetURL(db.GetURLArgs{
			ShortCode: urls[i].ShortCode,
		})
		require.Error(t, err)
	}
}
