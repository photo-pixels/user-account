-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL CHECK (LENGTH(name) <= 128) UNIQUE,
    description TEXT NOT NULL CHECK (LENGTH(description) <= 2096)
);

CREATE TABLE role_permission (
    permission_id UUID NOT NULL REFERENCES permission(id),
    role_id UUID NOT NULL REFERENCES role(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_role_permission ON role_permission(permission_id, role_id);

CREATE TABLE user_role (
    user_id UUID NOT NULL REFERENCES users(id),
    role_id UUID NOT NULL REFERENCES role(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_user_role ON user_role(user_id, role_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_role;
DROP TABLE role;
-- +goose StatementEnd
