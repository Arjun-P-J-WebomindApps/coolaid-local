-----------------------------------------------------------------------------------------------------------------------------
--Query
----------------------------------------------------------------------------------------------------------------------------

-- name: GetProductPartsByPartNo :one
SELECT * FROM product_parts WHERE part_no=$1;

-- name: GetProductPartInfoByPartNo :one
SELECT 
  cnpy.name AS company_name,
  mdl.name AS model_name,
  ctgr.name AS category_name,
  brnd.name AS brand_name
FROM product_parts p
LEFT JOIN companies cnpy ON cnpy.id = p.company_id
LEFT JOIN models mdl ON mdl.id = p.model_id
LEFT JOIN categories ctgr ON ctgr.id = p.category_id
LEFT JOIN brands brnd ON brnd.id = p.brand_id
WHERE p.part_no = $1;


-- name: GetProductPartNo :many
SELECT part_no from product_parts WHERE part_no=$1;

-- name: GetProductPartNos :many
SELECT part_no
FROM product_parts
WHERE part_no ILIKE '%' || $1 || '%'
ORDER BY part_no
LIMIT 10;


-- name: GetFilteredPartNo :many
SELECT *
FROM model_variants mv
WHERE
  -- no search term => return everything
  NULLIF($1, '') IS NULL
  OR LOWER(mv.part_no) ILIKE '%' || LOWER($1) || '%'
  OR EXISTS (
    SELECT 1
    FROM unnest(mv.oem_ids) AS oid
    JOIN oem_listings ol
      ON ol.id::text = oid           -- oem_ids is text[], id is UUID
    WHERE LOWER(ol.oem_number) ILIKE '%' || LOWER($1) || '%'
  );



-- name: GetFilteredDetails :many
-- @param CompanyName:sql.NullString
-- @param ModelName:sql.NullString
-- @param BrandName:sql.NullString
-- @param CategoryName:sql.NullString

SELECT
  pp.id,
  pp.part_no,
  c.name AS company_name,
  c.image_url AS company_image,
  m.name AS model_name,
  m.image_url AS model_image,
  b.name AS brand_name,
  b.image as brand_image,
  cat.name AS category_name,
  cat.image AS category_image
FROM product_parts pp
  JOIN companies c ON pp.company_id = c.id
  JOIN models m ON pp.model_id = m.id
  JOIN brands b ON pp.brand_id = b.id
  JOIN categories cat ON pp.category_id = cat.id
  JOIN model_variants mv ON pp.part_no = mv.part_no
WHERE
  -- Brand filter (only when non-empty)
  (NULLIF(@BrandName::text, '') IS NULL OR b.name = @BrandName::text)
  AND
  -- Category filter (only when non-empty)
  (NULLIF(@CategoryName::text, '') IS NULL OR cat.name = @CategoryName::text)
  AND
  (
    -- Model filter (only when non-empty)
    (
      NULLIF(@CompanyName::text, '') IS NOT NULL
      AND c.name = @CompanyName::text
      AND NULLIF(@ModelName::text, '') IS NOT NULL
      AND m.name = @ModelName::text
    )
    OR
    -- Unicode filter (only when non-empty)
    (
      NULLIF(@Unicode::text, '') IS NOT NULL
      AND EXISTS (
        SELECT 1
        FROM unnest(COALESCE(mv.unicode, '{}')) AS u
        WHERE u = @Unicode::text
      )
    )
    OR
    -- If neither model nor unicode provided → allow all
    (
      (NULLIF(@CompanyName::text, '') IS NULL OR c.name = @CompanyName::text)
      AND (NULLIF(@ModelName::text, '') IS NULL)
      AND (NULLIF(@Unicode::text, '') IS NULL)
    )
  );


-- name: GetFilteredProductsWithAllDetails :many
-- @param CompanyName:sql.NullString
-- @param ModelName:sql.NullString
-- @param PartNo:sql.NullString
-- @param CategoryList:text[]
-- @param BrandList:text[]
-- @param GenList:text[]
-- @param FuelList:text[]
-- @param MakeList:text[]
-- @param UnicodeCategoryList:text[]

SELECT
  pp.part_no,

  -- Master info
  c.name AS company_name,
  m.name AS model_name,
  b.name AS brand_name,
  b.image AS brand_image,
  cat.name AS category_name,

  -- Pricing info
  pr.basic_price,
  pr.freight,
  pr.gst,
  pr.tax,

  -- Ac Workshop (formerly retail)
  pr.ac_workshop,
  pr.ac_workshop_per,
  pr.ac_workshop_amt,

  -- Multibrand Workshop (formerly wholesale)
  pr.multibrand_workshop,
  pr.multibrand_workshop_per,
  pr.multibrand_workshop_amt,

  -- Auto Trader (formerly mrp)
  pr.auto_trader,
  pr.auto_trader_per,
  pr.auto_trader_amt,

  -- AC Trader (formerly Outstation Class B)
  pr.ac_trader,
  pr.ac_trader_per,
  pr.ac_trader_amt,

  -- Outstation & notes
  pr.outstation_class_a,
  pr.outstation_note,

  -- MRPs and misc
  pr.minimum_purchase_quantity,
  pr.mrp_temp,
  pr.oem_mrp,
  pr.unit_measure,


  --Inventory Info
  inv.qty_in_stock,
  inv.minimum_order_level,
  inv.maximum_order_level,
  inv.location,
  inv.is_flash,
  inv.is_requested_for_supply,
  inv.requested_date,

  --Offer Info
  ofr.is_offer_active AS offer_status,
  ofr.start_date      AS offer_start_date,
  ofr.end_date        AS offer_end_date,

  -- pricing-oriented aliases
  ofr.ac_trader       AS ac_trader_price,
  ofr.multi_brand     AS multi_brand_price,
  ofr.autotrader      AS autotrader_price,
  ofr.ac_workshop     AS ac_workshop_price,


  -- Variants
  vnt.part_no,
  vnt.fuel_types,
  vnt.type,
  vnt.gen,
  vnt.engine_cc,
  vnt.transmission_type,
  vnt.platform_codes,
  vnt.additional_info,
  vnt.image_1_link,
  vnt.image_2_link,
  vnt.image_3_link,
  vnt.image_4_link,
  vnt.hsn_code,
  vnt.unicode,
  vnt.description,
  vnt.oem_ids,
  vnt.make,
  vnt.vendor_id,
  vnt.placement,
  vnt.year_start,
  vnt.year_end

