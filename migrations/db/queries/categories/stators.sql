-----------------------------------------------------------------------------------------------------------------------------
--Queries for Stators
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetStatorByPartNo :one
SELECT *
FROM stators
WHERE part_no=$1;

-- name: GetStatorsForDownload :many
SELECT * FROM stators;
-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Stators
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateStator :one
INSERT INTO stators
    (id, part_no, voltage, compressor_details, notes)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateStatorByPartNo :one
UPDATE stators AS s SET
  voltage             = COALESCE(sqlc.narg('voltage'), s.voltage),
  compressor_details  = COALESCE(sqlc.narg('compressor_details'), s.compressor_details),
  notes               = COALESCE(sqlc.narg('notes'), s.notes)
WHERE s.part_no=$1
RETURNING *;

-- name: DeleteStatorByPartNo :exec
DELETE FROM stators
WHERE part_no=$1;
