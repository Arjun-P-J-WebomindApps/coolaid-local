-- -----------------------------------------------------------------------------------------------------------------------------
-- --Queries
-- -----------------------------------------------------------------------------------------------------------------------------

-- -------------used to get a cabin filter by part number--#

-- name: GetCabinFilterByPartNo :one
SELECT *
FROM cabin_filters
WHERE part_no = $1;

-- -------------used to search cabin filters by type--#

-- name: GetCabinFiltersByType :many
SELECT *
FROM cabin_filters
WHERE type = $1;

-- -------------used to search cabin filters by material--#

-- name: GetCabinFiltersByMaterial :many
SELECT *
FROM cabin_filters
WHERE material = $1;

-- name: GetCabinFiltersForDownload :many
SELECT * FROM cabin_filters;

-- -----------------------------------------------------------------------------------------------------------------------------
-- --Mutations
-- -----------------------------------------------------------------------------------------------------------------------------

-- name: CreateCabinFilter :one
INSERT INTO cabin_filters (
    id, part_no, type, dimensions, material, notes
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateCabinFilterByPartNo :one
UPDATE cabin_filters AS c
SET
    type       = COALESCE(sqlc.narg('type'), c.type),
    dimensions = COALESCE(sqlc.narg('dimensions'), c.dimensions),
    material   = COALESCE(sqlc.narg('material'), c.material),
    notes      = COALESCE(sqlc.narg('notes'), c.notes)
WHERE part_no = $1
RETURNING *;

-- name: DeleteCabinFilterByPartNo :exec
DELETE FROM cabin_filters
WHERE part_no = $1;
