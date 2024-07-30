package model

import (
	"fmt"

	"github.com/google/uuid"
)

// User информация о пользователе
type User struct {
	Base
	ID         uuid.UUID
	FirstName  string
	Surname    string
	Patronymic *string
}

// FullName полное имя человека
func (p User) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.Surname)
}

// UpdateUser параметры обновления пользователей
type UpdateUser struct {
	BaseUpdate
	FirstName  UpdateField[string]
	Surname    UpdateField[string]
	Patronymic UpdateField[*string]
}
