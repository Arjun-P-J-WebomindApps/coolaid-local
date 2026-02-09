---------------------------------------------------------------------
-- 📘 QUERIES (Read Operations)
-- These statements only READ data from the database.
---------------------------------------------------------------------

-- name: GetOrderListingByOrderId :many
SELECT *
FROM orders_listing
WHERE id = $1;



---------------------------------------------------------------------
-- ⚙️ MUTATIONS (Write Operations)
-- These statements CREATE, UPDATE, or DELETE data in the database.
---------------------------------------------------------------------

-- name: CreateOrderListing :one
INSERT INTO orders_listing (
    id,
    part_no,
    quantity
)
VALUES (
    $1,  -- id
    $2,  -- part_no
    $3   -- quantity
)
RETURNING *;

-- name: UpdateOrderListingById :one
UPDATE orders_listing
SET
  part_no = COALESCE($2, part_no),
  quantity = COALESCE($3, quantity)
WHERE id = $1
RETURNING *;

-- name: DeleteOrderListingById :exec
DELETE FROM orders_listing
WHERE id = $1;
