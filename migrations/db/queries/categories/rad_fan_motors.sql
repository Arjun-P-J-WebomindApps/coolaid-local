-----------------------------------------------------------------------------------------------------------------------------
--Queries for Rad Fan Motors
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetRadFanMotorByPartNo :one
SELECT *
FROM rad_fan_motors
WHERE part_no=$1;


-- name: GetRadFanMotorsForDownload :many
SELECT * FROM rad_fan_motors;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Rad Fan Motors
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateRadFanMotor :one
INSERT INTO rad_fan_motors
    (id, part_no, fan_blade_diameter, number_of_blades, voltage, number_of_sockets, connector_type, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateRadFanMotorByPartNo :one
UPDATE rad_fan_motors AS r SET
  fan_blade_diameter  = COALESCE(sqlc.narg('fan_blade_diameter'), r.fan_blade_diameter),
  number_of_blades    = COALESCE(sqlc.narg('number_of_blades'), r.number_of_blades),
  voltage             = COALESCE(sqlc.narg('voltage'), r.voltage),
  number_of_sockets   = COALESCE(sqlc.narg('number_of_sockets'), r.number_of_sockets),
  connector_type      = COALESCE(sqlc.narg('connector_type'), r.connector_type),
  notes               = COALESCE(sqlc.narg('notes'), r.notes)
WHERE r.part_no=$1
RETURNING *;

-- name: DeleteRadFanMotorByPartNo :exec
DELETE FROM rad_fan_motors
WHERE part_no=$1;
