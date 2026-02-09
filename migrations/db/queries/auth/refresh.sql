--Refresh Token --------------------------------------------------------


-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens
    (id, user_id,session_id, token_hash, created_at, expires_at, revoked_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE session_id = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: GetActiveRefreshTokenBySession :one
SELECT *
FROM refresh_tokens
WHERE session_id = $1
  AND revoked_at IS NULL
  AND expires_at > NOW()
ORDER BY created_at DESC
LIMIT 1;

-- name: GetRefreshTokenByHash :one
SELECT *
FROM refresh_tokens
WHERE token_hash = $1;



-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW() 
WHERE token_hash = $1;


-- name: RotateRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW(), replaced_by = $2
WHERE id = $1;


-- name: DeleteRefreshTokenWithId :exec
DELETE FROM refresh_tokens WHERE id=$1;

-- name: DeleteRefreshTokensBySession :exec
DELETE FROM refresh_tokens
WHERE session_id = $1;
