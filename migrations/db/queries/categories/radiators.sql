-----------------------------------------------------------------------------------------------------------------------------
--Queries for Radiators
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetRadiatorByPartNo :one
SELECT *
FROM radiators
WHERE part_no=$1;

-- name: GetRadiatorsForDownload :many
SELECT * FROM radiators;


-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Radiators
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateRadiator :one
INSERT INTO radiators
    (id, part_no, size, transmission, temp_sensor, tank, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateRadiatorByPartNo :one
UPDATE radiators AS r SET
  size         = COALESCE(sqlc.narg('size'), r.size),
  transmission = COALESCE(sqlc.narg('transmission'), r.transmission),
  temp_sensor  = COALESCE(sqlc.narg('temp_sensor'), r.temp_sensor),
  tank         = COALESCE(sqlc.narg('tank'), r.tank),
  notes        = COALESCE(sqlc.narg('notes'), r.notes)
WHERE r.part_no=$1
RETURNING *;

-- name: DeleteRadiatorByPartNo :exec
DELETE FROM radiators
WHERE part_no=$1;
