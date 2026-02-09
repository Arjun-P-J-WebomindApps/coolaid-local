-- -----------------------------------------------------------------------------------------------------------------------------
-- --Queries
-- -----------------------------------------------------------------------------------------------------------------------------

-- -------------used to get a compressor valve by part number--#

-- name: GetCompressorValveByPartNo :one
SELECT *
FROM compressor_valves
WHERE part_no = $1;

-- -------------used to search compressor valves by type--#

-- name: GetCompressorValvesByType :many
SELECT *
FROM compressor_valves
WHERE type = $1;

-- name: GetCompressorValvesForDownload :many
SELECT * FROM compressor_valves;

-- -----------------------------------------------------------------------------------------------------------------------------
-- --Mutations
-- -----------------------------------------------------------------------------------------------------------------------------

-- name: CreateCompressorValve :one
INSERT INTO compressor_valves (
    id, part_no, type, voltage, connector_type, compressor_details, notes
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCompressorValveByPartNo :one
UPDATE compressor_valves AS cv
SET
    type               = COALESCE(sqlc.narg('type'), cv.type),
    voltage            = COALESCE(sqlc.narg('voltage'), cv.voltage),
    connector_type     = COALESCE(sqlc.narg('connector_type'), cv.connector_type),
    compressor_details = COALESCE(sqlc.narg('compressor_details'), cv.compressor_details),
    notes              = COALESCE(sqlc.narg('notes'), cv.notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteCompressorValveByPartNo :exec
DELETE FROM compressor_valves
WHERE part_no = $1;
