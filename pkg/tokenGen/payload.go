package tokengen

import (
	"errors"
	"time"

	"github.com/BON4/gosubs/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Instance  domain.Account
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(account *domain.Account, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Instance:  *account,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func GetAccountFromContext(ctx *gin.Context, payloadkey string) (*Payload, error) {
	payload, ok := ctx.Get(payloadkey)
	if !ok {

		return nil, errors.New("TODO: custom error 1")
	}

	payloadParsed, ok := payload.(*Payload)
	if !ok {
		return nil, errors.New("TODO: custom error 1")

	}
	return payloadParsed, nil
}
