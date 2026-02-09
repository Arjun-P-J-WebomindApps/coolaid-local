-- Session----------------------------------------------------------

-- name: CreateUserSession :one
INSERT INTO user_sessions
    (session_id, user_id, created_at, expires_at, ip_address, user_agent)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetUserSessionByUserId :one
SELECT *
FROM user_sessions
WHERE user_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: GetUserSessionById :one
SELECT *
FROM user_sessions
WHERE  session_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: DeleteUserSession :exec 
DELETE FROM user_sessions
WHERE session_id=$1;



