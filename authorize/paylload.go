package authorize

import (
	"time"
	"errors"
	"github.com/google/uuid"
)


// Different types of error returned by the VerifyToken Function
var (
	ErrExpiredToken = errors.New("token has experied")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID    uuid.UUID `json:"id"`
	Owner string    `json:"owner"`
	Iat   time.Time `json:"issued_at"`
	Exp   time.Time `json:"expired_at"`
}


func (payload *Payload) Valid() error {
	if time.Now().After(payload.Exp) {
		return ErrExpiredToken
	}
	return nil
}


// implementing authorization