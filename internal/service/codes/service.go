package codes

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/storage"
)

// Storage интерфейс хранения данных
type Storage interface {
	storage.Transactor
	SaveConfirmCode(ctx context.Context, confirmCode model.ConfirmCode) error
	GetActiveConfirmCode(ctx context.Context, code string, confirmType model.ConfirmCodeType) (model.ConfirmCode, error)
	UpdateConfirmCode(ctx context.Context, personID uuid.UUID, confirmCodeType model.ConfirmCodeType, update model.UpdateConfirmCode) error
}

// Service сервис для работы с кодами повержения
type Service struct {
	logger  log.Logger
	storage Storage
}

// NewService создание сервиса для работы с паролями
func NewService(logger log.Logger, storage Storage) *Service {
	return &Service{
		logger:  logger.Named("confirm_code_service"),
		storage: storage,
	}
}

// GetActiveConfirmCode получение активного кода подтверждения
func (s *Service) GetActiveConfirmCode(ctx context.Context, code string, confirmType model.ConfirmCodeType) (model.ConfirmCode, error) {
	res, err := s.storage.GetActiveConfirmCode(ctx, code, confirmType)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return model.ConfirmCode{}, fmt.Errorf("s.storage.GetActiveConfirmCode: %w", ErrNotFound)
		}
		return model.ConfirmCode{}, fmt.Errorf("s.storage.GetActiveConfirmCode: %w", err)
	}
	return res, nil
}

// SendConfirmCode отправка кода подтверждения
func (s *Service) SendConfirmCode(ctx context.Context, userID uuid.UUID, confirmType model.ConfirmCodeType) error {
	code, err := generateCode()
	if err != nil {
		return fmt.Errorf("utils.GenerateCode: %w", err)
	}
	confirmCode := model.ConfirmCode{
		Base:   model.NewBase(),
		Code:   code,
		UserID: userID,
		Type:   confirmType,
		Active: true,
	}

	// Сначала отправляем, потом сохраняем
	// Не страшно если два раза отправим, страшно если кода не будет в базе
	// TODO: Реализация отправки кодов
	fmt.Println("*** *** *** *** *** *** ***")
	fmt.Printf("Confirm code: %s, for personID: %s", confirmCode.Code, userID.String())
	fmt.Println("*** *** *** *** *** *** ***")

	if err = s.storage.SaveConfirmCode(ctx, confirmCode); err != nil {
		return fmt.Errorf("s.storage.SaveConfirmCode: %w", err)
	}

	return nil
}

// DeactivateCode деактивация кода подтверждения
func (s *Service) DeactivateCode(ctx context.Context, personID uuid.UUID, confirmCodeType model.ConfirmCodeType) error {
	update := model.UpdateConfirmCode{
		BaseUpdate: model.NewBaseUpdate(),
		Active:     model.NewUpdateField(false),
	}
	if err := s.storage.UpdateConfirmCode(ctx, personID, confirmCodeType, update); err != nil {
		return fmt.Errorf("s.storage.UpdateConfirmCode: %w", err)
	}
	return nil
}
