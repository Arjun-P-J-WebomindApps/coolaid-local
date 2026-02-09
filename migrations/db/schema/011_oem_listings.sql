-- +goose Up
CREATE TABLE oem_listings (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  part_no     TEXT NOT NULL,
  oem_number  TEXT NOT NULL,
  price       FLOAT NOT NULL DEFAULT 0,
  created_at  TIMESTAMP NOT NULL DEFAULT now(),
  updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

-- Optional: prevent duplicate OEM numbers
CREATE UNIQUE INDEX oem_listings_number_key
  ON oem_listings (oem_number);

-- +goose Down
DROP TABLE oem_listings;
