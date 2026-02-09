
-- name: GetSimilarPricingByPartNo :many
SELECT
    brnd.name as brand_name,
    pr.basic_price,
    pr.freight,
    pr.gst,
    pr.tax,
    pr.ac_workshop,
    pr.ac_workshop_per,
    pr.ac_workshop_amt,
    pr.multibrand_workshop,
    pr.multibrand_workshop_per,
    pr.multibrand_workshop_amt,
    pr.auto_trader,
    pr.auto_trader_per,
    pr.auto_trader_amt,
    pr.ac_trader,
    pr.ac_trader_per,
    pr.ac_trader_amt,
    pr.mrp_temp,
    pr.oem_mrp,
    pr.unit_measure,
    pr.minimum_purchase_quantity
FROM product_parts p
JOIN companies   c   ON c.id   = p.company_id
JOIN models      m   ON m.id   = p.model_id
JOIN categories  cat ON cat.id = p.category_id
JOIN brands brnd     ON brnd.id = p.brand_id
JOIN product_part_pricing pr ON pr.product_part_id = p.id
JOIN model_variants vnts ON vnts.part_no=p.part_no
WHERE 
    c.name = $1
    AND m.name = $2
    AND cat.name = $3
    AND vnts.type=$4
ORDER BY p.part_no;
