package token

import (
	"context"
	"errors"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/photo-pixels/platform/basemodel"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/serviceerr"
	"github.com/photo-pixels/platform/utils"
	"github.com/samber/lo"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
)

// Storage интерфейс хранения данных
type Storage interface {
	storage.Transactor
	GetTokens(ctx context.Context, userID uuid.UUID) ([]model.Token, error)
	SaveToken(ctx context.Context, token model.Token) error
	DeleteToken(ctx context.Context, userID, tokenID uuid.UUID) error
	GetToken(ctx context.Context, token string) (model.Token, error)
	GetUser(ctx context.Context, userID uuid.UUID) (model.User, error)
}

// Service сервис токенов
type Service struct {
	logger   log.Logger
	storage  Storage
	validate *validator.Validate
	trans    ut.Translator
}

// NewService новый сервис
func NewService(
	logger log.Logger,
	storage Storage,
) *Service {
	validate, trans := utils.NewValidator()
	return &Service{
		logger:   logger.Named("token_service"),
		storage:  storage,
		validate: validate,
		trans:    trans,
	}
}

// GetTokens получение токенов пользователя
func (s *Service) GetTokens(ctx context.Context, form form.GetTokens) ([]dto.Token, error) {
	if err := s.validate.Struct(form); err != nil {
		return nil, serviceerr.InvalidInput(s.trans, err, "GetTokens")
	}

	tokens, err := s.storage.GetTokens(ctx, form.UserID)
	if err != nil {
		return nil, serviceerr.MakeErr(err, "s.storage.GetTokens")
	}

	return lo.Map(tokens, func(item model.Token, _ int) dto.Token {
		return dto.Token{
			Title:     item.Title,
			TokenType: item.TokenType,
			UserID:    item.UserID,
			ExpiredAt: item.ExpiredAt,
		}
	}), nil
}

// CreateToken создание нового токена пользователя
func (s *Service) CreateToken(ctx context.Context, form form.CreateToken) (string, error) {
	if err := s.validate.Struct(form); err != nil {
		return "", serviceerr.InvalidInput(s.trans, err, "CreateToken")
	}

	_, err := s.storage.GetUser(ctx, form.UserID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return "", serviceerr.NotFoundf("User not found")
	default:
		return "", serviceerr.MakeErr(err, "s.storage.GetUser")
	}

	apiToken, err := generateAPIToken()
	if err != nil {
		return "", serviceerr.MakeErr(err, "generateAPIToken")
	}

	var expiresAt *time.Time
	if form.TimeDuration != nil {
		expiresAt = lo.ToPtr(time.Now().Add(*form.TimeDuration))
	}

	token := model.Token{
		Base:      basemodel.NewBase(),
		ID:        uuid.New(),
		UserID:    form.UserID,
		Title:     form.Title,
		Token:     apiToken,
		TokenType: form.TokenType,
		ExpiredAt: expiresAt,
	}

	err = s.storage.SaveToken(ctx, token)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrAlreadyExist):
		return "", serviceerr.Conflictf("token already exists")
	default:
		return "", serviceerr.MakeErr(err, " storage.SaveToken")
	}

	return apiToken, nil
}

// DeleteToken удаление токена пользователя
func (s *Service) DeleteToken(ctx context.Context, form form.DeleteToken) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInput(s.trans, err, "DeleteToken")
	}

	err := s.storage.DeleteToken(ctx, form.UserID, form.TokenID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("Token not found")
	default:
		return serviceerr.MakeErr(err, "s.storage.DeleteToken")
	}

	return nil
}

// GetToken получение токена
func (s *Service) GetToken(ctx context.Context, form form.GetToken) (dto.Token, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.Token{}, serviceerr.InvalidInput(s.trans, err, "DeleteToken")
	}

	item, err := s.storage.GetToken(ctx, form.Token)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.Token{}, serviceerr.NotFoundf("Token not found")
	default:
		return dto.Token{}, serviceerr.MakeErr(err, "s.storage.GetToken")
	}

	if item.ExpiredAt != nil && (*item.ExpiredAt).Before(time.Now()) {
		return dto.Token{}, serviceerr.PermissionDeniedf("token has expired")
	}

	return dto.Token{
		Title:     item.Title,
		TokenType: item.TokenType,
		UserID:    item.UserID,
		ExpiredAt: item.ExpiredAt,
	}, nil
}
