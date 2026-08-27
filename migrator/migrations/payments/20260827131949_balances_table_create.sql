-- +goose Up
CREATE TABLE IF NOT EXISTS balances (
  user_id UUID PRIMARY KEY,
  bonuses BIGINT NOT NULL CHECK(bonuses >=0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS balances;