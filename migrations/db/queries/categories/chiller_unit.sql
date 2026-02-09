-----------------------------------------------------------------------------------------------------------------------------
--Queries for Chiller Units
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetChillerUnitByPartNo :one
SELECT *
FROM chiller_units
WHERE part_no=$1;

-- name: GetChillerUnitsForDownload :many
SELECT * FROM chiller_units;



-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Chiller Units
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateChillerUnit :one
INSERT INTO chiller_units
    (id, part_no, type, voltage, notes)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateChillerUnitByPartNo :one
UPDATE chiller_units AS c SET
  type     = COALESCE(sqlc.narg('type'), c.type),
  voltage  = COALESCE(sqlc.narg('voltage'), c.voltage),
  notes    = COALESCE(sqlc.narg('notes'), c.notes)
WHERE c.part_no=$1
RETURNING *;

-- name: DeleteChillerUnitByPartNo :exec
DELETE FROM chiller_units
WHERE part_no=$1;
