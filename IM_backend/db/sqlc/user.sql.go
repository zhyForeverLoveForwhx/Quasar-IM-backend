// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (username, password) VALUES (?,?)
`

type CreateUserParams struct {
	Username sql.NullString `json:"username"`
	Password sql.NullString `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Username, arg.Password)
}
