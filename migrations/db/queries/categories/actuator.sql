-----------------------------------------------------------------------------------------------------------------------------
--Queries
-----------------------------------------------------------------------------------------------------------------------------

-------------used to get an actuator by part number--#

-- name: GetActuatorByPartNo :one
SELECT *
FROM actuators
WHERE part_no = $1;

-- name: GetActuatorForDownload :many
SELECT * FROM actuators;

------------------------------------------------------------------------------------------------------------------------------
--Mutation
------------------------------------------------------------------------------------------------------------------------------

-- name: CreateActuator :one
INSERT INTO actuators (
    id, part_no, connector_type, mounting, voltage, rotation_angle, notes
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateActuatorByPartNo :one
UPDATE actuators AS a
SET
    connector_type = COALESCE(sqlc.narg('connector_type'), a.connector_type),
    mounting       = COALESCE(sqlc.narg('mounting'), a.mounting),
    voltage        = COALESCE(sqlc.narg('voltage'), a.voltage),
    rotation_angle = COALESCE(sqlc.narg('rotation_angle'), a.rotation_angle),
    notes          = COALESCE(sqlc.narg('notes'), a.notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteActuatorByPartNo :exec
DELETE FROM actuators
WHERE part_no = $1;
