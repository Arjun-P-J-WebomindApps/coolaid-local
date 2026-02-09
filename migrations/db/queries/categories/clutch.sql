-----------------------------------------------------------------------------------------------------------------------------
--Queries for Clutch Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetClutchAssyByPartNo :one
SELECT *
FROM clutch_assys
WHERE part_no=$1;

-- name: GetClutchAssysForDownload :many
SELECT * FROM clutch_assys;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Clutch Assys
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateClutchAssy :one
INSERT INTO clutch_assys
    (id, part_no, pulley_ribs, pulley_size, compressor_details, connector_type, voltage, shaft_type, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateClutchAssyByPartNo :one
UPDATE clutch_assys AS c SET
  pulley_ribs         = COALESCE(sqlc.narg('pulley_ribs'), c.pulley_ribs),
  pulley_size         = COALESCE(sqlc.narg('pulley_size'), c.pulley_size),
  compressor_details  = COALESCE(sqlc.narg('compressor_details'), c.compressor_details),
  connector_type      = COALESCE(sqlc.narg('connector_type'), c.connector_type),
  voltage             = COALESCE(sqlc.narg('voltage'), c.voltage),
  shaft_type          = COALESCE(sqlc.narg('shaft_type'), c.shaft_type),
  notes               = COALESCE(sqlc.narg('notes'), c.notes)
WHERE c.part_no=$1
RETURNING *;

-- name: DeleteClutchAssyByPartNo :exec
DELETE FROM clutch_assys
WHERE part_no=$1;
