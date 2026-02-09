-----------------------------------------------------------------------------------------------------------------------------
--Compressors Queries
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetCompressorByPartNo :one
SELECT *
FROM compressors
WHERE part_no = $1;

-- name: GetCompressorsForDownload :many
SELECT * FROM compressors;

-----------------------------------------------------------------------------------------------------------------------------
--Compressors Mutation
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateCompressor :one
INSERT INTO compressors (
    id,
    part_no,
    compressor_id,
    oil,
    refrigerant,
    voltage,
    pulley_ribs,
    pulley_size,
    pipe_connector,
    comp_type,
    comp_mounting,
    connector_type,
    notes
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateCompressorByPartNo :one
UPDATE compressors
SET
    compressor_id  = COALESCE(sqlc.narg('compressor_id'), compressor_id),
    oil            = COALESCE(sqlc.narg('oil'), oil),
    refrigerant    = COALESCE(sqlc.narg('refrigerant'), refrigerant),
    voltage        = COALESCE(sqlc.narg('voltage'), voltage),
    pulley_ribs    = COALESCE(sqlc.narg('pulley_ribs'), pulley_ribs),
    pulley_size    = COALESCE(sqlc.narg('pulley_size'), pulley_size),
    pipe_connector = COALESCE(sqlc.narg('pipe_connector'), pipe_connector),
    comp_type      = COALESCE(sqlc.narg('comp_type'), comp_type),
    comp_mounting  = COALESCE(sqlc.narg('comp_mounting'), comp_mounting),
    connector_type = COALESCE(sqlc.narg('connector_type'), connector_type),
    notes          = COALESCE(sqlc.narg('notes'), notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteCompressorByPartNo :exec
DELETE FROM compressors
WHERE part_no = $1;
