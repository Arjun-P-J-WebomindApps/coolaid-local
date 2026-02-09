------------------------------------------------------------------------------------------------------------------------------
--Query
------------------------------------------------------------------------------------------------------------------------------

-- name: GetModelVariantByPartNo :one
SELECT * FROM model_variants WHERE part_no=$1;

-- name: GetModelVariantDownloadDetails :many
SELECT
    pp.part_no                       AS "Part No",
    c.name                           AS "Company",
    mdl.name                         AS "Model",
    b.name                           AS "Brand",
    cat.name                         AS "Category",

    m.type                           AS "Type",
    m.make                           AS "Make",
    m.gen                            AS "Generation",
    m.year_start                     AS "Year Start",
    m.year_end                       AS "Year End",
    m.engine_cc                      AS "Engine CC",
    m.transmission_type              AS "Transmission",
    m.placement                      AS "Placement",
    m.hsn_code                       AS "HSN Code",
    m.unicode                        AS "Unicode",
    m.description                    AS "Description",
    m.additional_info                AS "Additional Info",
    m.fuel_types                     AS "Fuel Types",
    m.platform_codes                 AS "Platform Codes",
    m.oem_ids                        AS "OEM IDs",
    m.vendor_id                      AS "Vendor IDs"

FROM model_variants m
JOIN product_parts pp ON pp.part_no = m.part_no
JOIN companies     c  ON c.id       = pp.company_id
JOIN models        mdl ON mdl.id    = pp.model_id
JOIN brands        b   ON b.id      = pp.brand_id
JOIN categories    cat ON cat.id    = pp.category_id;




------------------------------------------------------------------------------------------------------------------------------
--Mutation
-------------------------------------------------------------------------------------------------------------------------------
-- name: CreateModelVariant :one
INSERT INTO model_variants (
    id,
    part_no,
    type,
    fuel_types,            -- Array of fuel types
    gen,
    engine_cc,
    transmission_type,
    platform_codes,        -- Array of platform codes
    additional_info,
    image_1_link,
    image_2_link,
    image_3_link,
    image_4_link,
    oem_ids,
    make,
    vendor_id,
    hsn_code,
    unicode,
    description,
    placement,             -- Default placement value
    year_start,            -- Newly added column
    year_end               -- Newly added column
)
VALUES (
    $1,  -- id
    $2,  -- part_no
    $3,  -- type
    $4,  -- fuel_types (array of fuel types)
    $5,  -- gen
    $6,  -- engine_cc
    $7,  -- transmission_type
    $8,  -- platform_codes (array of platform codes)
    $9,  -- additional_info
    $10, -- image_1_link
    $11, -- image_2_link
    $12, -- image_3_link
    $13, -- image_4_link
    $14, -- oem_ids
    $15, -- make
    $16, -- vendor_id
    $17, -- hsn_code
    $18, -- unicode
    $19, -- description
    $20, -- placement
    $21, -- year_start
    $22  -- year_end
)
RETURNING *;



-- name: UpdateModelVariant :one
UPDATE model_variants AS mv SET
    type = COALESCE($2, mv.type),                  -- Update type if provided
    fuel_types = COALESCE($3, mv.fuel_types),      -- Update fuel_types array if provided
    gen = COALESCE($4, mv.gen),                    -- Update gen if provided
    engine_cc = COALESCE($5, mv.engine_cc),        -- Update engine_cc if provided
    transmission_type = COALESCE($6, mv.transmission_type), -- Update transmission_type if provided
    platform_codes = COALESCE($7, mv.platform_codes), -- Update platform_codes array if provided
    additional_info = COALESCE($8, mv.additional_info), -- Update additional_info if provided
    image_1_link = COALESCE($9, mv.image_1_link),  -- Update image_1_link if provided
    image_2_link = COALESCE($10, mv.image_2_link), -- Update image_2_link if provided
    image_3_link = COALESCE($11, mv.image_3_link), -- Update image_3_link if provided
    image_4_link = COALESCE($12, mv.image_4_link), -- Update image_4_link if provided
    oem_ids = COALESCE($13, mv.oem_ids),           -- Update oem_ids array if provided
    make = COALESCE($14, mv.make),                 -- Update make if provided
    vendor_id = COALESCE($15, mv.vendor_id),       -- Update vendor_id array if provided
    hsn_code = COALESCE($16, mv.hsn_code),         -- Update hsn_code if provided
    unicode = COALESCE($17, mv.unicode),           -- Update unicode array if provided
    description = COALESCE($18, mv.description),   -- Update description if provided
    placement = COALESCE($19, mv.placement),       -- Update placement if provided
    year_start = COALESCE($20, mv.year_start),     -- Update year_start if provided
    year_end = COALESCE($21, mv.year_end)          -- Update year_end if provided
WHERE mv.part_no = $1                              -- Ensure update is done for the correct part_no
RETURNING *;                                       -- Return the updated record


-- name: DeleteModelVariantByPartNo :execresult
DELETE FROM model_variants
WHERE part_no = $1;

