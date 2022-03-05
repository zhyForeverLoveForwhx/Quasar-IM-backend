package token

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// Payload contains the payload data of the token
//ID to 防止token被泄露
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.New() //分配各自ID
	if err != nil {
		return nil, err
	}

	Payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return Payload, nil
}
