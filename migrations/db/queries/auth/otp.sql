-- OTP----------------------------------------------------------------

-- name: CreateOTP :one 
INSERT INTO user_otps
    (user_id,otp_code,expires_at,is_used,created_at)
-- id is created internally
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetLatestOTPFromUser :one
SELECT *
FROM user_otps
WHERE user_id=$1
ORDER BY created_at DESC LIMIT 1;

-- name: DeleteUserOTPByUserId :exec
DELETE FROM user_otps
WHERE user_id=$1;

-- name: MarkOTPAsUsed :exec
UPDATE user_otps SET is_used = TRUE WHERE user_id=$1;