-- name: BulkCreateModelVariants :exec
WITH data AS (
    SELECT
        unnest(@part_nos::text[])          AS part_no,
        unnest(@gens::text[])              AS gen,
        unnest(@engine_ccs::float8[])        AS engine_cc,
        unnest(@transmissions::text[])     AS transmission_type_raw,
        unnest(@additional_infos::text[])  AS additional_info,
        unnest(@image1_links::text[])      AS image_1_link,
        unnest(@image2_links::text[])      AS image_2_link,
        unnest(@image3_links::text[])      AS image_3_link,
        unnest(@image4_links::text[])      AS image_4_link,
        unnest(@hsn_codes::text[])         AS hsn_code,
        unnest(@descriptions::text[])      AS description,
        unnest(@oem_ids::text[])           AS oem_ids_raw,
        unnest(@makes::text[])             AS make,
        unnest(@unicode_vals::text[])      AS unicode_raw,
        unnest(@vendor_ids::text[])        AS vendor_id_raw,
        unnest(@platform_codes::text[])    AS platform_codes_raw,
        unnest(@fuel_types::text[])        AS fuel_types_raw,
        unnest(@placements::text[])        AS placement,
        unnest(@types::text[])             AS type_val,
        unnest(@year_starts::int[])        AS year_start,
        unnest(@year_ends::int[])          AS year_end
)
INSERT INTO model_variants (
    id,
    part_no,
    gen,
    engine_cc,
    transmission_type,
    additional_info,
    image_1_link,
    image_2_link,
    image_3_link,
    image_4_link,
    hsn_code,
    description,
    oem_ids,
    make,
    unicode,
    vendor_id,
    platform_codes,
    fuel_types,
    placement,
    type,
    year_start,
    year_end
)
SELECT
    gen_random_uuid(),
    d.part_no,
    d.gen,
    d.engine_cc,

    -- transmission_type text[] from CSV string (e.g. "MT,AT")
    COALESCE(
        string_to_array(NULLIF(d.transmission_type_raw, ''), ','),
        ARRAY[]::text[]
    ) AS transmission_type,

    d.additional_info,
    d.image_1_link,
    d.image_2_link,
    d.image_3_link,
    d.image_4_link,
    d.hsn_code,
    d.description,

    -- oem_ids text[]
    COALESCE(
        string_to_array(NULLIF(d.oem_ids_raw, ''), ','),
        ARRAY[]::text[]
    ) AS oem_ids,

    d.make,

    -- unicode text[]
    COALESCE(
        string_to_array(NULLIF(d.unicode_raw, ''), ','),
        ARRAY[]::text[]
    ) AS unicode,

    -- vendor_id text[]
    COALESCE(
        string_to_array(NULLIF(d.vendor_id_raw, ''), ','),
        ARRAY[]::text[]
    ) AS vendor_id,


    -- platform_codes text[]
    COALESCE(
        string_to_array(NULLIF(d.platform_codes_raw, ''), ','),
        ARRAY[]::text[]
    ) AS platform_codes,

    -- fuel_types text[] ("PETROL", "DIESEL", "PETROL,DIESEL", "NONE")
    COALESCE(
        string_to_array(NULLIF(d.fuel_types_raw, ''), ','),
        ARRAY[]::text[]
    ) AS fuel_types,

    d.placement,
    d.type_val,
    d.year_start,
    d.year_end
FROM data d
ON CONFLICT (part_no) DO NOTHING;


-- name: BulkUpdateModelVariantsSelective :exec
WITH data AS (
    SELECT
        unnest(@part_nos::text[])          AS part_no,
        unnest(@gens::text[])              AS gen,
        unnest(@engine_ccs::float8[])      AS engine_cc,
        unnest(@transmissions::text[])     AS transmission_type_raw,
        unnest(@additional_infos::text[])  AS additional_info,
        unnest(@hsn_codes::text[])         AS hsn_code,
        unnest(@descriptions::text[])      AS description,
        unnest(@oem_ids::text[])           AS oem_ids_raw,
        unnest(@makes::text[])             AS make,
        unnest(@unicode_vals::text[])      AS unicode_raw,
        unnest(@vendor_ids::text[])        AS vendor_id_raw,
        unnest(@platform_codes::text[])    AS platform_codes_raw,
        unnest(@fuel_types::text[])        AS fuel_types_raw,
        unnest(@placements::text[])        AS placement,
        unnest(@types::text[])             AS type_val,
        unnest(@year_starts::int[])        AS year_start,
        unnest(@year_ends::int[])          AS year_end
)
UPDATE model_variants mv
SET
    gen = COALESCE(NULLIF(d.gen, ''), mv.gen),
    engine_cc = COALESCE(d.engine_cc, mv.engine_cc),

    transmission_type = COALESCE(
        string_to_array(NULLIF(d.transmission_type_raw, ''), ','),
        mv.transmission_type
    ),

    additional_info = COALESCE(NULLIF(d.additional_info, ''), mv.additional_info),

    hsn_code = COALESCE(NULLIF(d.hsn_code, ''), mv.hsn_code),
    description = COALESCE(NULLIF(d.description, ''), mv.description),

    oem_ids = COALESCE(
        string_to_array(NULLIF(d.oem_ids_raw, ''), ','),
        mv.oem_ids
    ),

    make = COALESCE(NULLIF(d.make, ''), mv.make),

    unicode = COALESCE(
        string_to_array(NULLIF(d.unicode_raw, ''), ','),
        mv.unicode
    ),

    vendor_id = COALESCE(
        string_to_array(NULLIF(d.vendor_id_raw, ''), ','),
        mv.vendor_id
    ),

    platform_codes = COALESCE(
        string_to_array(NULLIF(d.platform_codes_raw, ''), ','),
        mv.platform_codes
    ),

    fuel_types = COALESCE(
        string_to_array(NULLIF(d.fuel_types_raw, ''), ','),
        mv.fuel_types
    ),

    placement = COALESCE(NULLIF(d.placement, ''), mv.placement),
    type = COALESCE(NULLIF(d.type_val, ''), mv.type),
    year_start = COALESCE(d.year_start, mv.year_start),
    year_end = COALESCE(d.year_end, mv.year_end)

FROM data d
WHERE mv.part_no = d.part_no;
