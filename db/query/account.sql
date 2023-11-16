-- name: CreateAccount :one
INSERT INTO  accounts (uuid, email, hash_password) VALUES ($1, $2, $3) RETURNING *;

-- name: GetOneAccount :one
SELECT email, created_at, updated_at FROM accounts WHERE uuid = $1 AND  deleted_at IS NULL;

-- name: ListAccounts :many
SELECT email, created_at, updated_at FROM accounts WHERE email LIKE $1 AND deleted_at IS NULL ORDER BY created_at ASC OFFSET $2 LIMIT $3;

-- name: SoftDeleteAccount :exec
UPDATE accounts SET deleted_at = 'now()' WHERE uuid = $1;

-- name: ChangePassword :exec
UPDATE accounts SET hash_password = $2, updated_at = 'now()' WHERE uuid = $1 AND deleted_at IS NULL;