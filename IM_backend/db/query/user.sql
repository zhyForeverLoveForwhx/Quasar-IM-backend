-- name: CreateUser :execresult
INSERT INTO users (username, password) VALUES (?,?);

-- name: GetUserByName :execresult
SELECT * FROM users WHERE username = ? LIMIT 1;