package test

import (
	"testing"

	"github.com/ekefan/panda_url_shortner/util"
	"github.com/stretchr/testify/require"
)


func TestGenerateHash(t *testing.T){
	password, err := util.RandomShortCode(8)
	require.NoError(t, err)
	require.NotEmpty(t, password)

	hash, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash)



	err = util.VerifyPassword(hash, password)
	require.NoError(t, err)
}