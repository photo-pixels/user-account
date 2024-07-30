package model

import (
	"github.com/google/uuid"
)

// AuthStatus статус авторизации
type AuthStatus string

const (
	// AuthStatusSentInvite был отправлен инвайт на вступление
	AuthStatusSentInvite AuthStatus = "SENT_INVITE"
	// AuthStatusNotActivated не активен
	AuthStatusNotActivated AuthStatus = "NOT_ACTIVATED"
	// AuthStatusActivated активен
	AuthStatusActivated AuthStatus = "ACTIVATED"
	// AuthStatusBlocked заблокирован
	AuthStatusBlocked AuthStatus = "BLOCKED"
)

// RefreshTokenStatus статус рефреш токена
type RefreshTokenStatus string

const (
	// RefreshTokenStatusActive active
	RefreshTokenStatusActive RefreshTokenStatus = "ACTIVE"
	// RefreshTokenStatusRevoked revoked
	RefreshTokenStatusRevoked RefreshTokenStatus = "REVOKED"
	// RefreshTokenStatusExpired expired
	RefreshTokenStatusExpired RefreshTokenStatus = "EXPIRED"
	// RefreshTokenStatusLogout logout
	RefreshTokenStatusLogout RefreshTokenStatus = "LOGOUT"
)

// Auth авторизация пользователя
type Auth struct {
	Base
	UserID       uuid.UUID
	Email        string
	PasswordHash []byte
	Status       AuthStatus
}

// UpdateAuth Обновление Auth
type UpdateAuth struct {
	BaseUpdate
	PasswordHash UpdateField[[]byte]
	Status       UpdateField[AuthStatus]
}

// RefreshToken структура рефреш токена
type RefreshToken struct {
	Base
	ID     uuid.UUID
	UserID uuid.UUID
	Status RefreshTokenStatus
}
