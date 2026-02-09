------------------------------------------------------------------------------------------------------------------------------------
--Query
------------------------------------------------------------------------------------------------------------------------------------

-- name: GetProductPricingFromProductId :one
SELECT *
FROM product_part_pricing
WHERE product_part_id=$1;


-- name: GetProductPricingDownloadDetails :many
SELECT 
  pp.part_no AS "Part No",
  c.name AS "Company",
  m.name AS "Model",
  b.name AS "Brand",
  cat.name AS "Category",

  -- Pricing info
  price.basic_price AS "Basic Price",
  price.freight AS "Freight",
  price.gst AS "GST",

  price.ac_workshop AS "Ac Workshop",
  price.ac_workshop_per AS "Ac Workshop Per",
  price.ac_workshop_amt AS "Ac Workshop Margin",

  price.multibrand_workshop AS "Multibrand Workshop",
  price.multibrand_workshop_per AS "Multibrand Workshop Per",
  price.multibrand_workshop_amt AS "Multibrand Workshop Margin",

  price.auto_trader AS "Auto Trader",
  price.auto_trader_per AS "Auto Trader Per",
  price.auto_trader_amt AS "Auto Trader Margin",

  price.ac_trader AS "Ac Trader",
  price.ac_trader_per AS "Ac Trader Per",
  price.ac_trader_amt AS "Ac Trader Margin",

  price.minimum_purchase_quantity AS "Minimum Purchase Quantity",
  price.oem_mrp AS "OEM MRP",
  price.unit_measure AS "Unit Measure",

  vnt.vendor_id AS "Vendor ID"

FROM product_part_pricing price
LEFT JOIN product_parts pp ON pp.id = price.product_part_id
JOIN companies c ON c.id = pp.company_id
JOIN models m ON m.id = pp.model_id
JOIN brands b ON b.id = pp.brand_id
JOIN categories cat ON cat.id = pp.category_id
JOIN model_variants vnt ON vnt.part_no=pp.part_no;


-----------------------------------------------------------------------------------------------------------------------------
--Mutation
-----------------------------------------------------------------------------------------------------------------------------

-- name: CreateProductPrice :one
INSERT INTO product_part_pricing (
  id,
  product_part_id,
  basic_price,
  freight,
  gst,
  tax,

  -- Ac Workshop
  ac_workshop,
  ac_workshop_per,
  ac_workshop_amt,

  -- Multibrand Workshop
  multibrand_workshop,
  multibrand_workshop_per,
  multibrand_workshop_amt,

  -- Auto Trader
  auto_trader,
  auto_trader_per,
  auto_trader_amt,

  -- AC Trader
  ac_trader,
  ac_trader_per,
  ac_trader_amt,

  -- Outstation
  outstation_class_a,
  outstation_note,

  -- MRPs and misc
  mrp_temp,
  oem_mrp,
  unit_measure,
  minimum_purchase_quantity,

  created_at,
  updated_at
)
VALUES (
    $1,  -- id
    $2,  -- product_part_id
    $3,  -- basic_price
    $4,  -- freight
    $5,  -- gst
    $6,  -- tax

    $7,  -- ac_workshop
    $8,  -- ac_workshop_per
    $9,  -- ac_workshop_amt

    $10, -- multibrand_workshop
    $11, -- multibrand_workshop_per
    $12, -- multibrand_workshop_amt

    $13, -- auto_trader
    $14, -- auto_trader_per
    $15, -- auto_trader_amt

    $16, -- ac_trader
    $17, -- ac_trader_per
    $18, -- ac_trader_amt

    $19, -- outstation_class_a
    $20, -- outstation_note

    $21, -- mrp_temp
    $22, -- oem_mrp
    $23, -- unit_measure
    $24, -- minimum_purchase_quantity

    $25, -- created_at
    $26  -- updated_at
)
RETURNING *;





