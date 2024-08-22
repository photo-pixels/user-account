package form

import (
	"time"

	"github.com/google/uuid"
)

// GetTokens получить токены пользователя
type GetTokens struct {
	UserID uuid.UUID `validate:"required,uuid"`
}

// CreateToken создание нового токена для пользователя
type CreateToken struct {
	Title        string         `validate:"required,min=3,max=128"`
	TokenType    string         `validate:"required,min=3,max=128"`
	UserID       uuid.UUID      `validate:"required,uuid"`
	TimeDuration *time.Duration `validate:"required,duration-min=5s,duration-max=4320h"`
}

// DeleteToken удаление токена
type DeleteToken struct {
	TokenID uuid.UUID `validate:"required,uuid"`
	UserID  uuid.UUID `validate:"required,uuid"`
}

// GetToken получить токен
type GetToken struct {
	Token string `validate:"required,min=32,max=64"`
}
