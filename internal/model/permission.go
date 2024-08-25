package model

import (
	"github.com/google/uuid"
	"github.com/photo-pixels/platform/basemodel"
)

// Permission подель пермисии
type Permission struct {
	basemodel.Base
	ID          uuid.UUID
	Name        string
	Description string
}

// Role роль
type Role struct {
	basemodel.Base
	ID          uuid.UUID
	Name        string
	Description string
}

// GetPermissions параметры получения списка пермисий
type GetPermissions struct {
	PermissionNameIn []string
	PermissionIDIn   []uuid.UUID
	UserIDIn         []uuid.UUID
	RoleIDIn         []uuid.UUID
}

// GetRoles параметры получения списка ролей
type GetRoles struct {
	RoleIDIn   []uuid.UUID
	RoleNameIn []string
	UserIDIn   []uuid.UUID
}
