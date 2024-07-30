-- +goose Up
-- +goose StatementBegin
CREATE TYPE auth_status AS ENUM (
    'NOT_ACTIVATED',
    'SENT_INVITE',
    'ACTIVATED',
    'BLOCKED'
);

CREATE TABLE auth (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    email TEXT NOT NULL CHECK (LENGTH(email) <= 1024) UNIQUE,
    password_hash BYTEA NOT NULL,
    status auth_status NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth;
DROP TYPE auth_status;
-- +goose StatementEnd
