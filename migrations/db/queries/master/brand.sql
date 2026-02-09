-------------------------------------------------------------------------------------------------------------------------------
--QUERY
-----------------------------------------------------------------------------------------------------------------------------

-------Check if the Brand exists-----------------------------------

-- name: GetBrandById :one
SELECT *
FROM brands
WHERE id=$1;

-- name: GetBrandByName :one
SELECT *
FROM brands
WHERE name=$1;


-----------Brand List in Suggestions------------------------------
-- name: GetBrandListByName :many
SELECT *
FROM brands
WHERE NULLIF(name,'') IS NULL OR LOWER(name) ILIKE '%' || LOWER($1) || '%' ORDER BY name ASC;;

-------------------------------------------------------------------------------------------------------------------------
--Mutation
------------------------------------------------------------------------------------------------------------------------------

-- name: CreateBrand :one 
INSERT INTO brands
    (id,name,image)
VALUES
    ($1, $2,$3)
RETURNING *;

-- name: UpdateBrandByID :one
UPDATE brands AS b SET
  name = COALESCE(sqlc.narg('brand_name'), b.name),
  image = COALESCE(sqlc.narg('image_url'), b.image)
WHERE b.id = $1
RETURNING *;

-- name: DeleteBrandByID :exec
DELETE FROM brands
WHERE id = $1;
