-- name: GetTypesenseData :many
SELECT 
    p.id                       AS id,        -- Typesense primary key (or use part_no if you prefer)
    b.name                     AS brand,
    cat.name                   AS category,
    c.name                     AS company,
    COALESCE(mv.gen, '')       AS gen,
    'none'                     AS fuel_type, -- use single quotes, not "
    m.name                     AS model,
    p.part_no                  AS part_no
FROM product_parts p
LEFT JOIN companies  c   ON p.company_id  = c.id
LEFT JOIN models     m   ON p.model_id    = m.id
LEFT JOIN brands     b   ON p.brand_id    = b.id
LEFT JOIN categories cat ON p.category_id = cat.id
LEFT JOIN model_variants mv ON mv.part_no = p.part_no;

