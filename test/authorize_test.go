package test

import (
	"testing"
	"time"

	"github.com/ekefan/panda_url_shortner/authorize"
	"github.com/ekefan/panda_url_shortner/util"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	key := "12345678901234567890123456789012"
	randomUser, err := util.RandomShortCode(4)
	require.NoError(t, err)
	require.NotEmpty(t, randomUser)
	duration := 3 * time.Minute

	//Case One: Happy case token is created token is verified
	tokenMaker, err := authorize.NewJwtToken(key)
	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)

	token, err := tokenMaker.CreateToken(randomUser, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := tokenMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, payload.ID)
	require.Equal(t, payload.Owner, randomUser)
	require.WithinDuration(t, payload.Iat, time.Now(), time.Second)

}