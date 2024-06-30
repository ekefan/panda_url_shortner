package authorize

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Different types of error returned by Functions to verify Tokens and payload
var (
	ErrExpiredToken = errors.New("token has experied")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID    uuid.UUID        `json:"id"`
	Owner string           `json:"owner"`
	Iat   time.Time `json:"iat"`
	Exp   time.Time `json:"exp"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Exp) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(owner string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	iat := time.Now()
	exp := time.Now().Add(duration)
	payload := &Payload{
		ID:    tokenID,
		Owner: owner,
		Iat:   iat,
		Exp:   exp,
	}
	return payload, nil
}


// code below is not used in the applcation
func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {

	panic("not implemented")
}
func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	panic("not implemented")
}
func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	panic("not implemented")
}
func (p *Payload) GetIssuer() (string, error) {
	return p.Owner, nil
}
func (p *Payload) GetSubject() (string, error) {
	panic("not implemented")
}
func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	panic ("not implemented")
}
