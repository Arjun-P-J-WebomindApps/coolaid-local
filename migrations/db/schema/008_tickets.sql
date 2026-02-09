-- +goose Up

CREATE TABLE tickets (
    id UUID PRIMARY KEY,
    ticket_date TIMESTAMP NOT NULL DEFAULT now(),
    customer_id UUID NOT NULL,
    daily_number TEXT NOT NULL,
    paf FLOAT NOT NULL,

    images_included BOOLEAN DEFAULT false,

    status TEXT NOT NULL CHECK (
        status IN ('pending', 'completed', 'cancelled')
    ),

    transport_mode TEXT NOT NULL DEFAULT 'porter'
        CHECK (transport_mode IN ('porter', 'self_pickup')),

    inspection_images TEXT[] NOT NULL DEFAULT ARRAY[]::TEXT[]
        CHECK (inspection_images <@ ARRAY['send', 'check', 'pending']::TEXT[]),

    confirm_order_date TIMESTAMP DEFAULT NULL,
    check_for_damage BOOLEAN DEFAULT false,
    remove_all_labels BOOLEAN DEFAULT false,
    exclude_documents BOOLEAN DEFAULT false,
    urgent_requirement BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE ticket_items (
    id UUID PRIMARY KEY,
    ticket_id UUID NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    part_no TEXT NOT NULL,
    model TEXT NOT NULL,
    quantity INT NOT NULL,
    unit_price FLOAT NOT NULL DEFAULT 0,
    is_selected BOOLEAN DEFAULT true
);

CREATE TABLE daily_ticket_counters (
    ticket_date DATE PRIMARY KEY,
    last_number INT NOT NULL
);


-- CREATE OR REPLACE FUNCTION set_daily_ticket_number()
-- RETURNS TRIGGER AS $$
-- DECLARE
--     ticket_day  DATE;
--     next_number INT;
-- BEGIN
--     -- Ensure we always have a timestamp
--     NEW.ticket_date := COALESCE(NEW.ticket_date, now());
--     ticket_day := NEW.ticket_date::date;

--     -- Serialize numbering per day to avoid duplicate numbers under concurrency
--     PERFORM pg_advisory_xact_lock(hashtext(ticket_day::text));

--     -- Count for the calendar day (not exact timestamp)
--     SELECT COALESCE(MAX(daily_number), 0) + 1
--       INTO next_number
--       FROM tickets
--      WHERE ticket_date::date = ticket_day;

--     NEW.daily_number := next_number;
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- +goose Down

DROP TRIGGER IF EXISTS trg_set_daily_ticket_number ON tickets;
-- CREATE TRIGGER trg_set_daily_ticket_number
-- BEFORE INSERT ON tickets
-- FOR EACH ROW
-- EXECUTE FUNCTION set_daily_ticket_number();
