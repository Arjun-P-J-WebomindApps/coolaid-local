---------------------------------------------------------------------
-- 📘 QUERIES (Read Operations)
-- These statements only READ data from the database.
-- They do not modify any existing data.
---------------------------------------------------------------------

-- name: GetInventoryProductByPartNo :one
SELECT *
FROM inventory
WHERE part_no = $1;

-- name: OutOfStockCommodities :many
SELECT inv.*, vndr.company_name as vendor,mdl.name as model,c.name as category,vndr_list.vendor_part_no as vendor_part_no
FROM inventory inv
  join vendors vndr on inv.vendor_id=vndr.id
  join product_parts pp  on pp.part_no=inv.part_no
  join models mdl on mdl.id=pp.model_id
  join categories c on c.id=pp.category_id
  LEFT JOIN vendor_listing vndr_list ON vndr_list.product_part_no = pp.part_no 
WHERE inv.is_requested_for_supply = FALSE
  AND inv.qty_in_stock < inv.minimum_order_level;
  
-- name: GetInventoryDetailsForBulkDownload :many
SELECT
  pp.part_no AS "Part No",
  c.name AS "Company",
  m.name AS "Model",
  b.name AS "Brand",
  cat.name AS "Category",
  inv.*
FROM inventory inv
  join product_parts pp on pp.part_no=inv.part_no
  JOIN companies c ON c.id = pp.company_id
  JOIN models m ON m.id = pp.model_id
  JOIN brands b ON b.id = pp.brand_id
  JOIN categories cat ON cat.id = pp.category_id
  JOIN model_variants vnt ON vnt.part_no=pp.part_no;



---------------------------------------------------------------------
-- ⚙️ MUTATIONS (Write Operations)
-- These statements CREATE, UPDATE, or DELETE data in the database.
-- They modify the state of the system.
---------------------------------------------------------------------

-- name: CreateInventoryProduct :one
INSERT INTO inventory
  (
  id,
  part_no,
  minimum_order_level,
  maximum_order_level,
  qty_in_stock,
  location,
  is_flash,
  is_requested_for_supply,
  vendor_id
  )
VALUES
  (
    $1, -- id
    $2, -- part_no
    $3, -- minimum_order_level
    $4, -- maximum_order_level
    $5, -- qty_in_stock
    $6, -- location
    $7, -- is_flash
    $8, -- is_requested_for_supply
    $9   -- vendor_id
)
RETURNING *;


-- name: UpdateInventoryProductByPartNo :one
UPDATE inventory AS inv
SET
  minimum_order_level
= COALESCE
($2, inv.minimum_order_level),
  maximum_order_level       = COALESCE
($3, inv.maximum_order_level),
  qty_in_stock              = COALESCE
($4, inv.qty_in_stock),
  location                  = COALESCE
($5, inv.location),
  is_flash                  = COALESCE
($6, inv.is_flash),
  is_requested_for_supply   = COALESCE
($7, inv.is_requested_for_supply),
  vendor_id                 = COALESCE
($8, inv.vendor_id)
WHERE part_no = $1
RETURNING *;

CREATE UNIQUE INDEX IF NOT EXISTS idx_inventory_part_no ON inventory(part_no);

---------------------------------------------------------------------
-- 🚀 BULK UPDATE (by part_no)
-- Updates many inventory rows in ONE statement using array UNNEST.
-- Keep arrays the same length and aligned by index.
---------------------------------------------------------------------

-- name: BulkUpdateInventorySelective :exec
WITH data AS (
  SELECT
    UNNEST(@part_nos::text[])  AS part_no,
    UNNEST(@qtys::int4[])      AS qty_in_stock_raw,
    UNNEST(@max_lvls::int4[])  AS maximum_order_level_raw,
    UNNEST(@min_lvls::int4[])  AS minimum_order_level_raw,
    UNNEST(@locations::text[]) AS location_raw,
    UNNEST(@is_flashes::bool[]) AS is_flash
)
UPDATE inventory inv
SET
  qty_in_stock        = COALESCE(NULLIF(d.qty_in_stock_raw,        -1), inv.qty_in_stock),
  maximum_order_level = COALESCE(NULLIF(d.maximum_order_level_raw, -1), inv.maximum_order_level),
  minimum_order_level = COALESCE(NULLIF(d.minimum_order_level_raw, -1), inv.minimum_order_level),
  location            = COALESCE(NULLIF(d.location_raw, ''),           inv.location),
  is_flash            = d.is_flash
FROM data d
WHERE inv.part_no = d.part_no;


-- name: DeleteInventoryProductByPartNo :exec
DELETE FROM inventory
WHERE part_no = $1;
