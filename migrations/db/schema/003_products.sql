-- +goose Up


-- Products -------------------------------------------------------------------------

CREATE TABLE product_parts(
    id UUID PRIMARY KEY,
    company_id UUID NOT NULL REFERENCES companies(id),
    model_id UUID NOT NULL REFERENCES models(id),
    brand_id UUID NOT  NULL REFERENCES brands(id),
    category_id UUID NOT NULL REFERENCES categories(id),
    part_no TEXT NOT NULL UNIQUE,
    oem_no TEXT,-- TODO: Not required
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down

DROP TABLE product_parts;
