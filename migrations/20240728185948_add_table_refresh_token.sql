-- +goose Up
-- +goose StatementBegin
CREATE TYPE refresh_token_status AS ENUM (
    'ACTIVE',
    'REVOKED',
    'EXPIRED',
    'LOGOUT'
);

CREATE TABLE refresh_token (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    status refresh_token_status NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_token;
DROP TYPE refresh_token_status;
-- +goose StatementEnd
