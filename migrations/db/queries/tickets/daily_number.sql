--------------------------------------------------------------------------------------------------------------------------------
--Mutation (Daily Number)
--------------------------------------------------------------------------------------------------------------------------------

-- name: GetNextDailyNumber :one
WITH upsert AS (
  INSERT INTO daily_ticket_counters (ticket_date, last_number)
  VALUES (CURRENT_DATE, 1)
  ON CONFLICT (ticket_date)
  DO UPDATE SET last_number = daily_ticket_counters.last_number + 1
  RETURNING last_number
)
SELECT last_number FROM upsert;
