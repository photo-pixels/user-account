package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/basemodel"
)

// User информация о пользователе
type User struct {
	basemodel.Base
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
	basemodel.BaseUpdate
	FirstName  basemodel.UpdateField[string]
	Surname    basemodel.UpdateField[string]
	Patronymic basemodel.UpdateField[*string]
}
