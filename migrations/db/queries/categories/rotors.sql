-----------------------------------------------------------------------------------------------------------------------------
--Queries for Rotors
-----------------------------------------------------------------------------------------------------------------------------

-- name: GetRotorByPartNo :one
SELECT *
FROM rotors
WHERE part_no=$1;


-- name: GetRotorsForDownload :many
SELECT * FROM rotors;

-----------------------------------------------------------------------------------------------------------------------------
--Mutation for Rotors
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateRotor :one
INSERT INTO rotors
    (id, part_no, pulley_ribs, pulley_size, compressor_details, notes)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateRotorByPartNo :one
UPDATE rotors AS r SET
  pulley_ribs         = COALESCE(sqlc.narg('pulley_ribs'), r.pulley_ribs),
  pulley_size         = COALESCE(sqlc.narg('pulley_size'), r.pulley_size),
  compressor_details  = COALESCE(sqlc.narg('compressor_details'), r.compressor_details),
  notes               = COALESCE(sqlc.narg('notes'), r.notes)
WHERE r.part_no=$1
RETURNING *;

-- name: DeleteRotorByPartNo :exec
DELETE FROM rotors
WHERE part_no=$1;
