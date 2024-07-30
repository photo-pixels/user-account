// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addPermissionToRole = `-- name: AddPermissionToRole :exec
INSERT INTO role_permission (role_id, permission_id)
VALUES ($1, $2)
`

type AddPermissionToRoleParams struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
}

func (q *Queries) AddPermissionToRole(ctx context.Context, arg AddPermissionToRoleParams) error {
	_, err := q.db.Exec(ctx, addPermissionToRole, arg.RoleID, arg.PermissionID)
	return err
}

const addRoleToUser = `-- name: AddRoleToUser :exec
INSERT INTO user_role (user_id, role_id)
VALUES ($1, $2)
`

type AddRoleToUserParams struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (q *Queries) AddRoleToUser(ctx context.Context, arg AddRoleToUserParams) error {
	_, err := q.db.Exec(ctx, addRoleToUser, arg.UserID, arg.RoleID)
	return err
}

const emailExists = `-- name: EmailExists :one

SELECT count(1) as count FROM auth WHERE email = $1
`

// ----------------------------------------------------------------------------------------------------------------------
func (q *Queries) EmailExists(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRow(ctx, emailExists, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getActiveConfirmCode = `-- name: GetActiveConfirmCode :one
SELECT code, user_id, created_at, updated_at, active, type FROM code
WHERE code = $1 AND type = $2 and active = true
`

type GetActiveConfirmCodeParams struct {
	Code string
	Type CodeType
}

func (q *Queries) GetActiveConfirmCode(ctx context.Context, arg GetActiveConfirmCodeParams) (Code, error) {
	row := q.db.QueryRow(ctx, getActiveConfirmCode, arg.Code, arg.Type)
	var i Code
	err := row.Scan(
		&i.Code,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Active,
		&i.Type,
	)
	return i, err
}

const getAuth = `-- name: GetAuth :one
SELECT user_id, created_at, updated_at, email, password_hash, status FROM auth WHERE user_id = $1
`

func (q *Queries) GetAuth(ctx context.Context, userID uuid.UUID) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuth, userID)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
		&i.Status,
	)
	return i, err
}

const getAuthByEmail = `-- name: GetAuthByEmail :one
SELECT user_id, created_at, updated_at, email, password_hash, status FROM auth WHERE email = $1
`

func (q *Queries) GetAuthByEmail(ctx context.Context, email string) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuthByEmail, email)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.PasswordHash,
		&i.Status,
	)
	return i, err
}

const getLastActiveRefreshToken = `-- name: GetLastActiveRefreshToken :one
SELECT id, user_id, created_at, updated_at, status FROM refresh_token
WHERE id=$1 and status='ACTIVE'
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetLastActiveRefreshToken(ctx context.Context, id uuid.UUID) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, getLastActiveRefreshToken, id)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return i, err
}

const getPermission = `-- name: GetPermission :one
SELECT id, created_at as created_at, updated_at, name, description FROM permission
WHERE id = $1
`

