package dbtest

import (
	db "github.com/ekefan/panda_url_shortner/database"
	"fmt"
	"testing"
	"time"

	"github.com/ekefan/panda_url_shortner/util"
	"github.com/stretchr/testify/require"

)

func randomLongURL() string {
	shortCode, err := util.RandomShortCode(5)
	if err != nil {
		return ""
	}
	longURL := fmt.Sprintf("https://github.com/testingMyApp/%s", shortCode)
	return longURL
}

//New Test Case: Check for when shortCode is an empty string
func createRandomURL(t *testing.T) (newURL db.URL){
	user:= createRandomUser(t)
	require.NotEmpty(t, user)
	shortCode, err := util.RandomShortCode(6)
	require.NoError(t, err)
	args := db.CreateURLArgs{
		Owner: user.Name,
		ShortCode: shortCode,
		LongURL: randomLongURL(),
	}
	newURL, err = ts.CreateURL(args)
	require.NoError(t, err)
	require.NotEmpty(t, newURL)
	require.Equal(t, newURL.ShortCode, args.ShortCode)
	require.Equal(t, newURL.LongURL, args.LongURL)
	require.NotZero(t, newURL.CreatedAt)
	require.Zero(t, newURL.UpdatedAt)
	timeStamp := time.Time(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	require.Equal(t, newURL.UpdatedAt, timeStamp)
	require.NotZero(t, newURL.Owner)
	require.NotZero(t, newURL.CreatedAt)
	return
}

func TestCreateURL(t *testing.T) {
	createRandomURL(t) //A case when randomArgs returns "" shortCo
	// fmt.Println("nw")
}