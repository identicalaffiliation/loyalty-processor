-- +goose Up
CREATE TYPE orders_status AS ENUM ('created', 'paid', 'failed');

CREATE TABLE IF NOT EXISTS orders (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  product_id UUID NOT NULL,
  bonus_amount BIGINT NOT NULL CHECK (bonus_amount > 0),
  status orders_status NOT NULL DEFAULT 'created',
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_orders_user_created ON orders(user_id, created_at DESC);

-- +goose Down
DROP INDEX IF EXISTS idx_orders_user_created;
DROP TABLE IF EXISTS orders;
DROP TYPE IF EXISTS orders_status;