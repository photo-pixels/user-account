package permission

import (
	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
)

func mapToPermission(item model.Permission) dto.Permission {
	return dto.Permission{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreateAt,
		UpdatedAt:   item.UpdateAt,
	}
}

func mapToRole(item model.Role) dto.Role {
	return dto.Role{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreateAt,
		UpdatedAt:   item.UpdateAt,
	}
}
