-----------------------------------------------------------------------------------------------------------------------------
--Queries for Filter Driers
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetFilterDrierByPartNo :one
SELECT *
FROM filter_driers
WHERE part_no=$1;

-- name: GetFilterDriersForDownload :many
SELECT * FROM filter_driers;


-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Filter Driers
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateFilterDrier :one
INSERT INTO filter_driers
    (id, part_no, pipe_connector, size, pressure_switch, notes)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateFilterDrierByPartNo :one
UPDATE filter_driers AS fd SET
  pipe_connector  = COALESCE(sqlc.narg('pipe_connector'), fd.pipe_connector),
  size            = COALESCE(sqlc.narg('size'), fd.size),
  pressure_switch = COALESCE(sqlc.narg('pressure_switch'), fd.pressure_switch),
  notes           = COALESCE(sqlc.narg('notes'), fd.notes)
WHERE fd.part_no=$1
RETURNING *;

-- name: DeleteFilterDrierByPartNo :exec
DELETE FROM filter_driers
WHERE part_no=$1;
