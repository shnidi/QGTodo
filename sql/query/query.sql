-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: ListUsers :many
SELECT username
FROM users;

-- name: CreateUser :one
INSERT INTO users (username, password, created_at, updated_at)
VALUES (
        $1,
        $2,
        $3,
        $4
        )
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

--name: CreateTask :one
INSERT INTO tasks (
                   id, title, comment, created_at, updated_at
)
VALUES(
       $1,$2,$3,$4,$5
      )
RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks;

