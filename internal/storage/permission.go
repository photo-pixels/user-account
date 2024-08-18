package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage/db"
)

// SavePermission .
func (a *Adapter) SavePermission(ctx context.Context, permission model.Permission) error {
	queries := a.getQueries(ctx)

	err := queries.SavePermission(ctx, db.SavePermissionParams{
		ID:          permission.ID,
		CreatedAt:   permission.CreateAt,
		UpdatedAt:   permission.UpdateAt,
		Name:        permission.Name,
		Description: permission.Description,
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// AddPermissionToRole .
func (a *Adapter) AddPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	queries := a.getQueries(ctx)

	err := queries.AddPermissionToRole(ctx, db.AddPermissionToRoleParams{
		RoleID:       roleID,
		PermissionID: permissionID,
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// SaveRole .
func (a *Adapter) SaveRole(ctx context.Context, role model.Role) error {
	queries := a.getQueries(ctx)

	err := queries.SaveRole(ctx, db.SaveRoleParams{
		ID:          role.ID,
		CreatedAt:   role.CreateAt,
		UpdatedAt:   role.UpdateAt,
		Name:        role.Name,
		Description: role.Description,
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// AddRoleToUser .
func (a *Adapter) AddRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	queries := a.getQueries(ctx)

	err := queries.AddRoleToUser(ctx, db.AddRoleToUserParams{
		UserID: userID,
		RoleID: roleID,
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

func mapPermission(item db.Permission, _ int) model.Permission {
	return model.Permission{
		Base: model.Base{
			CreateAt: item.CreatedAt,
			UpdateAt: item.UpdatedAt,
		},
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
	}
}

func mapRole(item db.Role, _ int) model.Role {
	return model.Role{
		Base: model.Base{
			CreateAt: item.CreatedAt,
			UpdateAt: item.UpdatedAt,
		},
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
	}
}

// GetUserPermissions .
func (a *Adapter) GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]model.Permission, error) {
	queries := a.getQueries(ctx)

	items, err := queries.GetUserPermissions(ctx, userID)
	if err != nil {
		return nil, printError(err)
	}

	return lo.Map(items, mapPermission), nil
}

// GetRolePermissions .
func (a *Adapter) GetRolePermissions(ctx context.Context, roleID uuid.UUID) ([]model.Permission, error) {
	queries := a.getQueries(ctx)

	items, err := queries.GetRolePermissions(ctx, roleID)
	if err != nil {
		return nil, printError(err)
	}

	return lo.Map(items, mapPermission), nil
}

// GetRolePermission .
func (a *Adapter) GetRolePermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) (model.Permission, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetRolePermission(ctx, db.GetRolePermissionParams{
		RoleID:       roleID,
		PermissionID: permissionID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Permission{}, ErrNotFound
		}
		return model.Permission{}, printError(err)
	}
	return mapPermission(item, 0), nil
}

// GetPermissionByName .
func (a *Adapter) GetPermissionByName(ctx context.Context, name string) (model.Permission, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetPermissionByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Permission{}, ErrNotFound
		}
		return model.Permission{}, printError(err)
	}
	return mapPermission(item, 0), nil
}

// GetRoleByName .
func (a *Adapter) GetRoleByName(ctx context.Context, name string) (model.Role, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetRoleByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Role{}, ErrNotFound
		}
		return model.Role{}, printError(err)
	}
	return mapRole(item, 0), nil
}

// GetPermission .
func (a *Adapter) GetPermission(ctx context.Context, id uuid.UUID) (model.Permission, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetPermission(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Permission{}, ErrNotFound
		}
		return model.Permission{}, printError(err)
	}
	return mapPermission(item, 0), nil
}

// GetRole .
func (a *Adapter) GetRole(ctx context.Context, id uuid.UUID) (model.Role, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetRole(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Role{}, ErrNotFound
		}
		return model.Role{}, printError(err)
	}
	return mapRole(item, 0), nil
}

// GetUserRole .
func (a *Adapter) GetUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (model.Role, error) {
	queries := a.getQueries(ctx)
	item, err := queries.GetUserRole(ctx, db.GetUserRoleParams{
		UserID: userID,
		ID:     roleID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Role{}, ErrNotFound
		}
		return model.Role{}, printError(err)
	}
	return mapRole(item, 0), nil
}
