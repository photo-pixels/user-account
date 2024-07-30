package session_manager

import (
	"time"

	"github.com/google/uuid"
)

// PermissionSession данные прав
type PermissionSession struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// AccessSession данные авторизации
type AccessSession struct {
	UserID      uuid.UUID           `json:"user_id"`
	Permissions []PermissionSession `json:"permissions"`
}

// RefreshSession данные для обновления токена
type RefreshSession struct {
	RefreshTokenID uuid.UUID `json:"refresh_token_id"`
	UserID         uuid.UUID `json:"user_id"`
}

// Token токен с временем истечения
type Token struct {
	Token     string
	ExpiresAt time.Time
}
