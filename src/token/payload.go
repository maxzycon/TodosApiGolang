package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token invalid")
)

type Payload struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	IssueAt time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	uuid,err := uuid.NewRandom()
	if err != nil {
		return nil,err
	}

	payload := &Payload{
		ID: uuid,
		Username: username,
		IssueAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload,nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}