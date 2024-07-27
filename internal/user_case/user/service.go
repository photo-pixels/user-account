package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/serviceerr"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
)

// Storage интерфейс хранения данных
type Storage interface {
	storage.Transactor
	GetUser(ctx context.Context, userID uuid.UUID) (model.User, error)
	GetAuth(ctx context.Context, userID uuid.UUID) (model.Auth, error)
}

// Service сервис пользователей
type Service struct {
	logger  log.Logger
	storage Storage
}

// NewService новый сервис пользователей
func NewService(logger log.Logger,
	storage Storage,
) *Service {
	return &Service{
		logger:  logger.Named("user_service"),
		storage: storage,
	}
}

// GetUser получение данных пользоватея
func (s *Service) GetUser(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	user, err := s.storage.GetUser(ctx, userID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.User{}, serviceerr.NotFoundf("User not found")
	default:
		return dto.User{}, serviceerr.MakeErr(err, "s.storage.GetUser")
	}

	auth, err := s.storage.GetAuth(ctx, userID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.User{}, serviceerr.NotFoundf("User not found")
	default:
		return dto.User{}, serviceerr.MakeErr(err, "s.storage.GetAuth")
	}

	return dto.User{
		ID:         user.ID,
		Firstname:  user.FirstName,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Email:      auth.Email,
		Status:     auth.Status,
		CreatedAt:  user.CreateAt,
	}, nil
}
