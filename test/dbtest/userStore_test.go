package dbtest



import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ekefan/panda_url_shortner/util"
	db "github.com/ekefan/panda_url_shortner/database"
)


func randomEmail() (email string) {
	emailAddress, _:= util.RandomShortCode(4)
	email = fmt.Sprintf("%s@gmail.com", emailAddress)
	return
}

func createRandomUser(t *testing.T) db.USER {
	name, err := util.RandomShortCode(6)
	require.NoError(t, err)
	require.NotEmpty(t, name)
	hash, err:= util.RandomShortCode(8)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	args := db.CreateUserArgs{
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