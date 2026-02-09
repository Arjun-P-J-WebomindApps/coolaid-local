-----------------------------------------------------------------------------------------------------------------------------
--Queries
-----------------------------------------------------------------------------------------------------------------------------

-------------used to get a condenser by part number--#

-- name: GetCondenserByPartNo :one
SELECT *
FROM condensers
WHERE part_no = $1;


-- name: GetCondensersForDownload :many
SELECT * FROM condensers;


------------------------------------------------------------------------------------------------------------------------------
--Mutation
------------------------------------------------------------------------------------------------------------------------------

-- name: CreateCondenser :one
INSERT INTO condensers (
    id, part_no, size, pipe_connector, drier, pressure_switch, notes
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCondenserByPartNo :one
UPDATE condensers AS c
SET
    size           = COALESCE(sqlc.narg('size'), c.size),
    pipe_connector = COALESCE(sqlc.narg('pipe_connector'), c.pipe_connector),
    drier          = COALESCE(sqlc.narg('drier'), c.drier),
    pressure_switch= COALESCE(sqlc.narg('pressure_switch'), c.pressure_switch),
    notes          = COALESCE(sqlc.narg('notes'), c.notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteCondenserByPartNo :exec
DELETE FROM condensers
WHERE part_no = $1;
