
-- User---------------------------------------------------------------

----------------------------------------------------------------------
-- Queries
---------------------------------------------------------------------


-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id=$1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email=$1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username=$1;

---------------------------------------------------------------------------------------------------------
-- Mutation
---------------------------------------------------------------------------------------------------------


-- name: CreateUser :one
INSERT INTO users
    (id,name,username,email,password,mobile,role,is_active,max_sessions,created_at,updated_at,deleted_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;


-- name: UpdateUser :exec
UPDATE users
SET
    name         = COALESCE(sqlc.narg(name), name),
    password     = COALESCE(sqlc.narg(password), password),
    mobile       = COALESCE(sqlc.narg(mobile), mobile),
    role         = COALESCE(sqlc.narg(role), role),
    is_active    = COALESCE(sqlc.narg(is_active), is_active),
    max_sessions = COALESCE(sqlc.narg(max_sessions), max_sessions),
    updated_at   = NOW()
WHERE id = $1;



