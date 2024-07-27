package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/photo-pixels/user-account/internal/model"
)

// User данные пользователя
type User struct {
	ID         uuid.UUID
	Firstname  string
	Surname    string
	Patronymic *string
	Email      string
	Status     model.AuthStatus
	CreatedAt  time.Time
}
