package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// By adding field to Payload also change NewPayload & GetMap & NewPaylodfromMap
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) GetMap() map[string]interface{} {
	m := map[string]interface{}{
		"id":         p.ID,
		"username":   p.Username,
		"issued_at":  p.IssuedAt,
		"expired_at": p.ExpiredAt,
	}
	return m
}

func PayloadFromMap(m map[string]interface{}) (*Payload, error) {
	idStr := m["id"].(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	username := m["username"].(string)
	issuedAtStr := m["issued_at"].(string)
	issuedAt, err := time.Parse(time.RFC3339Nano, issuedAtStr)
	if err != nil {
		return nil, err
	}

	expiredAtStr := m["expired_at"].(string)
	expiredAt, err := time.Parse(time.RFC3339Nano, expiredAtStr)
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        id,
		Username:  username,
		ExpiredAt: expiredAt,
		IssuedAt:  issuedAt,
	}
	return payload, nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

type ValidError error
type ExpiredError error

var (
	ErrExpiredToken ExpiredError = errors.New("token has expired")
	ErrInvalidToken ValidError   = errors.New("token is invalid")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		fmt.Println("ОШИБКААААААААААААААААААААААААА ТУТТТТТТТТТТТ")
		return ErrExpiredToken
	}
	return nil
}
