-----------------------------------------------------------------------------------------------------------------------------
--Queries for Cond Fan Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetCondFanAssyByPartNo :one
SELECT *
FROM cond_fan_assys
WHERE part_no=$1;


-- name: GetCondFanAssysForDownload :many
SELECT * FROM cond_fan_assys;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Cond Fan Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateCondFanAssy :one
INSERT INTO cond_fan_assys
    (id, part_no, voltage, motor_type, resistance, fan_blade_diameter, number_of_blades, shroud, connector_type, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateCondFanAssyByPartNo :one
UPDATE cond_fan_assys AS c SET
  voltage             = COALESCE(sqlc.narg('voltage'), c.voltage),
  motor_type          = COALESCE(sqlc.narg('motor_type'), c.motor_type),
  resistance          = COALESCE(sqlc.narg('resistance'), c.resistance),
  fan_blade_diameter  = COALESCE(sqlc.narg('fan_blade_diameter'), c.fan_blade_diameter),
  number_of_blades    = COALESCE(sqlc.narg('number_of_blades'), c.number_of_blades),
  shroud              = COALESCE(sqlc.narg('shroud'), c.shroud),
  connector_type      = COALESCE(sqlc.narg('connector_type'), c.connector_type),
  notes               = COALESCE(sqlc.narg('notes'), c.notes)
WHERE c.part_no=$1
RETURNING *;

-- name: DeleteCondFanAssyByPartNo :exec
DELETE FROM cond_fan_assys
WHERE part_no=$1;