FROM product_parts pp
  JOIN companies c ON pp.company_id = c.id
  JOIN models m ON pp.model_id = m.id
  JOIN brands b ON pp.brand_id = b.id
  JOIN categories cat ON pp.category_id = cat.id
  LEFT JOIN product_part_pricing pr ON pr.product_part_id = pp.id
  LEFT JOIN inventory inv ON inv.part_no = pp.part_no
  LEFT JOIN offer ofr ON ofr.part_no = pp.part_no
  LEFT JOIN model_variants vnt ON vnt.part_no = pp.part_no

WHERE
  (
    (@PartNo::text IS NOT NULL AND pp.part_no = @PartNo)
    OR
    (
      (@PartNo::text IS NULL OR @PartNo::text = '')
    
      AND
      (
        /* Company + Model + Category (Company only to models)*/
       (
         (@CompanyName::text IS NULL OR @CompanyName = '' OR c.name ILIKE @CompanyName)
          AND NULLIF(@ModelName::text, '') IS NOT NULL
          AND m.name ILIKE @ModelName::text
          AND (
            @CategoryList::text[] IS NOT NULL
            AND cat.name = ANY(@CategoryList::text[])
          ) 
        )

        OR
        /* 2) EF: Unicode + UnicodeCategories (require BOTH) */
        (
            NULLIF(@Unicode::text, '') IS NOT NULL
            AND @UnicodeCategoryList::text[] IS NOT NULL
            AND EXISTS (
                SELECT 1
                FROM unnest(COALESCE(vnt.unicode, '{}')) AS u
                WHERE u ILIKE @Unicode::text
            )
            AND (
              @UnicodeCategoryList::text[] IS NOT NULL
              AND cat.name = ANY(@UnicodeCategoryList::text[])
            )
        )
          OR
        /* 3) A B' E' F': Category-only (no model, no unicode, no unicodeCategories) */
        (
          NULLIF(@ModelName::text, '') IS NULL
          AND NULLIF(@Unicode::text, '') IS NULL
          AND @UnicodeCategoryList::text[] IS NULL
          AND (
              @CategoryList::text[] IS NULL
              OR @CategoryList::text[] = '{}' 
              OR cat.name = ANY(@CategoryList::text[])
          )
          AND (@CompanyName::text IS NULL OR @CompanyName = '' OR c.name ILIKE @CompanyName)
        )
      )
      AND (@BrandList::text[] IS NULL OR  @BrandList::text[] = '{}'  OR b.name = ANY(@BrandList::text[]))
      AND (@GenList::text[] IS NULL OR  @GenList::text[] = '{}'  OR vnt.gen = ANY(@GenList::text[]))
      AND (@FuelList::text[] IS NULL OR  @FuelList::text[] = '{}'  OR EXISTS (
        SELECT 1 
        FROM unnest(vnt.fuel_types) AS ft
        WHERE ft = ANY(@FuelList::text[])
      ))
      AND (@MakeList::text[] IS NULL OR vnt.make = ANY(@MakeList::text[]))
    )
  );





-------------------------------------------------------------------------------------------------------------------------------
--Mutation
--------------------------------------------------------------------------------------------------------------------------
-- name: CreateProductParts :one 
INSERT INTO product_parts
  (id, company_id, model_id, brand_id, category_id, part_no, is_active, created_at, updated_at)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;


-- name: UpdateProductPartByID :one
UPDATE product_parts AS pp
SET
  company_id    = COALESCE(sqlc.narg('new_company_id'), pp.company_id),
  model_id      = COALESCE(sqlc.narg('new_model_id'), pp.model_id),
  brand_id      = COALESCE(sqlc.narg('new_brand_id'), pp.brand_id),
  category_id   = COALESCE(sqlc.narg('new_category_id'), pp.category_id),
  part_no       = COALESCE(sqlc.narg('new_part_no'), pp.part_no),
  is_active     = COALESCE(sqlc.narg('new_is_active'), pp.is_active),
  updated_at    = NOW()
WHERE pp.id = $1
RETURNING *;


-- name: DeleteProductPartByID :exec
DELETE FROM product_parts
WHERE id = $1;

-- name: DeleteProductPartByPartNo :exec
DELETE FROM product_parts
WHERE part_no = $1;

