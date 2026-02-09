-- +goose Up
CREATE TABLE vendors (
  id                      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  company_name            TEXT NOT NULL UNIQUE,
  created_at              TIMESTAMP NOT NULL DEFAULT now(),
  updated_at              TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE vendor_contacts (
  id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  vendor_id      UUID NOT NULL REFERENCES vendors(id) ON DELETE CASCADE,
  contact_person TEXT NOT NULL,
  mobile_no      VARCHAR(15) NOT NULL DEFAULT '',
  email_id       TEXT NOT NULL DEFAULT '',
  created_at     TIMESTAMP NOT NULL DEFAULT now(),
  updated_at     TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE vendor_listing (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  product_part_no TEXT NOT NULL REFERENCES product_parts(part_no) ON UPDATE CASCADE ON DELETE RESTRICT,
  vendor_name     TEXT NOT NULL,
  vendor_part_no  TEXT NOT NULL,
  vendor_mrp      FLOAT NOT NULL DEFAULT 0,
  created_at      TIMESTAMP NOT NULL DEFAULT now(),
  updated_at      TIMESTAMP NOT NULL DEFAULT now()
);

-- Prevent duplicates for the same product + vendor/part combo
CREATE UNIQUE INDEX vendors_part_vendor_key
  ON vendor_listing (part_no, vendor_name, vendor_part_no);

-- +goose Down
DROP TABLE vendors;
DROP TABLE vendor_contacts;
DROP TABLE vendor_listing;
