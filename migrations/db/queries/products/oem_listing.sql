-- -------------------------------------------------------------------------------------------------------------------------------
-- QUERY
-- -------------------------------------------------------------------------------------------------------------------------------

-- name: GetOemListingById :one
SELECT *
FROM oem_listings
WHERE id = $1;

-- name: GetOemListingsByPartNo :many
SELECT *
FROM oem_listings
WHERE part_no = $1
ORDER BY oem_number ASC;

-- name: GetOemListingsByOemNumber :many
SELECT *
FROM oem_listings
WHERE oem_number = $1
ORDER BY part_no ASC;

-- name: GetOemListingsByIds :many
SELECT
  id,
  part_no,
  oem_number,
  price,
  created_at,
  updated_at
FROM oem_listings
WHERE id = ANY($1::uuid[]);

-- name: GetOemListingByPartAndOemNumber :one
SELECT *
FROM oem_listings
WHERE part_no = $1
  AND oem_number = $2;

-- -------------------------------------------------------------------------------------------------------------------------------
-- MUTATION
-- -------------------------------------------------------------------------------------------------------------------------------

-- name: CreateOemListingWithID :one
INSERT INTO oem_listings (
  id,
  part_no,
  oem_number,
  price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- Optional: if you also want server-generated UUIDs (no ID passed from app)
-- name: CreateOemListing :one
INSERT INTO oem_listings (
  part_no,
  oem_number,
  price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateOemListingByID :one
UPDATE oem_listings AS o
SET
  part_no    = COALESCE(sqlc.narg('part_no'),    o.part_no),
  oem_number = COALESCE(sqlc.narg('oem_number'), o.oem_number),
  price      = COALESCE(sqlc.narg('price'),      o.price),
  updated_at = now()
WHERE o.id = sqlc.arg('id')
RETURNING *;

-- name: DeleteOemListingByID :exec
DELETE FROM oem_listings
WHERE id = $1;

-- name: DeleteOemListingsByPartNo :exec
DELETE FROM oem_listings
WHERE part_no = $1;

-- Optional: delete by OEM number
-- name: DeleteOemListingsByOemNumber :exec
DELETE FROM oem_listings
WHERE oem_number = $1;
