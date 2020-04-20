-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT pkey FROM users;

-- name: CreateUser :one
INSERT INTO users (
  pkey
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteUser :exec

DELETE FROM users
WHERE id = $1;