-- name: CreateUser :one
INSERT INTO users (uuid, first_name, last_name, birthday, address, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetOneUser :one
SELECT first_name, last_name, birthday, address, phone, created_at FROM users WHERE uuid = $1 AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT first_name, last_name, birthday, address, phone, created_at
FROM users
WHERE first_name LIKE $1
OR last_name LIKE $1
AND deleted_at IS NULL
ORDER BY created_at ASC
OFFSET $2
LIMIT $3;

-- name: UpdateUser :one
UPDATE users SET first_name = $2, last_name = $3, birthday = $4, address = $5, phone = $6, updated_at = 'now()'
WHERE uuid = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteUser :exec
UPDATE users SET  deleted_at = 'now()' WHERE uuid = $1;