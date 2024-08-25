package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/photo-pixels/platform/basemodel"
	"github.com/samber/lo"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage/db"
)

func mapToken(item db.Token) model.Token {
	return model.Token{
		Base: basemodel.Base{
			CreateAt: item.CreatedAt,
			UpdateAt: item.UpdatedAt,
		},
		ID:        item.ID,
		UserID:    item.UserID,
		Title:     item.Title,
		Token:     item.Token,
		TokenType: item.TokenType,
		ExpiredAt: toTimePtr(item.ExpiredAt),
	}
}

// GetTokens получение токенов пользователя
func (a *Adapter) GetTokens(ctx context.Context, userID uuid.UUID) ([]model.Token, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetTokens(ctx, userID)
	if err != nil {
		return nil, printError(err)
	}

	return lo.Map(res, func(item db.Token, _ int) model.Token {
		return mapToken(item)
	}), nil
}

// SaveToken сохранение токена
func (a *Adapter) SaveToken(ctx context.Context, token model.Token) error {
	queries := a.getQueries(ctx)

	err := queries.SaveToken(ctx, db.SaveTokenParams{
		ID:        token.ID,
		UserID:    token.UserID,
		Title:     token.Title,
		Token:     token.Token,
		CreatedAt: token.CreateAt,
		UpdatedAt: token.UpdateAt,
		ExpiredAt: toTimestamptz(token.ExpiredAt),
		TokenType: token.TokenType,
	})

	if err != nil {
		if isAlreadyExist(err) {
			return ErrAlreadyExist
		}
		return printError(err)
	}

	return nil
}

// DeleteToken удаление токена
func (a *Adapter) DeleteToken(ctx context.Context, userID, tokenID uuid.UUID) error {
	queries := a.getQueries(ctx)

	_, err := queries.DeleteToken(ctx, db.DeleteTokenParams{
		ID:     tokenID,
		UserID: userID,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return printError(err)
	}

	return nil
}

// GetToken получение токена по токену
func (a *Adapter) GetToken(ctx context.Context, token string) (model.Token, error) {
	queries := a.getQueries(ctx)

	res, err := queries.GetToken(ctx, token)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Token{}, ErrNotFound
		}
		return model.Token{}, printError(err)
	}

	return mapToken(res), nil
}
