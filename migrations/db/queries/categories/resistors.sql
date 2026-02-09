-----------------------------------------------------------------------------------------------------------------------------
--Queries for Resistors
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetResistorByPartNo :one
SELECT *
FROM resistors
WHERE part_no=$1;

-- name: GetResistorsForDownload :many
SELECT * FROM resistors;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Resistors
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateResistor :one
INSERT INTO resistors
    (id, part_no, type, connector_type, voltage, notes)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateResistorByPartNo :one
UPDATE resistors AS r SET
  type           = COALESCE(sqlc.narg('type'), r.type),
  connector_type = COALESCE(sqlc.narg('connector_type'), r.connector_type),
  voltage        = COALESCE(sqlc.narg('voltage'), r.voltage),
  notes          = COALESCE(sqlc.narg('notes'), r.notes)
WHERE r.part_no=$1
RETURNING *;

-- name: DeleteResistorByPartNo :exec
DELETE FROM resistors
WHERE part_no=$1;
