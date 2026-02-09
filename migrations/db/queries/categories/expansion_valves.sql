-----------------------------------------------------------------------------------------------------------------------------
--Queries for Expansion Valves
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetExpansionValveByPartNo :one
SELECT *
FROM expansion_valves
WHERE part_no=$1;


-- name: GetExpansionValvesForDownload :many
SELECT * FROM expansion_valves;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Expansion Valves
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateExpansionValve :one
INSERT INTO expansion_valves
    (id, part_no, type, material, refrigerant, notes)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateExpansionValveByPartNo :one
UPDATE expansion_valves AS ev SET
  type       = COALESCE(sqlc.narg('type'), ev.type),
  material   = COALESCE(sqlc.narg('material'), ev.material),
  refrigerant= COALESCE(sqlc.narg('refrigerant'), ev.refrigerant),
  notes      = COALESCE(sqlc.narg('notes'), ev.notes)
WHERE ev.part_no=$1
RETURNING *;

-- name: DeleteExpansionValveByPartNo :exec
DELETE FROM expansion_valves
WHERE part_no=$1;
