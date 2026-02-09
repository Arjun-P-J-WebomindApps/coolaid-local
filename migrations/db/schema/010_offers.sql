-- +goose Up
CREATE TABLE offer (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  part_no TEXT NOT NULL UNIQUE,  -- Unique constraint added for part_no
  is_offer_active BOOLEAN DEFAULT FALSE,
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  ac_trader TEXT[],  -- Array of strings for AcTrader
  multi_brand TEXT[],  -- Array of strings for Mulitbrand
  autotrader TEXT[],  -- Array of strings for autotrader
  ac_workshop TEXT[]  -- Array of strings for acworkshop
);


-- +goose Down
DROP TABLE offer;