func (q *Queries) GetPermission(ctx context.Context, id uuid.UUID) (Permission, error) {
	row := q.db.QueryRow(ctx, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getPermissionByName = `-- name: GetPermissionByName :one
SELECT id, created_at as created_at, updated_at, name, description FROM permission
WHERE name = $1
`

func (q *Queries) GetPermissionByName(ctx context.Context, name string) (Permission, error) {
	row := q.db.QueryRow(ctx, getPermissionByName, name)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getRole = `-- name: GetRole :one
SELECT id, created_at as created_at, updated_at, name, description FROM role
WHERE id = $1
`

func (q *Queries) GetRole(ctx context.Context, id uuid.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT id, created_at as created_at, updated_at, name, description FROM role
WHERE name = $1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getRolePermission = `-- name: GetRolePermission :one
SELECT id, p.created_at as created_at, updated_at, name, description FROM role_permission r
JOIN permission p ON p.id = r.permission_id
WHERE r.role_id = $1 and r.permission_id = $2
`

type GetRolePermissionParams struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
}

func (q *Queries) GetRolePermission(ctx context.Context, arg GetRolePermissionParams) (Permission, error) {
	row := q.db.QueryRow(ctx, getRolePermission, arg.RoleID, arg.PermissionID)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const getRolePermissions = `-- name: GetRolePermissions :many
SELECT id, p.created_at as created_at, updated_at, name, description FROM role_permission r
JOIN permission p ON p.id = r.permission_id
WHERE r.role_id = $1
`

func (q *Queries) GetRolePermissions(ctx context.Context, roleID uuid.UUID) ([]Permission, error) {
	rows, err := q.db.Query(ctx, getRolePermissions, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one

SELECT id, created_at, updated_at, firstname, surname, patronymic FROM users WHERE id = $1
`

type GetUserRow struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Firstname  string
	Surname    string
	Patronymic *string
}

// ----------------------------------------------------------------------------------------------------------------------
func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Firstname,
		&i.Surname,
		&i.Patronymic,
	)
	return i, err
}

const getUserPermissions = `-- name: GetUserPermissions :many
select p.id as id, p.created_at as created_at, p.updated_at as updated_at, p.name as name, p.description as description from user_role u
join public.role_permission rp on u.role_id = rp.role_id
join public.permission p on rp.permission_id = p.id
where user_id=$1
`

func (q *Queries) GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]Permission, error) {
	rows, err := q.db.Query(ctx, getUserPermissions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserRole = `-- name: GetUserRole :one
select r.id as id, r.created_at as created_at, r.updated_at as updated_at, r.name as name, r.description as description from user_role u
join public.role r on u.role_id = r.id
where user_id=$1 and r.id = $2
`

type GetUserRoleParams struct {
	UserID uuid.UUID
	ID     uuid.UUID
}

func (q *Queries) GetUserRole(ctx context.Context, arg GetUserRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, getUserRole, arg.UserID, arg.ID)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const saveConfirmCode = `-- name: SaveConfirmCode :exec

INSERT INTO code (code, user_id, created_at, updated_at, active, type)
VALUES ($1, $2, $3, $4, $5, $6)
`

type SaveConfirmCodeParams struct {
	Code      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
	Type      CodeType
}

// ----------------------------------------------------------------------------------------------------------------------
func (q *Queries) SaveConfirmCode(ctx context.Context, arg SaveConfirmCodeParams) error {
	_, err := q.db.Exec(ctx, saveConfirmCode,
		arg.Code,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Active,
		arg.Type,
	)
	return err
}

const savePermission = `-- name: SavePermission :exec

INSERT INTO permission (id, created_at, updated_at, name, description)
VALUES ($1, $2, $3, $4, $5)
`

type SavePermissionParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

// ----------------------------------------------------------------------------------------------------------------------
func (q *Queries) SavePermission(ctx context.Context, arg SavePermissionParams) error {
	_, err := q.db.Exec(ctx, savePermission,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
	)
	return err
}

const savePersonAuth = `-- name: SavePersonAuth :exec
INSERT INTO auth (user_id, created_at, updated_at, email, password_hash, status)
VALUES ($1, $2, $3, $4, $5, $6)
`

type SavePersonAuthParams struct {
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash []byte
	Status       AuthStatus
}

func (q *Queries) SavePersonAuth(ctx context.Context, arg SavePersonAuthParams) error {
	_, err := q.db.Exec(ctx, savePersonAuth,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Email,
		arg.PasswordHash,
		arg.Status,
	)
	return err
}

const saveRefreshToken = `-- name: SaveRefreshToken :exec

INSERT INTO refresh_token (id, user_id, created_at, updated_at, status)
VALUES ($1, $2, $3, $4, $5)
`

type SaveRefreshTokenParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    RefreshTokenStatus
}

// ----------------------------------------------------------------------------------------------------------------------
func (q *Queries) SaveRefreshToken(ctx context.Context, arg SaveRefreshTokenParams) error {
	_, err := q.db.Exec(ctx, saveRefreshToken,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Status,
	)
	return err
}

const saveRole = `-- name: SaveRole :exec
INSERT INTO role (id, created_at, updated_at, name, description)
VALUES ($1, $2, $3, $4, $5)
`

type SaveRoleParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

func (q *Queries) SaveRole(ctx context.Context, arg SaveRoleParams) error {
	_, err := q.db.Exec(ctx, saveRole,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
	)
	return err
}

const saveUser = `-- name: SaveUser :exec
INSERT INTO users (id, created_at, updated_at, firstname, surname, patronymic)
VALUES ($1, $2, $3, $4, $5, $6)
`

type SaveUserParams struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Firstname  string
	Surname    string
	Patronymic *string
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) error {
	_, err := q.db.Exec(ctx, saveUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Firstname,
		arg.Surname,
		arg.Patronymic,
	)
	return err
}

const updateRefreshTokenStatus = `-- name: UpdateRefreshTokenStatus :exec
UPDATE refresh_token
SET status = $1, updated_at = $2
WHERE id = $3
`

type UpdateRefreshTokenStatusParams struct {
	Status    RefreshTokenStatus
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateRefreshTokenStatus(ctx context.Context, arg UpdateRefreshTokenStatusParams) error {
	_, err := q.db.Exec(ctx, updateRefreshTokenStatus, arg.Status, arg.UpdatedAt, arg.ID)
	return err
}
