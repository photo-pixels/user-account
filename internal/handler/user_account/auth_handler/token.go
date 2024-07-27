package auth_handler

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	desc "github.com/photo-pixels/user-account/pkg/gen/api/user_account"
)

func mapAuthData(res dto.AuthData) *desc.AuthData {
	return &desc.AuthData{
		UserId:                 res.UserID.String(),
		AccessToken:            res.AccessToken,
		AccessTokenExpiration:  timestamppb.New(res.AccessTokenExpiration),
		RefreshToken:           res.RefreshToken,
		RefreshTokenExpiration: timestamppb.New(res.RefreshTokenExpiration),
		// TODO: Roles
	}
}
