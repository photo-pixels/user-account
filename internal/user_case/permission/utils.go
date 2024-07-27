package permission

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/serviceerr"

	"github.com/photo-pixels/user-account/internal/model"
)

func (s *Service) checkRoleExistByID(ctx context.Context, roleID uuid.UUID) error {
	roles, err := s.storage.GetRoles(ctx, model.GetRoles{
		RoleIDIn: []uuid.UUID{roleID},
	})
	if err != nil {
		return fmt.Errorf("s.storage.GetRoles: %w", err)
	}
	if len(roles) == 0 {
		return serviceerr.NotFoundf("Role not found")
	}
	return nil
}

func (s *Service) checkPermissionExistByID(ctx context.Context, permissionID uuid.UUID) error {
	permissions, err := s.storage.GetPermissions(ctx, model.GetPermissions{
		PermissionIDIn: []uuid.UUID{permissionID},
	})
	if err != nil {
		return fmt.Errorf("s.storage.GetPermissions: %w", err)
	}
	if len(permissions) == 0 {
		return serviceerr.NotFoundf("Permission not found")
	}
	return nil
}

func (s *Service) checkConflictRoleName(ctx context.Context, roleName string) error {
	roles, err := s.storage.GetRoles(ctx, model.GetRoles{
		RoleNameIn: []string{roleName},
	})
	if err != nil {
		return serviceerr.MakeErr(err, "s.storage.GetRoles")
	}
	if len(roles) > 0 {
		return serviceerr.Conflictf("Role %s aready exists", roles[0].Name)
	}
	return nil
}

func (s *Service) checkConflictPermissionName(ctx context.Context, permissionName string) error {
	permissions, err := s.storage.GetPermissions(ctx, model.GetPermissions{
		PermissionNameIn: []string{permissionName},
	})
	if err != nil {
		return serviceerr.MakeErr(err, "s.storage.GetRoles")
	}
	if len(permissions) > 0 {
		return serviceerr.Conflictf("Permission %s aready exists", permissions[0].Name)
	}

	return nil
}

func (s *Service) checkConflictPermissionInRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	permissions, err := s.storage.GetPermissions(ctx, model.GetPermissions{
		RoleIDIn:       []uuid.UUID{roleID},
		PermissionIDIn: []uuid.UUID{permissionID},
	})
	if err != nil {
		return fmt.Errorf("s.storage.GetUserPermissions: %w", err)
	}
	if len(permissions) > 0 {
		return serviceerr.Conflictf("Permission %s aready exists", permissions[0].Name)
	}
	return nil
}

func (s *Service) checkConflictRoleInUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	roles, err := s.storage.GetRoles(ctx, model.GetRoles{
		UserIDIn: []uuid.UUID{userID},
		RoleIDIn: []uuid.UUID{roleID},
	})
	if err != nil {
		return serviceerr.MakeErr(err, "s.storage.GetRoles")
	}
	if len(roles) > 0 {
		return serviceerr.Conflictf("Role %s aready exists", roles[0].Name)
	}
	return nil
}
