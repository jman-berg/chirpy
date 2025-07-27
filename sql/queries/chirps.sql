-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(),
	$1,
	$2

)
RETURNING *;

-- name: GetChirps :many
SELECT * FROM chirps
WHERE (sqlc.arg(uuid_filter)::uuid = '00000000-0000-0000-0000-000000000000' OR user_id = sqlc.arg(uuid_filter)::uuid)
ORDER BY created_at ASC;


-- name: GetChirp :one
SELECT * FROM chirps
WHERE user_id = $1
LIMIT (1);

-- name: DeleteChirp :exec
DELETE FROM chirps WHERE id = $1 AND user_id = $2;
