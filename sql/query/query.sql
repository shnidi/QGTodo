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

-- name: ParanoidListUsers :many
SELECT username
FROM users
WHERE deleted_at IS NULL;

-- name: CreateUser :one
INSERT INTO users (username, password, created_at, updated_at)
VALUES ( $1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET password=$1;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: ParanoidDeleteUser :exec
UPDATE users
SET deleted_at=$1;

-- name: CreateTask :one
INSERT INTO tasks (id, title, comment, created_at, updated_at)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: UpdateTaskTitle :exec
UPDATE tasks
SET title=$1;

-- name: ListTasksFromUser :many
SELECT * FROM tasks WHERE fk_user=$1;

-- name: ParanoidListTasksFromUser :many
SELECT * FROM tasks
WHERE deleted_at
IS NULL AND fk_user=$1;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = $1;

-- name: ParanoidDeleteTask :exec
UPDATE users SET deleted_at=$1;
