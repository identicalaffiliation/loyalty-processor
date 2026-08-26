-- +goose Up
CREATE TYPE outbox_status AS ENUM ('created', 'published');

CREATE TABLE IF NOT EXISTS outbox (
  id UUID PRIMARY KEY,
  order_id UUID NOT NULL,
  status outbox_status NOT NULL DEFAULT 'created',
  payload JSONB NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  published_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_outbox_event_unpublished
    ON outbox(published_at) WHERE published_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS idx_outbox_event_unpublished;
DROP TABLE IF EXISTS outbox;
DROP TYPE IF EXISTS outbox_status;