-- name: UpdateProductPriceByID :one
UPDATE product_part_pricing AS p
SET
  basic_price             = COALESCE(sqlc.narg('new_basic_price'), p.basic_price),
  freight                 = COALESCE(sqlc.narg('new_freight'), p.freight),
  gst                     = COALESCE(sqlc.narg('new_gst'), p.gst),
  tax                     = COALESCE(sqlc.narg('new_tax'), p.tax),

  -- Ac Workshop (formerly retail)
  ac_workshop             = COALESCE(sqlc.narg('new_ac_workshop'), p.ac_workshop),
  ac_workshop_per         = COALESCE(sqlc.narg('new_ac_workshop_per'), p.ac_workshop_per),
  ac_workshop_amt         = COALESCE(sqlc.narg('new_ac_workshop_amt'), p.ac_workshop_amt),

  -- Multibrand Workshop (formerly wholesale)
  multibrand_workshop     = COALESCE(sqlc.narg('new_multibrand_workshop'), p.multibrand_workshop),
  multibrand_workshop_per = COALESCE(sqlc.narg('new_multibrand_workshop_per'), p.multibrand_workshop_per),
  multibrand_workshop_amt = COALESCE(sqlc.narg('new_multibrand_workshop_amt'), p.multibrand_workshop_amt),

  -- Auto Trader (formerly mrp)
  auto_trader             = COALESCE(sqlc.narg('new_auto_trader'), p.auto_trader),
  auto_trader_per         = COALESCE(sqlc.narg('new_auto_trader_per'), p.auto_trader_per),
  auto_trader_amt         = COALESCE(sqlc.narg('new_auto_trader_amt'), p.auto_trader_amt),

  -- AC Trader
  ac_trader               = COALESCE(sqlc.narg('new_ac_trader'), p.ac_trader),
  ac_trader_per           = COALESCE(sqlc.narg('new_ac_trader_per'), p.ac_trader_per),
  ac_trader_amt           = COALESCE(sqlc.narg('new_ac_trader_amt'), p.ac_trader_amt),

  -- Outstation & notes
  outstation_class_a      = COALESCE(sqlc.narg('new_outstation_class_a'), p.outstation_class_a),
  outstation_note         = COALESCE(sqlc.narg('new_outstation_note'), p.outstation_note),

  -- MRPs & misc
  mrp_temp                = COALESCE(sqlc.narg('new_mrp_temp'), p.mrp_temp),
  oem_mrp                 = COALESCE(sqlc.narg('new_oem_mrp'), p.oem_mrp),
  unit_measure            = COALESCE(sqlc.narg('new_unit_measure'), p.unit_measure),
  minimum_purchase_quantity = COALESCE(sqlc.narg('new_minimum_purchase_quantity'), p.minimum_purchase_quantity),

  updated_at              = NOW()
FROM product_parts AS pp
WHERE pp.part_no = sqlc.arg(part_no)
  AND p.product_part_id = pp.id
RETURNING p.*;




-- name: UpdateProductPriceByProductPartID :one
UPDATE product_part_pricing AS p
SET
  basic_price             = COALESCE(sqlc.narg('new_basic_price'), p.basic_price),
  freight                 = COALESCE(sqlc.narg('new_freight'), p.freight),
  tax                     = COALESCE(sqlc.narg('new_tax'), p.tax),
  gst                     = COALESCE(sqlc.narg('new_gst'), p.gst),

  -- Ac Workshop (was retail)
  ac_workshop             = COALESCE(sqlc.narg('new_ac_workshop'), p.ac_workshop),
  ac_workshop_per         = COALESCE(sqlc.narg('new_ac_workshop_per'), p.ac_workshop_per),
  ac_workshop_amt         = COALESCE(sqlc.narg('new_ac_workshop_amt'), p.ac_workshop_amt),

  -- Multibrand Workshop (was wholesale)
  multibrand_workshop     = COALESCE(sqlc.narg('new_multibrand_workshop'), p.multibrand_workshop),
  multibrand_workshop_per = COALESCE(sqlc.narg('new_multibrand_workshop_per'), p.multibrand_workshop_per),
  multibrand_workshop_amt = COALESCE(sqlc.narg('new_multibrand_workshop_amt'), p.multibrand_workshop_amt),

  -- Auto Trader (was mrp)
  auto_trader             = COALESCE(sqlc.narg('new_auto_trader'), p.auto_trader),
  auto_trader_per         = COALESCE(sqlc.narg('new_auto_trader_per'), p.auto_trader_per),
  auto_trader_amt         = COALESCE(sqlc.narg('new_auto_trader_amt'), p.auto_trader_amt),

  -- AC Trader
  ac_trader               = COALESCE(sqlc.narg('new_ac_trader'), p.ac_trader),
  ac_trader_per           = COALESCE(sqlc.narg('new_ac_trader_per'), p.ac_trader_per),
  ac_trader_amt           = COALESCE(sqlc.narg('new_ac_trader_amt'), p.ac_trader_amt),

  -- Outstation
  outstation_class_a      = COALESCE(sqlc.narg('new_outstation_class_a'), p.outstation_class_a),
  outstation_note         = COALESCE(sqlc.narg('new_outstation_note'), p.outstation_note),

  -- MRPs & misc
  mrp_temp                = COALESCE(sqlc.narg('new_mrp_temp'), p.mrp_temp),
  oem_mrp                 = COALESCE(sqlc.narg('new_oem_mrp'), p.oem_mrp),
  unit_measure            = COALESCE(sqlc.narg('new_unit_measure'), p.unit_measure),
  minimum_purchase_quantity = COALESCE(sqlc.narg('new_minimum_purchase_quantity'), p.minimum_purchase_quantity),


  updated_at              = NOW()
