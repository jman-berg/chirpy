-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at)
VALUES (
	$1,
	NOW(),
	NOW(),
	$2,
	NOW() + INTERVAL '60 day'
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT
user_id
FROM refresh_tokens
WHERE token = $1 AND revoked_at is null AND expires_at > NOW();

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW(), updated_at = NOW()
WHERE $1 = token;

