package dbtest

import (
	"fmt"
	"testing"
	"time"

	db "github.com/ekefan/panda_url_shortner/database"

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

// New Test Case: Check for when shortCode is an empty string
func createRandomURL(t *testing.T, user db.USER) (newURL db.URL) {
	require.NotEmpty(t, user)
	shortCode, err := util.RandomShortCode(6)
	require.NoError(t, err)
	args := db.CreateURLArgs{
		Owner:     user.Name,
		ShortCode: shortCode,
		LongURL:   randomLongURL(),
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
	user := createRandomUser(t)
	createRandomURL(t, user) //A case when randomArgs returns "" shortCo

}

func TestGetUrls(t *testing.T) {
	// fmt.Println("nw")
	urls := []db.URL{}
	user := createRandomUser(t)
	for i := 0; i < 10; i++ {
		urls = append(urls, createRandomURL(t, user))
	}
	args := db.GetURLsArg{
		Owner:  user.Name,
		Limit:  10,
		Offset: 5,
	}
	dbUrls, err := ts.GetURLs(args)
	require.NoError(t, err)
	require.NotEmpty(t, dbUrls)
	require.Equal(t, len(dbUrls), 5)
	for i, url := range urls{
		require.Equal(t, url.Owner, urls[i].Owner)
	}
}

func TestTxUpdateShortCode(t *testing.T) {
	user := createRandomUser(t)
	url := createRandomURL(t, user)

	require.Equal(t, user.Name, url.Owner)
	shortCode, err := util.RandomShortCode(5)
	require.NoError(t, err)
	args := db.TxUrlArgs{
		Owner: url.Owner,
		// CurrentShortCode: url.ShortCode,
		ShortCode: shortCode,
	}
	updatedUrl, err := ts.TxUpdateShortCode(args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUrl)
	require.Equal(t, args.ShortCode, updatedUrl.ShortCode)
	require.NotEqual(t, updatedUrl.ShortCode, url.ShortCode)
	require.Equal(t, args.ShortCode, updatedUrl.ShortCode)
	require.Equal(t, url.Owner, updatedUrl.Owner)
	
}

//case for already existing shortcode in the database


func  TestTxDeleteUrl(t *testing.T) {
	usr := createRandomUser(t)

	usersUrls := []db.URL{}
	for i := 0; i < 5; i++ {
		usersUrls = append(usersUrls, createRandomURL(t, usr))
	}

	for i := 0; i < 5; i++ {
		arg := db.TxUrlArgs{
			Owner: usr.Name,
			ShortCode: usersUrls[i].ShortCode,
		}
		err := ts.TxDeleteUrl(arg)
		require.NoError(t, err)

		// check that the urls don't exist in the database anymore
		_, err = ts.GetURL(db.GetURLArgs{
			ShortCode: usersUrls[i].ShortCode,
		})
		require.Error(t, err)
	}
	


}