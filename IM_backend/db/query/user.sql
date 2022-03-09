-- name: CreateUser :execresult
INSERT INTO users (username, password) VALUES (?,?);

-- name: GetUserByName :one
SELECT * FROM users WHERE username = ? LIMIT 1;