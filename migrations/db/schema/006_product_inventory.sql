-- +goose Up

CREATE TABLE inventory (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  part_no TEXT NOT NULL UNIQUE,            -- Unique constraint added for part_no
  minimum_order_level INT NOT NULL,
  maximum_order_level INT NOT NULL,
  qty_in_stock INT NOT NULL,               -- Renamed to qty_in_stock
  location TEXT NOT NULL,
  is_flash BOOLEAN DEFAULT FALSE,          -- Higher priority for sale
  is_requested_for_supply BOOLEAN DEFAULT FALSE,  -- New column: marks if restock/supply is requested
  requested_date TIMESTAMPTZ,
  vendor_id UUID REFERENCES vendors(id)
);

-- +goose Down
DROP TABLE inventory;