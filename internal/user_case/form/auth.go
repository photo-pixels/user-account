package form

// SendInviteForm отправление инвайта на присоединение в системе
type SendInviteForm struct {
	Email string `validate:"required,email"`
}

// ActivateInviteForm активация инвайта
type ActivateInviteForm struct {
	FirstName   string  `validate:"required,min=3,max=128"`
	Surname     string  `validate:"required,min=3,max=128"`
	Patronymic  *string `validate:"omitempty,min=3,max=128"`
	CodeConfirm string  `validate:"required,max=128"`
	Password    string  `validate:"required,min=8,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ"`
}

// LoginForm форма для логина
type LoginForm struct {
	Email    string `validate:"required,max=128"`
	Password string `validate:"required,max=128"`
}

// RegisterForm форма для регистрации
type RegisterForm struct {
	FirstName  string  `validate:"required,min=3,max=128"`
	Surname    string  `validate:"required,min=3,max=128"`
	Patronymic *string `validate:"omitempty,min=3,max=128"`
	Email      string  `validate:"required,max=128"`
	Password   string  `validate:"required,max=128"`
}

// ActivateRegisterForm активация регистрации
type ActivateRegisterForm struct {
	CodeConfirm string `validate:"required,max=128"`
}

// EmailAvailableForm форма для проверки доступности email
type EmailAvailableForm struct {
	Email string `validate:"required,max=128"`
}

// LogoutForm форма для logout
type LogoutForm struct {
	Token string `validate:"required,max=2048"`
}

// RefreshForm форма для рефреша авторизационных данных по рефреш токену
type RefreshForm struct {
	Token string `validate:"required,max=2048"`
}
