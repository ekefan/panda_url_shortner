package authorize

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtMaker struct {
	key string
}

const minSecretKeySize = 32

func NewJwtToken(secretKey string) (*JwtMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JwtMaker{
		key: secretKey,
	}, nil
}

func (t *JwtMaker) CreateToken(username string, duration time.Duration) (string, error) {
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

// token verification
func (t *JwtMaker) VerifyToken(token string) (*Payload, error) {
	// keyfunc checks the method used for signing the parsed
	// token in the header. If it is not the same then it is
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(t.key), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyfunc)
	if err != nil {
		// golang's implementation of jwt returns ErrInvalidClaims error on error
		return nil, err
	}
	// fmt.Println("got here")  test printing
	//get claims and return it
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

//Write test for tokens
// create a username, duration,
// create a new token, using a random key
// case one: Happy case, Verify the token,
// check the contents of the payload
