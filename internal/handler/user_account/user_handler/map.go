package user_handler

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

func mapToAuthStatus(status model.AuthStatus) (desc.AuthStatus, error) {
	switch status {
	case model.AuthStatusSentInvite:
		return desc.AuthStatus_AUTH_STATUS_SENT_INVITE, nil
	case model.AuthStatusNotActivated:
		return desc.AuthStatus_AUTH_STATUS_NOT_ACTIVATED, nil
	case model.AuthStatusActivated:
		return desc.AuthStatus_AUTH_STATUS_ACTIVATED, nil
	case model.AuthStatusBlocked:
		return desc.AuthStatus_AUTH_STATUS_BLOCKED, nil
	default:
		return desc.AuthStatus_AUTH_STATUS_UNKNOWN, fmt.Errorf("invalid auth status: %s", status)
	}
}

func mapUserResponse(user dto.User) (*desc.GetUserResponse, error) {
	status, err := mapToAuthStatus(user.Status)
	if err != nil {
		return nil, err
	}
	return &desc.GetUserResponse{
		Id:         user.ID.String(),
		Status:     status,
		Firstname:  user.Firstname,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Email:      user.Email,
		CreatedAt:  timestamppb.New(user.CreatedAt),
	}, nil
}
