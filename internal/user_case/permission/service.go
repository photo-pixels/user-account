package permission

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/serviceerr"
	"github.com/samber/lo"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	"github.com/photo-pixels/user-account/internal/utils"
)

// Storage интерфейс хранения данных
type Storage interface {
	storage.Transactor
	SavePermission(ctx context.Context, permission model.Permission) error
	AddPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error
	SaveRole(ctx context.Context, role model.Role) error
	AddRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]model.Permission, error)
	GetRolePermissions(ctx context.Context, roleID uuid.UUID) ([]model.Permission, error)
	GetPermissionByName(ctx context.Context, name string) (model.Permission, error)
	GetRoleByName(ctx context.Context, name string) (model.Role, error)
	GetPermission(ctx context.Context, id uuid.UUID) (model.Permission, error)
	GetRole(ctx context.Context, nid uuid.UUID) (model.Role, error)
	GetRolePermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) (model.Permission, error)
	GetUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (model.Role, error)
}

// Service роли и права пользователей
type Service struct {
	logger   log.Logger
	storage  Storage
	validate *validator.Validate
}

// NewService новый сервис
func NewService(logger log.Logger,
	storage Storage,
) *Service {
	return &Service{
		logger:   logger.Named("auth_service"),
		storage:  storage,
		validate: utils.NewValidator(),
	}
}

// CreateRole создание новой роли
func (s *Service) CreateRole(ctx context.Context, form form.CreateRole) (dto.Role, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.Role{}, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверка роли с этим именем
	_, err := s.storage.GetRoleByName(ctx, form.Name)
	switch {
	case err == nil:
		return dto.Role{}, serviceerr.Conflictf("Permission %s aready exists", form.Name)
	case errors.Is(err, storage.ErrNotFound): // Значит не найдено, продолжаем
	default:
		return dto.Role{}, serviceerr.MakeErr(err, "s.storage.GetRoleByName")
	}

	role := model.Role{
		Base:        model.NewBase(),
		ID:          uuid.New(),
		Name:        form.Name,
		Description: form.Description,
	}

	if err := s.storage.SaveRole(ctx, role); err != nil {
		return dto.Role{}, serviceerr.MakeErr(err, "s.storage.SaveRole")
	}

	return mapToRole(role), nil
}

// CreatePermission создание новой пермиссии
func (s *Service) CreatePermission(ctx context.Context, form form.CreatePermission) (dto.Permission, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.Permission{}, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверка пермисии с этим именем
	_, err := s.storage.GetPermissionByName(ctx, form.Name)
	switch {
	case err == nil:
		return dto.Permission{}, serviceerr.Conflictf("Role %s aready exists", form.Name)
	case errors.Is(err, storage.ErrNotFound): // Значит не найдено, продолжаем
	default:
		return dto.Permission{}, serviceerr.MakeErr(err, "s.storage.GetPermissionByName")
	}

	permission := model.Permission{
		Base:        model.NewBase(),
		ID:          uuid.New(),
		Name:        form.Name,
		Description: form.Description,
	}

	if err := s.storage.SavePermission(ctx, permission); err != nil {
		return dto.Permission{}, serviceerr.MakeErr(err, "s.storage.SavePermission")
	}

	return mapToPermission(permission), nil
}

// AddPermissionToRole добавить пермиссию в роль
func (s *Service) AddPermissionToRole(ctx context.Context, form form.AddPermissionToRole) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверить существование RoleID
	_, err := s.storage.GetRole(ctx, form.RoleID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("Role %s not found", form.RoleID.String())
	default:
		serviceerr.MakeErr(err, "s.storage.GetRoleByName")
	}
	// Проверить существование PermissionID
	_, err = s.storage.GetPermission(ctx, form.PermissionID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("Permission %s not found", form.RoleID.String())
	default:
		serviceerr.MakeErr(err, "s.storage.GetPermission")
	}
	// Проверка наличии этой пермиссии в роли
	_, err = s.storage.GetRolePermission(ctx, form.RoleID, form.PermissionID)
	switch {
	case err == nil:
		return serviceerr.Conflictf("Permission %s aready exists", form.PermissionID.String())
	case errors.Is(err, storage.ErrNotFound):
	default:
		serviceerr.MakeErr(err, "s.storage.GetPermission")
	}

	if err := s.storage.AddPermissionToRole(ctx, form.RoleID, form.PermissionID); err != nil {
		return serviceerr.MakeErr(err, "s.storage.AddPermissionToRole")
	}

	return nil
}

// GetUserPermissions список пермисий пользователя
func (s *Service) GetUserPermissions(ctx context.Context, form form.GetUserPermissions) ([]dto.Permission, error) {
	if err := s.validate.Struct(form); err != nil {
		return nil, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	permissions, err := s.storage.GetUserPermissions(ctx, form.UserID)
	if err != nil {
		return nil, fmt.Errorf("s.storage.GetUserPermissions: %w", err)
	}

	return lo.Map(permissions, func(item model.Permission, _ int) dto.Permission {
		return mapToPermission(item)
	}), nil
}

// AddRoleToUser добавить роль пользователю
func (s *Service) AddRoleToUser(ctx context.Context, form form.AddRoleToUser) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверить существование RoleID
	_, err := s.storage.GetRole(ctx, form.RoleID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("Role %s not found", form.RoleID.String())
	default:
		serviceerr.MakeErr(err, "s.storage.GetRoleByName")
	}
	// Проврка наличии этой роли у пользователя
	_, err = s.storage.GetUserRole(ctx, form.UserID, form.RoleID)
	switch {
	case err == nil:
		return serviceerr.Conflictf("Role %s aready exists", form.RoleID.String())
	case errors.Is(err, storage.ErrNotFound):
	default:
		serviceerr.MakeErr(err, "s.storage.GetUserRole")
	}

	if err := s.storage.AddRoleToUser(ctx, form.UserID, form.RoleID); err != nil {
		return serviceerr.MakeErr(err, "s.storage.AddRoleToUser")
	}

	return nil
}
