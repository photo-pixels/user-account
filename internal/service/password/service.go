package password

import (
	"github.com/photo-pixels/platform/log"
	"golang.org/x/crypto/bcrypt"
)

// Service сервис для работы с паролями
type Service struct {
	logger log.Logger
}

// NewService создание сервиса для работы с паролями
func NewService(logger log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

// HashPassword получение хеша пароля
func (s *Service) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash сравнение пароля с хешем
func (s *Service) CheckPasswordHash(password string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}
