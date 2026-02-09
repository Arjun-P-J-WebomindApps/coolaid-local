---------------------------------------------------------------------
-- 📘 QUERIES (Read Operations)
-- These statements only READ data from the database.
---------------------------------------------------------------------

-- name: GetOrdersByUserId :many
SELECT *
FROM orders
WHERE user_id = $1;

-- name: GetAllRestockOrders :many
SELECT
    o.id,
    o.created_at AS order_date,
    v.company_name AS vendor_name,
    o.is_requested_via_email AS is_requested_by_email,
    o.is_requested_via_whatsapp AS is_requested_by_whatsapp,
    o2.part_no AS part_number,
    o2.quantity AS quantity_ordered
FROM
    orders AS o
LEFT JOIN
    orders_listing AS o2
    ON o2.id = ANY(o.order_listing_ids::uuid[])
LEFT JOIN
    vendors AS v
    ON v.id = o.vendor_id
ORDER BY
    o.created_at DESC;

---------------------------------------------------------------------
-- ⚙️ MUTATIONS (Write Operations)
-- These statements CREATE, UPDATE, or DELETE data in the database.
---------------------------------------------------------------------

-- name: CreateOrder :one
INSERT INTO orders (
    id,
    order_listing_ids,
    vendor_id,
    user_id,
    is_requested_via_email,
    is_requested_via_whatsapp
)
VALUES (
    $1,  -- id
    $2,  -- order_listing_ids (ARRAY)
    $3,  -- vendor_id
    $4,  -- user_id
    $5,  -- is_requested_via_email
    $6   -- is_requested_via_whatsapp
)
RETURNING *;

-- name: UpdateOrderStatusById :one
UPDATE orders SET
  is_requested_via_email = COALESCE($2, is_requested_via_email),
  is_requested_via_whatsapp = COALESCE($3, is_requested_via_whatsapp)
WHERE id = $1
RETURNING *;

-- name: DeleteOrderById :exec
DELETE FROM orders
WHERE id = $1;
