-- name: FilterVariantTypes :many
SELECT DISTINCT
  m2."type"
FROM product_parts p
JOIN model_variants m2 ON m2.part_no = p.part_no
JOIN companies c       ON c.id = p.company_id
JOIN models m          ON m.id = p.model_id
JOIN categories cat    ON cat.id = p.category_id
WHERE (NULLIF($1::text, '') IS NULL OR UPPER(c.name) = UPPER($1))
  AND (NULLIF($2::text, '') IS NULL OR UPPER(m.name) = UPPER($2))
  AND (NULLIF($3::text, '') IS NULL OR UPPER(cat.name) = UPPER($3))
  AND (NULLIF($4::text, '') IS NULL OR UPPER(m2."type") LIKE '%' || UPPER($4) || '%')
ORDER BY m2."type";
