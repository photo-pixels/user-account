package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/serviceerr"
	"github.com/samber/lo"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/service/session_manager"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
)

func (s *Service) createAuthData(ctx context.Context, personAuth model.Auth) (dto.AuthData, error) {
	refreshToken := model.RefreshToken{
		Base:   model.NewBase(),
		ID:     uuid.New(),
		UserID: personAuth.UserID,
		Status: model.RefreshTokenStatusActive,
	}

	if err := s.storage.SaveRefreshToken(ctx, refreshToken); err != nil {
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.storage.SaveOrCreateRefreshSession")
	}

	permissions, err := s.storage.GetUserPermissions(ctx, personAuth.UserID)
	if err != nil {
		return dto.AuthData{}, fmt.Errorf("s.storage.GetUserPermissions: %w", err)
	}

	session := session_manager.AccessSession{
		UserID: personAuth.UserID,
		Permissions: lo.Map(permissions, func(item model.Permission, _ int) session_manager.PermissionSession {
			return session_manager.PermissionSession{
				ID:   item.ID,
				Name: item.Name,
			}
		}),
	}

	access, err := s.sessionService.CreateTokenByAccessSession(session)
	if err != nil {
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.sessionService.CreateTokenBySession")
	}

	refreshSession := session_manager.RefreshSession{
		RefreshTokenID: refreshToken.ID,
		UserID:         refreshToken.UserID,
	}

	refresh, err := s.sessionService.CreateTokenByRefreshSession(refreshSession)
	if err != nil {
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.sessionService.CreateTokenByRefresh")
	}

	return dto.AuthData{
		UserID:                 personAuth.UserID,
		AccessToken:            access.Token,
		AccessTokenExpiration:  access.ExpiresAt,
		RefreshToken:           refresh.Token,
		RefreshTokenExpiration: refresh.ExpiresAt,
	}, nil
}
