package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/basemodel"
)

// Token токен сервиса
type Token struct {
	basemodel.Base
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Token     string
	TokenType string
	ExpiredAt *time.Time
}