WHERE p.product_part_id = $1
RETURNING *;



-- name: UpdateProductPriceByPartNo :one
UPDATE product_part_pricing AS p
SET
  basic_price             = COALESCE(sqlc.narg('new_basic_price'), p.basic_price),
  freight                 = COALESCE(sqlc.narg('new_freight'), p.freight),
  tax                     = COALESCE(sqlc.narg('new_tax'), p.tax),
  gst                     = COALESCE(sqlc.narg('new_gst'), p.gst),

  -- Ac Workshop (was retail)
  ac_workshop             = COALESCE(sqlc.narg('new_ac_workshop'), p.ac_workshop),
  ac_workshop_per         = COALESCE(sqlc.narg('new_ac_workshop_per'), p.ac_workshop_per),
  ac_workshop_amt         = COALESCE(sqlc.narg('new_ac_workshop_amt'), p.ac_workshop_amt),

  -- Multibrand Workshop (was wholesale)
  multibrand_workshop     = COALESCE(sqlc.narg('new_multibrand_workshop'), p.multibrand_workshop),
  multibrand_workshop_per = COALESCE(sqlc.narg('new_multibrand_workshop_per'), p.multibrand_workshop_per),
  multibrand_workshop_amt = COALESCE(sqlc.narg('new_multibrand_workshop_amt'), p.multibrand_workshop_amt),

  -- Auto Trader (was mrp)
  auto_trader             = COALESCE(sqlc.narg('new_auto_trader'), p.auto_trader),
  auto_trader_per         = COALESCE(sqlc.narg('new_auto_trader_per'), p.auto_trader_per),
  auto_trader_amt         = COALESCE(sqlc.narg('new_auto_trader_amt'), p.auto_trader_amt),

  -- AC Trader (formerly Outstation Class B)
  ac_trader               = COALESCE(sqlc.narg('new_ac_trader'), p.ac_trader),
  ac_trader_per           = COALESCE(sqlc.narg('new_ac_trader_per'), p.ac_trader_per),
  ac_trader_amt           = COALESCE(sqlc.narg('new_ac_trader_amt'), p.ac_trader_amt),

  -- Outstation
  outstation_class_a      = COALESCE(sqlc.narg('new_outstation_class_a'), p.outstation_class_a),
  outstation_note         = COALESCE(sqlc.narg('new_outstation_note'), p.outstation_note),

  -- MRPs and misc
  mrp_temp                = COALESCE(sqlc.narg('new_mrp_temp'), p.mrp_temp),
  oem_mrp                 = COALESCE(sqlc.narg('new_oem_mrp'), p.oem_mrp),
  unit_measure            = COALESCE(sqlc.narg('new_unit_measure'), p.unit_measure),
  minimum_purchase_quantity = COALESCE(sqlc.narg('new_minimum_purchase_quantity'), p.minimum_purchase_quantity),

  updated_at              = NOW()
FROM product_parts AS pp
WHERE pp.id = p.product_part_id
  AND pp.part_no = sqlc.arg(part_no)
RETURNING *;


-- name: DeleteProductPriceByID :exec
DELETE FROM product_part_pricing
WHERE id = $1;

-- name: DeleteProductPriceByProductPartID :exec
DELETE FROM product_part_pricing
WHERE product_part_id = $1;





