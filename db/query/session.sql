-- name: CreateSession :one
INSERT INTO sessions (uuid, refresh_token, user_agent, client_id, expired_at)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE uuid = $1;