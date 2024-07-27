package form

import "github.com/google/uuid"

// GetUserPermissions форма получения пермисий пользователя
type GetUserPermissions struct {
	UserID uuid.UUID `validate:"required,uuid"`
}

// AddPermissionToRole добавление новой пермисии к роли
type AddPermissionToRole struct {
	PermissionID uuid.UUID `validate:"required,uuid"`
	RoleID       uuid.UUID `validate:"required,uuid"`
}

// AddRoleToUser добавление роли к пользователю
type AddRoleToUser struct {
	UserID uuid.UUID `validate:"required,uuid"`
	RoleID uuid.UUID `validate:"required,uuid"`
}

// CreateRole создание новой роли
type CreateRole struct {
	Name        string `validate:"required,min=3,max=128"`
	Description string `validate:"required,min=3,max=1024"`
}

// CreatePermission создание новой пермисии
type CreatePermission struct {
	Name        string `validate:"required,min=3,max=128"`
	Description string `validate:"required,min=3,max=1024"`
}
