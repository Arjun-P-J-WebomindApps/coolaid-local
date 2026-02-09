-----------------------------------------------------------------------------------------------------------------------------
--Queries for Rad Fan Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetRadFanAssyByPartNo :one
SELECT *
FROM rad_fan_assys
WHERE part_no=$1;

-- name: GetRadFanAssysForDownload :many
SELECT * FROM rad_fan_assys;


-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Rad Fan Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateRadFanAssy :one
INSERT INTO rad_fan_assys
    (id, part_no, voltage, motor_type, resistance, number_of_sockets, shroud, connector_type, fan_blade_diameter, number_of_blades, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateRadFanAssyByPartNo :one
UPDATE rad_fan_assys AS r SET
  voltage              = COALESCE(sqlc.narg('voltage'), r.voltage),
  motor_type           = COALESCE(sqlc.narg('motor_type'), r.motor_type),
  resistance           = COALESCE(sqlc.narg('resistance'), r.resistance),
  number_of_sockets    = COALESCE(sqlc.narg('number_of_sockets'), r.number_of_sockets),
  shroud               = COALESCE(sqlc.narg('shroud'), r.shroud),
  connector_type       = COALESCE(sqlc.narg('connector_type'), r.connector_type),
  fan_blade_diameter   = COALESCE(sqlc.narg('fan_blade_diameter'), r.fan_blade_diameter),
  number_of_blades     = COALESCE(sqlc.narg('number_of_blades'), r.number_of_blades),
  notes                = COALESCE(sqlc.narg('notes'), r.notes)
WHERE r.part_no=$1
RETURNING *;

-- name: DeleteRadFanAssyByPartNo :exec
DELETE FROM rad_fan_assys
WHERE part_no=$1;
