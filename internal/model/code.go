package model

import (
	"github.com/google/uuid"
	"github.com/photo-pixels/platform/basemodel"
)

// ConfirmCodeType тип кода подтверждения
type ConfirmCodeType string

const (
	// ConfirmCodeTypeActivateInvite активация инвайта
	ConfirmCodeTypeActivateInvite ConfirmCodeType = "ACTIVATE_INVITE"
	// ConfirmCodeTypeActivateRegistration активация регистрации
	ConfirmCodeTypeActivateRegistration ConfirmCodeType = "ACTIVATE_REGISTRATION"
)

// ConfirmCode код подтверждения
type ConfirmCode struct {
	basemodel.Base
	Code   string
	UserID uuid.UUID
	Type   ConfirmCodeType
	Active bool
}

// UpdateConfirmCode Обновление Person
type UpdateConfirmCode struct {
	basemodel.BaseUpdate
	Active basemodel.UpdateField[bool]
}
