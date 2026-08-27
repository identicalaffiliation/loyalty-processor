-- +goose Up
CREATE TABLE IF NOT EXISTS outbox (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    published_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_outbox_event_unpublished
    ON outbox(published_at) WHERE published_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS idx_outbox_event_unpublished;
DROP TABLE IF EXISTS outbox;
