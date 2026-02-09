---------------------------------------------------------------------
-- 📘 QUERIES (Read Operations)
-- These statements only READ data from the database.
---------------------------------------------------------------------

-- name: GetOfferByPartNo :one
SELECT *
FROM offer
WHERE part_no = $1;



---------------------------------------------------------------------
-- ⚙️ MUTATIONS (Write Operations)
-- These statements CREATE, UPDATE, or DELETE data in the database.
---------------------------------------------------------------------

-- name: CreateOffer :one
INSERT INTO offer (
    id,
    part_no,
    is_offer_active,
    start_date,
    end_date,
    ac_trader,
    multi_brand,
    autotrader,
    ac_workshop
)
VALUES (
    $1,  -- id
    $2,  -- part_no
    $3,  -- is_offer_active
    $4,  -- start_date
    $5,  -- end_date
    $6,  -- ac_trader (ARRAY)
    $7,  -- multi_brand (ARRAY)
    $8,  -- autotrader (ARRAY)
    $9   -- ac_workshop (ARRAY)
)
RETURNING *;

-- name: UpdateOfferByPartNo :one
UPDATE offer SET
  is_offer_active = COALESCE($2, is_offer_active),
  start_date      = COALESCE($3, start_date),
  end_date        = COALESCE($4, end_date),
  ac_trader       = COALESCE($5, ac_trader),
  multi_brand     = COALESCE($6, multi_brand),
  autotrader      = COALESCE($7, autotrader),
  ac_workshop     = COALESCE($8, ac_workshop)
WHERE part_no = $1
RETURNING *;

-- name: DeleteOfferByPartNo :exec
DELETE FROM offer
WHERE part_no = $1;
