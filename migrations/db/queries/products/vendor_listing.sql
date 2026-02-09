-- -------------------------------------------------------------------------------------------------------------------------------
-- QUERY
-- -------------------------------------------------------------------------------------------------------------------------------

-- name: GetVendorListingById :one
SELECT *
FROM vendor_listing
WHERE id = $1;

-- name: GetVendorListingsByProductPartNo :many
SELECT *
FROM vendor_listing
WHERE product_part_no = $1
ORDER BY vendor_name ASC, vendor_part_no ASC;

-- name: GetVendorListingsByVendorName :many
SELECT *
FROM vendor_listing
WHERE vendor_name ILIKE '%' || $1 || '%'
ORDER BY vendor_name ASC, product_part_no ASC;

-- name: GetVendorListingsByIds :many
SELECT
  id,
  product_part_no,
  vendor_name,
  vendor_part_no,
  vendor_mrp,
  created_at,
  updated_at
FROM vendor_listing
WHERE id = ANY($1::uuid[]);

-- name: GetVendorListingByProductAndVendorPart :one
SELECT *
FROM vendor_listing
WHERE product_part_no = $1
  AND vendor_name = $2
  AND vendor_part_no = $3;

-- -------------------------------------------------------------------------------------------------------------------------------
-- MUTATION
-- -------------------------------------------------------------------------------------------------------------------------------


-- name: CreateVendorListingWithID :one
INSERT INTO vendor_listing (
  id,
  product_part_no,
  vendor_name,
  vendor_part_no,
  vendor_mrp
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateVendorListingByID :one
UPDATE vendor_listing AS v
SET
  product_part_no = COALESCE(sqlc.narg('product_part_no'), v.product_part_no),
  vendor_name     = COALESCE(sqlc.narg('vendor_name'), v.vendor_name),
  vendor_part_no  = COALESCE(sqlc.narg('vendor_part_no'), v.vendor_part_no),
  vendor_mrp      = COALESCE(sqlc.narg('vendor_mrp'), v.vendor_mrp),
  updated_at      = now()
WHERE v.id = sqlc.arg('id')
RETURNING *;

-- name: DeleteVendorListingByID :exec
DELETE FROM vendor_listing
WHERE id = $1;

-- name: DeleteVendorListingsByProductPartNo :exec
DELETE FROM vendor_listing
WHERE product_part_no = $1;
