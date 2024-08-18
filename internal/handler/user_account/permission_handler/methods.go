package permission_handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/server"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/photo-pixels/user-account/internal/handler"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

// CreatePermission создание новой пермиссии
func (h *PermissionHandler) CreatePermission(ctx context.Context, request *desc.CreatePermissionRequest) (*desc.CreatePermissionResponse, error) {
	permission, err := h.permission.CreatePermission(ctx, form.CreatePermission{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.permission.SavePermission")
	}

	return &desc.CreatePermissionResponse{
		Permission: mapPermission(permission),
	}, nil
}

// GetUserPermissions получение пермиссий пользователей
func (h *PermissionHandler) GetUserPermissions(ctx context.Context, request *desc.GetUserPermissionsRequest) (*desc.GetUserPermissionsResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	permissions, err := h.permission.GetUserPermissions(ctx, form.GetUserPermissions{
		UserID: userID,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.permission.GetUserPermissions")
	}

	return &desc.GetUserPermissionsResponse{
		Permissions: lo.Map(permissions, func(item dto.Permission, _ int) *desc.Permission {
			return mapPermission(item)
		}),
	}, nil
}

// CreateRole создание роли
func (h *PermissionHandler) CreateRole(ctx context.Context, request *desc.CreateRoleRequest) (*desc.CreateRoleResponse, error) {
	role, err := h.permission.CreateRole(ctx, form.CreateRole{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.permission.SaveRole")
	}

	return &desc.CreateRoleResponse{
		Role: mapRole(role),
	}, nil
}

// AddPermissionToRole добавление пермисси к роли
func (h *PermissionHandler) AddPermissionToRole(ctx context.Context, request *desc.AddPermissionToRoleRequest) (*emptypb.Empty, error) {
	permissionID, err := uuid.Parse(request.PermissionId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("permissionID is invalid: %w", err))
	}

	roleID, err := uuid.Parse(request.RoleId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("roleID is invalid: %w", err))
	}

	err = h.permission.AddPermissionToRole(ctx, form.AddPermissionToRole{
		PermissionID: permissionID,
		RoleID:       roleID,
	})

	if err != nil {
		return nil, handler.HandleError(err, "h.permission.AddPermissionToRole")
	}

	return &emptypb.Empty{}, nil
}

// AddRoleToUser добавлении роли к пользователю
func (h *PermissionHandler) AddRoleToUser(ctx context.Context, request *desc.AddRoleToUserRequest) (*emptypb.Empty, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("userID is invalid: %w", err))
	}

	roleID, err := uuid.Parse(request.RoleId)
	if err != nil {
		return nil, server.ErrInvalidArgument(fmt.Errorf("roleID is invalid: %w", err))
	}

	err = h.permission.AddRoleToUser(ctx, form.AddRoleToUser{
		UserID: userID,
		RoleID: roleID,
	})

	if err != nil {
		return nil, handler.HandleError(err, "h.permission.AddRoleToUser")
	}

	return &emptypb.Empty{}, nil
}
