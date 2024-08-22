package dto

import (
	"time"

	"github.com/google/uuid"
)

// Token токен сервиса
type Token struct {
	Title     string
	TokenType string
	UserID    uuid.UUID
	ExpiredAt *time.Time
}
