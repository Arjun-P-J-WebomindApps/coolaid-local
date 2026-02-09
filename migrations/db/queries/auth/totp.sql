-- TOTP -------------------------------------------------------------

---------------------------------------------------------------------
--Queries
---------------------------------------------------------------------

-- name: GetUserTotpByUserID :one
SELECT *
FROM user_totp
WHERE user_id = $1;

-- name: EnableUserTotp :exec
UPDATE user_totp
SET
  is_enabled = TRUE,
  updated_at = NOW()
WHERE user_id = $1;


----------------------------------------------------------------------
-------Mutation
----------------------------------------------------------------------


-- name: CreateUserTotp :one
INSERT INTO user_totp (user_id, secret, is_enabled)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpsertUserTotp :one
INSERT INTO user_totp (user_id, secret, is_enabled)
VALUES ($1, $2, $3)
ON CONFLICT (user_id) DO UPDATE
SET
  secret     = EXCLUDED.secret,
  is_enabled = EXCLUDED.is_enabled,
  updated_at = NOW()
RETURNING *;
