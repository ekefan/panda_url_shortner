package database

import (
	"testing"
	"github.com/stretchr/testify/require"
)


var ts *Store
func createRandomURL(t *testing.T) (newURL URL){
	args := createURLArgs{
		shortCode: "newcode",
		longURL: "https://github.com/stretchr/testify",
	}
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
	createRandomURL(t)
}