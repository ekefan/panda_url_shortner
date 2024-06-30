package authorize

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtToken struct {
	key string
}

func CreateToken(username string, duration time.Duration, secretKey string) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", fmt.Errorf("couldn't generate payload id: %v", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenKey := jwtToken{
		key: secretKey,
	}
	tokenSigned, err := token.SignedString(tokenKey)
	return tokenSigned, nil
}

//token verification
