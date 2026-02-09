-- +goose Up
CREATE TABLE orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  order_listing_ids TEXT[],  -- Array of order listing IDs
  vendor_id UUID REFERENCES vendors(id),  -- Foreign key referencing the vendors table
  user_id UUID,  -- Assuming you have a users table, reference the user_id here
  is_requested_via_email BOOLEAN DEFAULT FALSE,
  is_requested_via_whatsapp BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders_listing (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  part_no TEXT NOT NULL,  -- Part number for the product
  quantity INT NOT NULL,  -- Quantity of the product in the order
  created_at TIMESTAMP DEFAULT NOW()  -- Timestamp for when the record was created
);


-- +goose Down
DROP TABLE orders_listing;
DROP TABLE orders;
