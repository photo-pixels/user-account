-- +goose Up
-- +goose StatementBegin
CREATE TABLE permission (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL CHECK (LENGTH(name) <= 128) UNIQUE,
    description TEXT NOT NULL CHECK (LENGTH(description) <= 2096)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_permission;
DROP TABLE permission;
-- +goose StatementEnd
