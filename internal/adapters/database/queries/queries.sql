-- name: ListUsers :many
SELECT *
FROM users;

-- name: FindUserById :one
SELECT *
FROM users
Where id = $1;

-- name: ListTasks :many
SELECT *
FROM tasks;
