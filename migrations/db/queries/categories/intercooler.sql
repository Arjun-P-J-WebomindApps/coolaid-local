-----------------------------------------------------------------------------------------------------------------------------
--Queries for Intercoolers
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetIntercoolerByPartNo :one
SELECT *
FROM intercoolers
WHERE part_no=$1;

-- name: GetIntercoolersForDownload :many
SELECT * FROM intercoolers;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Intercoolers
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateIntercooler :one
INSERT INTO intercoolers
    (id, part_no, size, temp_sensor, notes)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateIntercoolerByPartNo :one
UPDATE intercoolers AS ic SET
  size        = COALESCE(sqlc.narg('size'), ic.size),
  temp_sensor = COALESCE(sqlc.narg('temp_sensor'), ic.temp_sensor),
  notes       = COALESCE(sqlc.narg('notes'), ic.notes)
WHERE ic.part_no=$1
RETURNING *;

-- name: DeleteIntercoolerByPartNo :exec
DELETE FROM intercoolers
WHERE part_no=$1;
