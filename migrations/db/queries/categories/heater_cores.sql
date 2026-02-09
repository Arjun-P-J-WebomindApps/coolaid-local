-----------------------------------------------------------------------------------------------------------------------------
--Queries for Heater Cores
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetHeaterCoreByPartNo :one
SELECT *
FROM heater_cores
WHERE part_no=$1;


-- name: GetHeaterCoresForDownload :many
SELECT * FROM heater_cores;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Heater Cores
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateHeaterCore :one
INSERT INTO heater_cores
    (id, part_no, size, pipe, type, notes)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateHeaterCoreByPartNo :one
UPDATE heater_cores AS hc SET
  size     = COALESCE(sqlc.narg('size'), hc.size),
  pipe     = COALESCE(sqlc.narg('pipe'), hc.pipe),
  type     = COALESCE(sqlc.narg('type'), hc.type),
  notes    = COALESCE(sqlc.narg('notes'), hc.notes)
WHERE hc.part_no=$1
RETURNING *;

-- name: DeleteHeaterCoreByPartNo :exec
DELETE FROM heater_cores
WHERE part_no=$1;
