-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    firstname TEXT NOT NULL CHECK (LENGTH(firstname) <= 1024),
    surname TEXT NOT NULL CHECK (LENGTH(surname) <= 1024),
    patronymic TEXT CHECK (LENGTH(patronymic) <= 1024),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
