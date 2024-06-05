package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/ekefan/panda_url_shortner/util"
	"github.com/stretchr/testify/require"
)

/*
 ======================   POSSIBLE ERRORS =========================================
1. Creating new URL row with similar shortcode:UNIQUE constraint failed: urls.
 		short_code is the error when the unique code is broken
2.
*/


func randomLongURL() string {
	shortCode, err := util.RandomShortCode(5)
	if err != nil {
		return ""
	}
	longURL := fmt.Sprintf("https://github.com/testingMyApp/%s", shortCode)
	return longURL
}

//New Test Case: Check for when shortCode is an empty string
func createRandomURL(t *testing.T) (newURL URL){
	user:= createRandomUser(t)
	require.NotEmpty(t, user)
	shortCode, err := util.RandomShortCode(6)
	require.NoError(t, err)
	args := CreateURLArgs{
		Owner: user.ID,
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
	createRandomURL(t,) //A case when randomArgs returns "" shortCode
}

// func TestGetURL(t *testing.T) {
// 	urlRow := createRandomURL(t, randomArgs())
// 	args := GetURLArgs{ShortCode: urlRow.ShortCode}
// 	storedURL, err := ts.GetURL(args)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, storedURL)
	
// 	require.Equal(t, urlRow.ShortCode, storedURL.ShortCode)
// 	require.Equal(t, urlRow.LongURL, storedURL.LongURL)
// 	require.NotZero(t, urlRow.Owner)
// 	require.NotZero(t, urlRow.CreatedAt)
// }