// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type User struct {
	ID             int32     `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
}
