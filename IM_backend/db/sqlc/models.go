// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int32          `json:"id"`
	Username  sql.NullString `json:"username"`
	Password  sql.NullString `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
}
