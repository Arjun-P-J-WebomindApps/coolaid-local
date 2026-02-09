--------------------------------------------------------------------------------------------------------------------------------
--Query (Tickets)
-------------------------------------------------------------------------------------------------------------------------------


-- -- name: FilterTickets :many
-- SELECT *
-- FROM tickets
-- WHERE ($1::date  IS NULL OR ticket_date::date = $1)
--   AND ($2::int   IS NULL OR daily_number      = $2)
--   AND ($3::text  IS NULL OR status            = $3)
-- ORDER BY ticket_date DESC;

-- name: FilterTickets :many
WITH p AS (
  SELECT NULLIF($1::text, '')::date AS d
)
SELECT
  t.id as ticket_id,
  t.ticket_date,
  t.daily_number,
  t.status,
  t.paf,
  t.transport_mode,
  t.inspection_images,
  t.check_for_damage,
  t.remove_all_labels,
  t.exclude_documents,
  t.urgent_requirement,
  t.images_included,
  c.customer_company_name AS customer_company_name,
  c.contact_person AS contact_person,
  c.mobile AS customer_phone,
  c.type as customer_type,
  c.payment_mode as payment_mode,
  i.id as item_id,
  i.part_no,
  i.model,
  i.quantity,
  i.unit_price,
  i.is_selected,
  b.name                     AS brand_name,
  cat.name                   AS category_name
FROM tickets t
LEFT JOIN ticket_items i ON i.ticket_id = t.id
LEFT JOIN customers c   ON c.id = t.customer_id
LEFT JOIN product_parts pp  ON pp.part_no  = i.part_no 
LEFT JOIN brands        b   ON b.id        = pp.brand_id
LEFT JOIN categories    cat ON cat.id      = pp.category_id
CROSS JOIN p
WHERE (p.d IS NULL OR (t.ticket_date >= p.d AND t.ticket_date < p.d + INTERVAL '1 day'))
  AND (NULLIF($2::text,'') IS NULL OR t.daily_number = $2)
  AND (NULLIF($3::text,'') IS NULL OR t.status = $3)
ORDER BY t.ticket_date DESC, t.id, i.id;

-------------------------------------------------------------------------------------------------------------------------------
--Mutation (Tickets)
------------------------------------------------------------------------------------------------------------------------------------

-- name: CreateTicket :one
INSERT INTO tickets (
  id,
  ticket_date,
  customer_id,
  daily_number,
  status,
  images_included,
  transport_mode,
  inspection_images,
  check_for_damage,
  remove_all_labels,
  exclude_documents,
  urgent_requirement,
  paf,
  created_at
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('ticket_date'),
  sqlc.arg('customer_id'),
  sqlc.arg('daily_number'),
  sqlc.arg('status'),
  COALESCE(sqlc.narg('images_included'), false),
  COALESCE(sqlc.narg('transport_mode'), 'porter'),
  COALESCE(sqlc.narg('inspection_images')::TEXT[], ARRAY[]::TEXT[]),
  COALESCE(sqlc.narg('check_for_damage'), false),
  COALESCE(sqlc.narg('remove_all_labels'), false),
  COALESCE(sqlc.narg('exclude_documents'), false),
  COALESCE(sqlc.narg('urgent_requirement'), false),
  COALESCE(sqlc.narg('paf'), 0),
  COALESCE(sqlc.narg('created_at'), now())
)
RETURNING *;


-- name: UpdateTicket :one
UPDATE tickets AS t
SET
  status              = COALESCE(sqlc.narg('status'), t.status),
  confirm_order_date  = COALESCE(sqlc.narg('confirm_order_date'), t.confirm_order_date),
  images_included     = COALESCE(sqlc.narg('images_included'), t.images_included),
  transport_mode      = COALESCE(NULLIF(sqlc.narg('transport_mode'), ''), t.transport_mode),
  inspection_images   = COALESCE(sqlc.narg('inspection_images')::TEXT[], t.inspection_images),
  check_for_damage    = COALESCE(sqlc.narg('check_for_damage'), t.check_for_damage),
  remove_all_labels   = COALESCE(sqlc.narg('remove_all_labels'), t.remove_all_labels),
  exclude_documents   = COALESCE(sqlc.narg('exclude_documents'), t.exclude_documents),
  urgent_requirement  = COALESCE(sqlc.narg('urgent_requirement'), t.urgent_requirement),
  paf                 = COALESCE(sqlc.narg('paf'), t.paf) 
WHERE t.ticket_date::date = sqlc.arg('ticket_date')
  AND t.daily_number      = sqlc.arg('daily_number')
RETURNING *;


-- name: DeletePendingTicketsOlderThan2Days :many
DELETE FROM tickets
WHERE status = 'pending'
  AND created_at < NOW() - INTERVAL '2 days'
RETURNING *;


-- name: DeleteTicketByDateAndDailyNumber :exec
DELETE FROM tickets
WHERE ticket_date::date = sqlc.arg('ticket_date')
  AND daily_number      = sqlc.arg('daily_number');



---------------------------------------------------------------------------------------------------------------------------------
--Query-(Ticket Item)
---------------------------------------------------------------------------------------------------------------------------------


-- name: GetTicketItemByTicketId :many
SELECT * FROM ticket_items WHERE ticket_id=$1;



-- name: GetNonSelectedTicketItemsByTicketId :many
SELECT *
FROM ticket_items
WHERE ticket_id = $1
  AND is_selected = false;



---------------------------------------------------------------------------------------------------------------------------------
--Mutation-(Ticket Item)
-----------------------------------------------------------------------------------------------------------------------------------

-- name: CreateTicketItem :one
INSERT INTO ticket_items (id,ticket_id,part_no,model,quantity,unit_price) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *;

-- name: UpdateTicketItem :one
UPDATE ticket_items AS ti SET
  model        = COALESCE(sqlc.narg('model'), ti.model),
  quantity     = COALESCE(sqlc.narg('quantity'), ti.quantity),
  unit_price   = COALESCE(sqlc.narg('unit_price'), ti.unit_price),
  is_selected  = COALESCE(sqlc.narg('is_selected'), ti.is_selected)
WHERE ti.ticket_id = sqlc.arg('ticket_id')
  AND ti.part_no   = sqlc.arg('part_no')
RETURNING *;


-- name: DeleteTicketItemsByTicketId :many
DELETE FROM ticket_items
WHERE ticket_id = $1
RETURNING *;


-- name: DeleteTicketItemsByIds :many
DELETE FROM ticket_items
WHERE id = ANY($1::uuid[])
RETURNING *;

