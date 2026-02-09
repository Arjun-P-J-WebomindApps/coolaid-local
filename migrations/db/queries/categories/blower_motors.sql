-- -----------------------------------------------------------------------------------------------------------------------------
-- --Queries
-- -----------------------------------------------------------------------------------------------------------------------------

-- -------------used to get a blower motor by part number--#

-- name: GetBlowerMotorByPartNo :one
SELECT *
FROM blower_motors
WHERE part_no = $1;

-- -------------used to search blower motors by mounting type--#

-- name: GetBlowerMotorsByMounting :many
SELECT *
FROM blower_motors
WHERE mounting = $1;

-- -------------used to search blower motors by motor type--#

-- name: GetBlowerMotorsByMotorType :many
SELECT *
FROM blower_motors
WHERE motor_type = $1;

-- name: GetBlowerMotorsForDownload :many
SELECT * FROM blower_motors;

-- -----------------------------------------------------------------------------------------------------------------------------
-- --Mutations
-- -----------------------------------------------------------------------------------------------------------------------------

-- name: CreateBlowerMotor :one
INSERT INTO blower_motors (
    id, part_no, mounting, connector_type, impeller, resistance, motor_mounting, motor_type, voltage, notes
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateBlowerMotorByPartNo :one
UPDATE blower_motors AS b
SET
    mounting        = COALESCE(sqlc.narg('mounting'), b.mounting),
    connector_type  = COALESCE(sqlc.narg('connector_type'), b.connector_type),
    impeller        = COALESCE(sqlc.narg('impeller'), b.impeller),
    resistance      = COALESCE(sqlc.narg('resistance'), b.resistance),
    motor_mounting  = COALESCE(sqlc.narg('motor_mounting'), b.motor_mounting),
    motor_type      = COALESCE(sqlc.narg('motor_type'), b.motor_type),
    voltage         = COALESCE(sqlc.narg('voltage'), b.voltage),
    notes           = COALESCE(sqlc.narg('notes'), b.notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteBlowerMotorByPartNo :exec
DELETE FROM blower_motors
WHERE part_no = $1;
