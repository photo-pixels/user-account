package permission

import (
	"context"
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
	CreatePermission(ctx context.Context, permission model.Permission) error
	AddPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error
	CreateRole(ctx context.Context, role model.Role) error
	AddRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	GetPermissions(ctx context.Context, params model.GetPermissions) ([]model.Permission, error)
	GetRoles(ctx context.Context, params model.GetRoles) ([]model.Role, error)
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
	if err := s.checkConflictRoleName(ctx, form.Name); err != nil {
		return dto.Role{}, err
	}

	role := model.Role{
		Base:        model.NewBase(),
		ID:          uuid.New(),
		Name:        form.Name,
		Description: form.Description,
	}

	if err := s.storage.CreateRole(ctx, role); err != nil {
		return dto.Role{}, serviceerr.MakeErr(err, "s.storage.CreateRole")
	}

	return mapToRole(role), nil
}

// CreatePermission создание новой пермиссии
func (s *Service) CreatePermission(ctx context.Context, form form.CreatePermission) (dto.Permission, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.Permission{}, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверка пермисии с этим именем
	if err := s.checkConflictPermissionName(ctx, form.Name); err != nil {
		return dto.Permission{}, err
	}

	permission := model.Permission{
		Base:        model.NewBase(),
		ID:          uuid.New(),
		Name:        form.Name,
		Description: form.Description,
	}

	if err := s.storage.CreatePermission(ctx, permission); err != nil {
		return dto.Permission{}, serviceerr.MakeErr(err, "s.storage.CreatePermission")
	}

	return mapToPermission(permission), nil
}

// AddPermissionToRole добавить пермиссию в роль
func (s *Service) AddPermissionToRole(ctx context.Context, form form.AddPermissionToRole) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Проверить существование RoleID
	if err := s.checkRoleExistByID(ctx, form.RoleID); err != nil {
		return err
	}
	// Проверить существование PermissionID
	if err := s.checkPermissionExistByID(ctx, form.PermissionID); err != nil {
		return err
	}
	// Проверка наличии этой пермиссии в роли
	if err := s.checkConflictPermissionInRole(ctx, form.RoleID, form.PermissionID); err != nil {
		return err
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

	permissions, err := s.storage.GetPermissions(ctx, model.GetPermissions{
		UserIDIn: []uuid.UUID{form.UserID},
	})
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
	if err := s.checkRoleExistByID(ctx, form.RoleID); err != nil {
		return err
	}
	// Проврка наличии этой роли у пользователя
	if err := s.checkConflictRoleInUser(ctx, form.UserID, form.RoleID); err != nil {
		return err
	}

	if err := s.storage.AddRoleToUser(ctx, form.UserID, form.RoleID); err != nil {
		return serviceerr.MakeErr(err, "s.storage.AddRoleToUser")
	}

	return nil
}
