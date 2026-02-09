-----------------------------------------------------------------------------------------------------------------------------
--Queries
----------------------------------------------------------------------------------------------------------------------------

-------------used to identify if the input model exists in database or not--#

-- name: GetModelById :one
SELECT *
FROM models
WHERE id=$1;

-- name: GetModelByName :one
SELECT *
FROM models
WHERE name=$1;

-- name: GetModelByNameAndCompany :one
--params
SELECT m.id,
m.name as model_name,
c.name as company_name
FROM models m
JOIN companies c ON c.id = m.company_id
WHERE c.name = @CompanyName::text
  AND m.name = @ModelName::text;

--------------used in search engine to identify models like 800,i20 etc-----#

-- name: GetModelSimilar :many
SELECT *
FROM models
WHERE LOWER(name) ILIKE LOWER($1);

---------------- used by forms to get suggestions --------------------------#

-- name: GetModelsByCompanyAndModel :many
SELECT m.name as model_name, c.name as company_name, m.id,m.image_url
FROM models m
JOIN companies c ON c.id = m.company_id
WHERE
  ($1 = '' OR c.name ILIKE '%' || $1 || '%')
  AND
  ($2 = '' OR m.name ILIKE '%' || $2 || '%') ORDER BY m.name ASC;


------------------------------------------------------------------------------------------------------------------------------
--Mutation
------------------------------------------------------------------------------------------------------------------------------

-- name: CreateModel :one
INSERT INTO models
    (id, company_id, name, image_url)
VALUES
    ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateModelByID :one
UPDATE models AS m SET
  company_id = COALESCE(sqlc.narg('company_id'), m.company_id),
  name       = COALESCE(sqlc.narg('model_name'), m.name),
  image_url  = COALESCE(sqlc.narg('image_url'), m.image_url)
WHERE m.id = $1
RETURNING *;

-- name: DeleteModelByID :exec
DELETE FROM models
WHERE id = $1;

