------------------------------------------------------------------------------------------------------------------------------
--Queries
------------------------------------------------------------------------------------------------------------------------------

------used in identifying whether catgeory exists before using

-- name: GetCategoryById :one
SELECT *
FROM categories
WHERE id=$1;

-- name: GetCategoryByName :one
SELECT *
FROM categories
WHERE name=$1;

-------used in form suggestion API 

-- name: GetCategoriesByName :many
SELECT *
FROM categories
WHERE NULLIF($1,'') IS NULL OR UPPER(name)
ILIKE '%'||$1||'%' ORDER BY name ASC;;


------------------------------------------------------------------------------------------------------------------------------
--Mutation
------------------------------------------------------------------------------------------------------------------------------

-- name: CreateCategories :one
INSERT INTO categories
    (id,name,image)
VALUES
    ($1, $2, $3)
RETURNING *;

-- name: UpdateCategoryByID :one
UPDATE categories AS c SET
  name  = COALESCE(sqlc.narg('category_name'), c.name),
  image = COALESCE(sqlc.narg('category_image'), c.image)
WHERE c.id = $1
RETURNING *;

-- name: DeleteCategoryByID :exec
DELETE FROM categories
WHERE id = $1;
