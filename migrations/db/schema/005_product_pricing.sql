-- +goose Up

--ProductPartPricing--------------------------------------------------------------------

CREATE TABLE product_part_pricing (
    id UUID PRIMARY KEY,
    product_part_id UUID UNIQUE NOT NULL REFERENCES product_parts(id) ON DELETE CASCADE,

    basic_price FLOAT NOT NULL DEFAULT 0,
    freight FLOAT NOT NULL DEFAULT 0,
    gst FLOAT NOT NULL DEFAULT 0,

    --
    tax FLOAT NOT NULL DEFAULT 0,

    -- Ac Workshop (formerly retail)
    ac_workshop FLOAT NOT NULL DEFAULT 0,
    ac_workshop_per FLOAT NOT NULL DEFAULT 0,
    ac_workshop_amt FLOAT NOT NULL DEFAULT 0,

    -- Multibrand Workshop (formerly wholesale)
    multibrand_workshop FLOAT NOT NULL DEFAULT 0,
    multibrand_workshop_per FLOAT NOT NULL DEFAULT 0,
    multibrand_workshop_amt FLOAT NOT NULL DEFAULT 0,

    -- Auto Trader (formerly mrp)
    auto_trader FLOAT NOT NULL DEFAULT 0,
    auto_trader_per FLOAT NOT NULL DEFAULT 0,
    auto_trader_amt FLOAT NOT NULL DEFAULT 0,

    -- Outstation & Ac Trader
    ac_trader FLOAT NOT NULL DEFAULT 0,
    ac_trader_per FLOAT NOT NULL DEFAULT 0,
    ac_trader_amt FLOAT NOT NULL DEFAULT 0,

    --Not used
    outstation_class_a FLOAT NOT NULL DEFAULT 0,           
    outstation_note TEXT,                        

    minimum_purchase_quantity INT NOT NULL DEFAULT 0,
    mrp_temp FLOAT NOT NULL DEFAULT 0, -- temp storage of new mrp
    oem_mrp FLOAT NOT NULL DEFAULT 0,
    unit_measure VARCHAR(20),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);



-- +goose Down

DROP TABLE product_part_pricing;