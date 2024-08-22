package storage

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/photo-pixels/platform/basemodel"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage/db"
)

// SaveUserAuth .
func (a *Adapter) SaveUserAuth(ctx context.Context, auth model.Auth) error {
	queries := a.getQueries(ctx)

	err := queries.SavePersonAuth(ctx, db.SavePersonAuthParams{
		UserID:       auth.UserID,
		CreatedAt:    auth.CreateAt,
		UpdatedAt:    auth.UpdateAt,
		Email:        auth.Email,
		PasswordHash: auth.PasswordHash,
		Status:       db.AuthStatus(auth.Status),
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// GetAuth .
func (a *Adapter) GetAuth(ctx context.Context, userID uuid.UUID) (model.Auth, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetAuth(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Auth{}, ErrNotFound
		}
		return model.Auth{}, printError(err)
	}

	return model.Auth{
		Base: basemodel.Base{
			CreateAt: res.CreatedAt,
			UpdateAt: res.UpdatedAt,
		},
		UserID:       res.UserID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Status:       model.AuthStatus(res.Status),
	}, nil
}

// UpdateUserAuth .
func (a *Adapter) UpdateUserAuth(ctx context.Context, userID uuid.UUID, updateAuth model.UpdateAuth) error {
	tx := a.getTX(ctx)

	builder := sq.Update("auth").
		Where(sq.Eq{"user_id": userID}).
		Set("updated_at", updateAuth.UpdateAt)

	if updateAuth.PasswordHash.NeedUpdate {
		builder = builder.Set("password_hash", updateAuth.PasswordHash.Value)
	}
	if updateAuth.Status.NeedUpdate {
		builder = builder.Set("status", db.AuthStatus(updateAuth.Status.Value))
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return printError(err)
	}

	return nil
}

// GetAuthByEmail .
func (a *Adapter) GetAuthByEmail(ctx context.Context, email string) (model.Auth, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetAuthByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Auth{}, ErrNotFound
		}
		return model.Auth{}, printError(err)
	}

	return model.Auth{
		Base: basemodel.Base{
			CreateAt: res.CreatedAt,
			UpdateAt: res.UpdatedAt,
		},
		UserID:       res.UserID,
		Email:        res.Email,
		PasswordHash: res.PasswordHash,
		Status:       model.AuthStatus(res.Status),
	}, nil
}

// EmailExists .
func (a *Adapter) EmailExists(ctx context.Context, email string) (bool, error) {
	queries := a.getQueries(ctx)

	res, err := queries.EmailExists(ctx, email)
	if err != nil {
		return false, printError(err)
	}

	return res > 0, nil
}
