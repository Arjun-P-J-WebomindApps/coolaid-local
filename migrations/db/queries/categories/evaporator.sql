-----------------------------------------------------------------------------------------------------------------------------
--Queries for Evaporators
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetEvaporatorByPartNo :one
SELECT *
FROM evaporators
WHERE part_no=$1;

-- name: GetEvaporatorsForDownload :many
SELECT * FROM evaporators;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Evaporators
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateEvaporator :one
INSERT INTO evaporators
    (id, part_no, mounting, exp_valve, additional_info, dimensions, pipe_connector, notes)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateEvaporatorByPartNo :one
UPDATE evaporators AS e SET
  mounting         = COALESCE(sqlc.narg('mounting'), e.mounting),
  exp_valve        = COALESCE(sqlc.narg('exp_valve'), e.exp_valve),
  additional_info  = COALESCE(sqlc.narg('additional_info'), e.additional_info),
  dimensions       = COALESCE(sqlc.narg('dimensions'), e.dimensions),
  pipe_connector   = COALESCE(sqlc.narg('pipe_connector'), e.pipe_connector),
  notes            = COALESCE(sqlc.narg('notes'), e.notes)
WHERE e.part_no=$1
RETURNING *;

-- name: DeleteEvaporatorByPartNo :exec
DELETE FROM evaporators
WHERE part_no=$1;