-- name: BulkUpdatePricingSelective :exec
WITH data AS (
  SELECT
    UNNEST(@part_nos::text[]) AS part_no,

    UNNEST(@basic_prices::numeric[]) AS basic_price,
    UNNEST(@freights::numeric[]) AS freight,
    UNNEST(@gst_vals::numeric[]) AS gst,
    UNNEST(@tax_vals::numeric[]) AS tax,

    UNNEST(@ac_workshops::numeric[]) AS ac_workshop,
    UNNEST(@ac_workshop_pers::numeric[]) AS ac_workshop_per,
    UNNEST(@ac_workshop_amts::numeric[]) AS ac_workshop_amt,

    UNNEST(@multibrand_workshops::numeric[]) AS multibrand_workshop,
    UNNEST(@multibrand_workshop_pers::numeric[]) AS multibrand_workshop_per,
    UNNEST(@multibrand_workshop_amts::numeric[]) AS multibrand_workshop_amt,

    UNNEST(@auto_traders::numeric[]) AS auto_trader,
    UNNEST(@auto_trader_pers::numeric[]) AS auto_trader_per,
    UNNEST(@auto_trader_amts::numeric[]) AS auto_trader_amt,

    UNNEST(@ac_traders::numeric[]) AS ac_trader,
    UNNEST(@ac_trader_pers::numeric[]) AS ac_trader_per,
    UNNEST(@ac_trader_amts::numeric[]) AS ac_trader_amt,

    UNNEST(@outstation_class_as::numeric[]) AS outstation_class_a,
    UNNEST(@outstation_notes::text[]) AS outstation_note,

    UNNEST(@mrp_temps::numeric[]) AS mrp_temp,
    UNNEST(@oem_mrps::numeric[]) AS oem_mrp,
    UNNEST(@unit_measures::text[]) AS unit_measure,
    UNNEST(@min_purchase_qtys::int4[]) AS minimum_purchase_quantity
)
UPDATE product_part_pricing p
SET
  basic_price = COALESCE(NULLIF(d.basic_price, -1), p.basic_price),
  freight = COALESCE(NULLIF(d.freight, -1), p.freight),
  gst = COALESCE(NULLIF(d.gst, -1), p.gst),
  tax = COALESCE(NULLIF(d.tax, -1), p.tax),

  ac_workshop = COALESCE(NULLIF(d.ac_workshop, -1), p.ac_workshop),
  ac_workshop_per = COALESCE(NULLIF(d.ac_workshop_per, -1), p.ac_workshop_per),
  ac_workshop_amt = COALESCE(NULLIF(d.ac_workshop_amt, -1), p.ac_workshop_amt),

  multibrand_workshop = COALESCE(NULLIF(d.multibrand_workshop, -1), p.multibrand_workshop),
  multibrand_workshop_per = COALESCE(NULLIF(d.multibrand_workshop_per, -1), p.multibrand_workshop_per),
  multibrand_workshop_amt = COALESCE(NULLIF(d.multibrand_workshop_amt, -1), p.multibrand_workshop_amt),

  auto_trader = COALESCE(NULLIF(d.auto_trader, -1), p.auto_trader),
  auto_trader_per = COALESCE(NULLIF(d.auto_trader_per, -1), p.auto_trader_per),
  auto_trader_amt = COALESCE(NULLIF(d.auto_trader_amt, -1), p.auto_trader_amt),

  ac_trader = COALESCE(NULLIF(d.ac_trader, -1), p.ac_trader),
  ac_trader_per = COALESCE(NULLIF(d.ac_trader_per, -1), p.ac_trader_per),
  ac_trader_amt = COALESCE(NULLIF(d.ac_trader_amt, -1), p.ac_trader_amt),

  outstation_class_a = COALESCE(NULLIF(d.outstation_class_a, -1), p.outstation_class_a),
  outstation_note = COALESCE(NULLIF(d.outstation_note, ''), p.outstation_note),

  mrp_temp = COALESCE(NULLIF(d.mrp_temp, -1), p.mrp_temp),
  oem_mrp = COALESCE(NULLIF(d.oem_mrp, -1), p.oem_mrp),
  unit_measure = COALESCE(NULLIF(d.unit_measure, ''), p.unit_measure),
  minimum_purchase_quantity = COALESCE(NULLIF(d.minimum_purchase_quantity, -1), p.minimum_purchase_quantity),

  updated_at = NOW()
FROM data d
JOIN product_parts pp ON pp.part_no = d.part_no
WHERE p.product_part_id = pp.id;
