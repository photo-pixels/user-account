------------------------------------------------------------------------------------------------------------------------

-- name: GetUser :one
SELECT id, created_at, updated_at, firstname, surname, patronymic FROM users WHERE id = $1;

-- name: SaveUser :exec
INSERT INTO users (id, created_at, updated_at, firstname, surname, patronymic)
VALUES ($1, $2, $3, $4, $5, $6);

------------------------------------------------------------------------------------------------------------------------

-- name: EmailExists :one
SELECT count(1) as count FROM auth WHERE email = $1;

-- name: SavePersonAuth :exec
INSERT INTO auth (user_id, created_at, updated_at, email, password_hash, status)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetAuth :one
SELECT user_id, created_at, updated_at, email, password_hash, status FROM auth WHERE user_id = $1;

-- name: GetAuthByEmail :one
SELECT user_id, created_at, updated_at, email, password_hash, status FROM auth WHERE email = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: SaveRefreshToken :exec
INSERT INTO refresh_token (id, user_id, created_at, updated_at, status)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateRefreshTokenStatus :exec
UPDATE refresh_token
SET status = $1, updated_at = $2
WHERE id = $3;

-- name: GetLastActiveRefreshToken :one
SELECT id, user_id, created_at, updated_at, status FROM refresh_token
WHERE id=$1 and status='ACTIVE'
ORDER BY created_at DESC
LIMIT 1;

------------------------------------------------------------------------------------------------------------------------

-- name: SaveConfirmCode :exec
INSERT INTO code (code, user_id, created_at, updated_at, active, type)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetActiveConfirmCode :one
SELECT code, user_id, created_at, updated_at, active, type FROM code
WHERE code = $1 AND type = $2 and active = true;

------------------------------------------------------------------------------------------------------------------------

-- name: SavePermission :exec
INSERT INTO permission (id, created_at, updated_at, name, description)
VALUES ($1, $2, $3, $4, $5);

-- name: SaveRole :exec
INSERT INTO role (id, created_at, updated_at, name, description)
VALUES ($1, $2, $3, $4, $5);

-- name: AddPermissionToRole :exec
INSERT INTO role_permission (role_id, permission_id)
VALUES ($1, $2);

-- name: AddRoleToUser :exec
INSERT INTO user_role (user_id, role_id)
VALUES ($1, $2);

-- name: GetUserPermissions :many
select p.id as id, p.created_at as created_at, p.updated_at as updated_at, p.name as name, p.description as description from user_role u
join public.role_permission rp on u.role_id = rp.role_id
join public.permission p on rp.permission_id = p.id
where user_id=$1;

-- name: GetUserRole :one
select r.id as id, r.created_at as created_at, r.updated_at as updated_at, r.name as name, r.description as description from user_role u
join public.role r on u.role_id = r.id
where user_id=$1 and r.id = $2;

-- name: GetRolePermissions :many
SELECT id, p.created_at as created_at, updated_at, name, description FROM role_permission r
JOIN permission p ON p.id = r.permission_id
WHERE r.role_id = $1;

-- name: GetRolePermission :one
SELECT id, p.created_at as created_at, updated_at, name, description FROM role_permission r
JOIN permission p ON p.id = r.permission_id
WHERE r.role_id = $1 and r.permission_id = $2;

-- name: GetPermissionByName :one
SELECT id, created_at as created_at, updated_at, name, description FROM permission
WHERE name = $1;

-- name: GetRoleByName :one
SELECT id, created_at as created_at, updated_at, name, description FROM role
WHERE name = $1;

-- name: GetPermission :one
SELECT id, created_at as created_at, updated_at, name, description FROM permission
WHERE id = $1;

-- name: GetRole :one
SELECT id, created_at as created_at, updated_at, name, description FROM role
WHERE id = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: GetTokens :many
SELECT id, user_id, title, token, token_type, created_at, updated_at, expired_at FROM token
WHERE user_id=$1
ORDER BY created_at;

-- name: SaveToken :exec
INSERT INTO token (id, user_id, title, token, token_type, created_at, updated_at, expired_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: DeleteToken :one
DELETE FROM token
WHERE id=$1 and user_id=$2
RETURNING id;

-- name: GetToken :one
SELECT id, user_id, title, token, token_type, created_at, updated_at, expired_at FROM token
WHERE token=$1
LIMIT 1;