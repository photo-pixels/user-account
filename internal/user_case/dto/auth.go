package dto

import (
	"time"

	"github.com/google/uuid"
)

// AuthData токены авторизации
type AuthData struct {
	UserID                 uuid.UUID
	AccessToken            string
	AccessTokenExpiration  time.Time
	RefreshToken           string
	RefreshTokenExpiration time.Time
}
