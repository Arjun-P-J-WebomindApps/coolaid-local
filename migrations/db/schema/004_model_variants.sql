-- +goose Up


--Model Variants--------------------------------------------------------------------------

CREATE TABLE model_variants (
    id UUID PRIMARY KEY,

    -- Specs
    part_no TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL DEFAULT 'TYPE-NONE',
    gen TEXT,
    fuel_types TEXT[] DEFAULT '{}', 
    hsn_code varchar(20) DEFAULT '',

    engine_cc FLOAT,
    transmission_type TEXT[] DEFAULT '{}' ,
    platform_codes TEXT[] DEFAULT '{}',
    placement TEXT NOT NULL DEFAULT 'NOT APPLICABLE', -- Default value set to 'NOT APPLICABLE'

    -- Four display images
    image_1_link TEXT,
    image_2_link TEXT,
    image_3_link TEXT,
    image_4_link TEXT,

    make TEXT NOT NULL,
    oem_ids TEXT[] NOT NULL DEFAULT '{}',
    vendor_id TEXT[] NOT NULL DEFAULT '{}'::TEXT[],
    unicode TEXT[] NOT NULL DEFAULT '{}',
  
    description TEXT,
    additional_info TEXT,    -- now plain text

    year_start INTEGER,   -- e.g. 1983
    year_end   INTEGER,   -- e.g. 1989

    CONSTRAINT model_variants_make_chk CHECK (make IN ('OEM', 'Aftermarket')),
    CONSTRAINT year_range_chk CHECK (
        year_start IS NULL OR year_end IS NULL OR year_start <= year_end
    )
);


-- +goose Down

DROP TABLE model_variants;
