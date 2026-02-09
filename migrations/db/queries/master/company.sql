------------------------------------------------------------------------------------------------------------------------------
--Queries
------------------------------------------------------------------------------------------------------------------------------

-- used to find whether the  given company name exists in the db

-- name: GetCompanyByName :one
SELECT *
FROM companies
WHERE name=$1;

-- name: GetCompanyById :one
SELECT *
FROM companies
WHERE id=$1;


-- used for accessing existing company list (add forms etc)

-- name: GetCompaniesByName :many   
SELECT * FROM Companies WHERE  NULLIF(name,'') IS NULL OR LOWER(name) ILIKE '%' || LOWER($1) || '%' ORDER BY name ASC;


-----------------------------------------------------------------------------------------------------------------------------------
-- Mutation
-----------------------------------------------------------------------------------------------------------------------------------

-- name: CreateCompanies :one
INSERT INTO companies
    (id,name,status,image_url)
VALUES
    ($1, $2, $3,$4)
RETURNING *;

-- name: UpdateCompanyByID :one
UPDATE companies AS c SET
  name   = COALESCE(sqlc.narg('company_name'), c.name),
  status = COALESCE(sqlc.narg('company_status'), c.status),
  image_url = COALESCE(sqlc.narg('image_url'), c.image_url)
WHERE c.id = $1
RETURNING *;

-- name: UpdateCompanyByName :one
UPDATE companies AS c SET
  name   = COALESCE(sqlc.narg('company_name'), c.name),
  status = COALESCE(sqlc.narg('company_status'), c.status),
  image_url = COALESCE(sqlc.narg('image_url'), c.image_url)
WHERE c.name = $1
RETURNING *;

-- -- name: SoftDeleteCompanyByID :one
-- UPDATE companies
-- SET status = 'deleted'
-- WHERE id = $1
-- RETURNING *;

-- -- name: SoftDeleteCompanyByName :one
-- UPDATE companies
-- SET status = 'deleted'
-- WHERE name = $1
-- RETURNING *;

-- -- name: DeleteCompanyByID :one
-- DELETE FROM companies
-- WHERE id = $1
-- RETURNING *;
