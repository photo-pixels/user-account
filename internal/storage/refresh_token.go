package storage

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/photo-pixels/platform/basemodel"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage/db"
)

// GetLastActiveRefreshToken .
func (a *Adapter) GetLastActiveRefreshToken(ctx context.Context, refreshTokenID uuid.UUID) (model.RefreshToken, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetLastActiveRefreshToken(ctx, refreshTokenID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.RefreshToken{}, ErrNotFound
		}
		return model.RefreshToken{}, printError(err)
	}

	return model.RefreshToken{
		Base: basemodel.Base{
			CreateAt: res.CreatedAt,
			UpdateAt: res.UpdatedAt,
		},
		ID:     res.ID,
		UserID: res.UserID,
		Status: model.RefreshTokenStatus(res.Status),
	}, nil
}

// SaveRefreshToken .
func (a *Adapter) SaveRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error {
	queries := a.getQueries(ctx)

	err := queries.SaveRefreshToken(ctx, db.SaveRefreshTokenParams{
		ID:        refreshToken.ID,
		UserID:    refreshToken.UserID,
		CreatedAt: refreshToken.CreateAt,
		UpdatedAt: refreshToken.UpdateAt,
		Status:    db.RefreshTokenStatus(refreshToken.Status),
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// UpdateRefreshTokenStatus .
func (a *Adapter) UpdateRefreshTokenStatus(ctx context.Context, refreshTokenID uuid.UUID, status model.RefreshTokenStatus) error {
	queries := a.getQueries(ctx)

	err := queries.UpdateRefreshTokenStatus(ctx, db.UpdateRefreshTokenStatusParams{
		ID:        refreshTokenID,
		UpdatedAt: time.Now(),
		Status:    db.RefreshTokenStatus(status),
	})

	if err != nil {
		return printError(err)
	}

	return nil
}
