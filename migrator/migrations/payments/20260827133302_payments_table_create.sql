-- +goose Up
CREATE TABLE IF NOT EXISTS payments (
  id UUID PRIMARY KEY,
  order_id UUID NOT NULL UNIQUE,
  user_id UUID NOT NULL,
  bonuses_amount BIGINT NOT NULL check(bonuses_amount >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS payments;