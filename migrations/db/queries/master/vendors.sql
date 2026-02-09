-------------------------------------------------------------------------------------------------------------------------------
-- QUERY
-------------------------------------------------------------------------------------------------------------------------------

-- name: GetVendorById :one
SELECT *
FROM vendors
WHERE id = $1;

-- name: GetVendorByCompanyName :one
SELECT *
FROM vendors
WHERE company_name=$1;

-- -- name: GetVendorsByCompanyName :many
-- SELECT *
-- FROM vendors
-- WHERE NULLIF($1,'') IS NULL OR company_name ILIKE '%' || LOWER($1) || '%'
-- ORDER BY company_name ASC;

-- name: GetVendorsByCompanyName :many
SELECT 
  vndrs.id,
  vndrs.company_name,
  vndr_cnts.contact_person,
  vndr_cnts.mobile_no,
  vndr_cnts.email_id
FROM vendor_contacts AS vndr_cnts
LEFT JOIN vendors AS vndrs ON vndrs.id = vndr_cnts.vendor_id
WHERE 
  NULLIF($1, '') IS NULL 
  OR LOWER(vndrs.company_name) ILIKE '%' || LOWER($1) || '%'
ORDER BY vndrs.company_name ASC;



-------------------------------------------------------------------------------------------------------------------------------
-- MUTATION
-------------------------------------------------------------------------------------------------------------------------------



-- name: CreateVendor :one
-- If you want to provide your own UUID occasionally
INSERT INTO vendors (
  id,
  company_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateVendorContacts :one
INSERT INTO vendor_contacts (
  id,
  vendor_id,
  contact_person,
  mobile_no,
  email_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateVendorByID :one
UPDATE vendors AS v
SET
  company_name   = COALESCE(sqlc.narg('company_name'), v.company_name),
  updated_at     = now()
WHERE v.id = sqlc.arg('id')
RETURNING *;



-- name: DeleteVendorByID :exec
DELETE FROM vendors
WHERE id = $1;

-- name: DeleteVendorContactsByID :exec
DELETE FROM vendor_contacts
WHERE vendor_id=$1;

-- name: DeleteVendorsByCompanyName :exec
DELETE FROM vendors
WHERE company_name ILIKE '%' || $1 || '%';
