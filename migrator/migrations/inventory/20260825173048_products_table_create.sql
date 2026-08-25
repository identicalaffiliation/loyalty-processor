-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    stock INTEGER NOT NULL CHECK (stock > 0),
    created_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose Down
DROP TABLE IF EXISTS products;