package database

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ekefan/panda_url_shortner/util"
)


var ts Store

/*
 ======================   POSSIBLE ERRORS =========================================
1. Creating new URL row with similar shortcode:UNIQUE constraint failed: urls.
 		short_code is the error when the unique code is broken
2. 
*/


func randomLongURL() string {
	shortCode, err := util.RandomShortCode()
	if err != nil {
		return ""
	}
	longURL := fmt.Sprintf("https://github.com/testingMyApp/%s", shortCode)
	return longURL
}

func randomArgs() CreateURLArgs {
	shortCode, _ := util.RandomShortCode()
	// Error here... this function shoul not return createURLARgs if err shortCode is ""
	return CreateURLArgs{
		ShortCode: shortCode,
		LongURL: randomLongURL(),
	}
}

//New Test Case: Check for when shortCode is an empty string


func createRandomURL(t *testing.T, args CreateURLArgs) (newURL URL){
	newURL, err := ts.CreateURL(args)
	require.NoError(t, err)
	require.NotEmpty(t, newURL)
	require.Equal(t, newURL.ShortCode, args.ShortCode)
	require.Equal(t, newURL.LongURL, args.LongURL)
	require.NotZero(t, newURL.ID)
	require.NotZero(t, newURL.CreatedAt)
	return
}

func TestCreateURL(t *testing.T) {
	createRandomURL(t, randomArgs()) //A case when randomArgs returns "" shortCode
}

func TestGetURL(t *testing.T) {
	urlRow := createRandomURL(t, randomArgs())
	args := GetURLArgs{ShortCode: urlRow.ShortCode}
	storedURL, err := ts.GetURL(args)
	require.NoError(t, err)
	require.NotEmpty(t, storedURL)
	
	require.Equal(t, urlRow.ShortCode, storedURL.ShortCode)
	require.Equal(t, urlRow.LongURL, storedURL.LongURL)
	require.NotZero(t, urlRow.ID)
	require.NotZero(t, urlRow.CreatedAt)
}