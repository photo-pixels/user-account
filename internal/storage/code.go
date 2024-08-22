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

// SaveConfirmCode .
func (a *Adapter) SaveConfirmCode(ctx context.Context, confirmCode model.ConfirmCode) error {
	queries := a.getQueries(ctx)

	err := queries.SaveConfirmCode(ctx, db.SaveConfirmCodeParams{
		Code:      confirmCode.Code,
		UserID:    confirmCode.UserID,
		CreatedAt: confirmCode.CreateAt,
		UpdatedAt: confirmCode.UpdateAt,
		Active:    confirmCode.Active,
		Type:      db.CodeType(confirmCode.Type),
	})

	if err != nil {
		return printError(err)
	}

	return nil
}

// GetActiveConfirmCode .
func (a *Adapter) GetActiveConfirmCode(ctx context.Context, code string, confirmType model.ConfirmCodeType) (model.ConfirmCode, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetActiveConfirmCode(ctx, db.GetActiveConfirmCodeParams{
		Code: code,
		Type: db.CodeType(confirmType),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ConfirmCode{}, ErrNotFound
		}
		return model.ConfirmCode{}, printError(err)
	}

	return model.ConfirmCode{
		Base: basemodel.Base{
			CreateAt: res.CreatedAt,
			UpdateAt: res.UpdatedAt,
		},
		Code:   res.Code,
		UserID: res.UserID,
		Type:   model.ConfirmCodeType(res.Type),
		Active: res.Active,
	}, nil
}

// UpdateConfirmCode .
func (a *Adapter) UpdateConfirmCode(ctx context.Context, userID uuid.UUID, confirmCodeType model.ConfirmCodeType, update model.UpdateConfirmCode) error {
	tx := a.getTX(ctx)

	builder := sq.Update("code").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"type": db.CodeType(confirmCodeType)}).
		Set("updated_at", update.UpdateAt)

	if update.Active.NeedUpdate {
		builder = builder.Set("active", update.Active.Value)
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
