-- +goose Up
-- +goose StatementBegin
CREATE TABLE token (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    title TEXT NOT NULL CHECK (LENGTH(title) <= 128),
    token TEXT NOT NULL UNIQUE CHECK (LENGTH(token) <= 64),
    token_type TEXT NOT NULL CHECK (LENGTH(token_type) <= 128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    expired_at TIMESTAMPTZ
);

CREATE INDEX idx_token_user_id ON token(user_id);
CREATE INDEX idx_token_token ON token(token);
CREATE UNIQUE INDEX idx_title_user_token ON token(user_id, title, token_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE token;
-- +goose StatementEnd
