-----------------------------------------------------------------------------------------------------------------------------
--Queries for Pressure Switches
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetPressureSwitchByPartNo :one
SELECT *
FROM pressure_switches
WHERE part_no=$1;

-- name: GetPressureSwitchesForDownload :many
SELECT * FROM pressure_switches;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Pressure Switches
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreatePressureSwitch :one
INSERT INTO pressure_switches
    (id, part_no, connector_type, thread_type, notes)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdatePressureSwitchByPartNo :one
UPDATE pressure_switches AS ps SET
  connector_type = COALESCE(sqlc.narg('connector_type'), ps.connector_type),
  thread_type    = COALESCE(sqlc.narg('thread_type'), ps.thread_type),
  notes          = COALESCE(sqlc.narg('notes'), ps.notes)
WHERE ps.part_no=$1
RETURNING *;

-- name: DeletePressureSwitchByPartNo :exec
DELETE FROM pressure_switches
WHERE part_no=$1;
