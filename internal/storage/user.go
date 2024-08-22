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

// GetUser .
func (a *Adapter) GetUser(ctx context.Context, userID uuid.UUID) (model.User, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, printError(err)
	}

	return model.User{
		Base: basemodel.Base{
			CreateAt: res.CreatedAt,
			UpdateAt: res.UpdatedAt,
		},
		ID:         res.ID,
		FirstName:  res.Firstname,
		Surname:    res.Surname,
		Patronymic: res.Patronymic,
	}, nil
}

// SaveUser .
func (a *Adapter) SaveUser(ctx context.Context, user model.User) error {
	queries := a.getQueries(ctx)

	err := queries.SaveUser(ctx, db.SaveUserParams{
		ID:         user.ID,
		CreatedAt:  user.CreateAt,
		UpdatedAt:  user.UpdateAt,
		Firstname:  user.FirstName,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// UpdateUser .
func (a *Adapter) UpdateUser(ctx context.Context, userID uuid.UUID, updateUser model.UpdateUser) error {
	tx := a.getTX(ctx)

	builder := sq.Update("people").
		Where(sq.Eq{"id": userID}).
		Set("updated_at", updateUser.UpdateAt)

	if updateUser.FirstName.NeedUpdate {
		builder = builder.Set("firstname", updateUser.FirstName.Value)
	}
	if updateUser.Surname.NeedUpdate {
		builder = builder.Set("surname", updateUser.Surname.Value)
	}
	if updateUser.Patronymic.NeedUpdate {
		builder = builder.Set("patronymic", updateUser.Patronymic.Value)
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
