package permission_handler

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

func mapPermission(item dto.Permission) *desc.Permission {
	return &desc.Permission{
		Id:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   timestamppb.New(item.CreatedAt),
		UpdatedAt:   timestamppb.New(item.UpdatedAt),
	}
}

func mapRole(item dto.Role) *desc.Role {
	return &desc.Role{
		Id:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   timestamppb.New(item.CreatedAt),
		UpdatedAt:   timestamppb.New(item.UpdatedAt),
	}
}
