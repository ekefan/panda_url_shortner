package authorize

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtToken struct {
	key string
}

const minSecretKeySize = 32

func NewJwtToken(secretKey string)(*jwtToken, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &jwtToken{
		key: secretKey,
	}, nil
}

func (t *jwtToken)CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", fmt.Errorf("couldn't generate payload id: %v", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenKey := t.key
	tokenSigned, err := token.SignedString([]byte(tokenKey))
	if err != nil {
		return "", err
	}
	return tokenSigned, nil
}

//token verification
// func (t *jwtTOken) VerifyToken(f)(*Payload, error)