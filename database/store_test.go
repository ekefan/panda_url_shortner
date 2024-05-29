package database

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)


var ts Store

/*
 ======================   POSSIBLE ERRORS =========================================
1. Creating new URL row with similar shortcode:UNIQUE constraint failed: urls.
 		short_code is the error when the unique code is broken
2. 
*/
func randomShortCode() string{
	letters := "abcdefghijklmnopqrstuvwxyz"
	var shortCode strings.Builder
	for i := 0; i < 5; i++{
		idx := rand.Intn(len(letters))
		shortCode.WriteByte(letters[idx])
	}
	return shortCode.String()
}

func randomLongURL() string {
	return fmt.Sprintf("https://github.com/testingMyApp/%s", randomShortCode())
}

func randomArgs() createURLArgs {
	return createURLArgs{
		shortCode: randomShortCode(),
		longURL: randomLongURL(),
	}
}


func createRandomURL(t *testing.T, args createURLArgs) (newURL URL){
	newURL, err := ts.CreateURL(args)
	require.NoError(t, err)
	require.NotEmpty(t, newURL)
	require.Equal(t, newURL.ShortCode, args.shortCode)
	require.Equal(t, newURL.LongURL, args.longURL)
	require.NotZero(t, newURL.ID)
	require.NotZero(t, newURL.CreatedAt)
	return
}

func TestCreateURL(t *testing.T) {
	createRandomURL(t, randomArgs())
}

func TestGetURL(t *testing.T) {
	urlRow := createRandomURL(t, randomArgs())
	args := getURLArgs{shortCode: urlRow.ShortCode}
	storedURL, err := ts.GetURL(args)
	require.NoError(t, err)
	require.NotEmpty(t, storedURL)
	
	require.Equal(t, urlRow.ShortCode, storedURL.ShortCode)
	require.Equal(t, urlRow.LongURL, storedURL.LongURL)
	require.NotZero(t, urlRow.ID)
	require.NotZero(t, urlRow.CreatedAt)
}