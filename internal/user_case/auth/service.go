package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/serviceerr"

	"github.com/photo-pixels/user-account/internal/model"
	"github.com/photo-pixels/user-account/internal/service/codes"
	"github.com/photo-pixels/user-account/internal/service/session_manager"
	"github.com/photo-pixels/user-account/internal/storage"
	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	"github.com/photo-pixels/user-account/internal/utils"
)

// Config конфигурация авторизации
type Config struct {
	AllowRegistration bool `yaml:"allow_registration"`
}

// Storage интерфейс хранения данных
type Storage interface {
	storage.Transactor
	EmailExists(ctx context.Context, email string) (bool, error)
	SaveUser(ctx context.Context, user model.User) error
	SaveUserAuth(ctx context.Context, auth model.Auth) error
	GetAuth(ctx context.Context, userID uuid.UUID) (model.Auth, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, updateUser model.UpdateUser) error
	UpdateUserAuth(ctx context.Context, userID uuid.UUID, updateAuth model.UpdateAuth) error
	GetAuthByEmail(ctx context.Context, email string) (model.Auth, error)
	SaveRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error
	UpdateRefreshTokenStatus(ctx context.Context, refreshTokenID uuid.UUID, status model.RefreshTokenStatus) error
	GetLastActiveRefreshToken(ctx context.Context, refreshTokenID uuid.UUID) (model.RefreshToken, error)
	GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]model.Permission, error)
}

// ConfirmCodeService сервис кодов подтверждения
type ConfirmCodeService interface {
	GetActiveConfirmCode(ctx context.Context, code string, confirmType model.ConfirmCodeType) (model.ConfirmCode, error)
	SendConfirmCode(ctx context.Context, userID uuid.UUID, confirmType model.ConfirmCodeType) error
	DeactivateCode(ctx context.Context, userID uuid.UUID, confirmType model.ConfirmCodeType) error
}

// PasswordService сервис для работы с паролями
type PasswordService interface {
	HashPassword(password string) ([]byte, error)
	CheckPasswordHash(password string, hash []byte) bool
}

// SessionManagerService сервис для генерации jwt сессий
type SessionManagerService interface {
	CreateTokenByAccessSession(session session_manager.AccessSession) (session_manager.Token, error)
	CreateTokenByRefreshSession(refresh session_manager.RefreshSession) (session_manager.Token, error)
	GetRefreshSessionByToken(token string) (session_manager.RefreshSession, error)
}

// Service авторизации
type Service struct {
	cfg                Config
	logger             log.Logger
	storage            Storage
	validate           *validator.Validate
	confirmCodeService ConfirmCodeService
	passwordService    PasswordService
	sessionService     SessionManagerService
}

// NewService новый сервис
func NewService(logger log.Logger,
	storage Storage,
	cfg Config,
	confirmCodeService ConfirmCodeService,
	passwordService PasswordService,
	sessionService SessionManagerService,
) *Service {
	return &Service{
		cfg:                cfg,
		logger:             logger.Named("auth_service"),
		storage:            storage,
		validate:           utils.NewValidator(),
		confirmCodeService: confirmCodeService,
		passwordService:    passwordService,
		sessionService:     sessionService,
	}
}

// SendInvite Отправка приглашения зарегистрироваться
func (s *Service) SendInvite(ctx context.Context, form form.SendInviteForm) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}
	if emailExists, err := s.storage.EmailExists(ctx, form.Email); err != nil {
		return serviceerr.MakeErr(err, "s.storage.EmailExists")
	} else if emailExists {
		return serviceerr.Conflictf("email already exists")
	}

	newUser := model.User{
		Base: model.NewBase(),
		ID:   uuid.New(),
	}

	newAuth := model.Auth{
		Base:         model.NewBase(),
		UserID:       newUser.ID,
		Email:        form.Email,
		PasswordHash: []byte{},
		Status:       model.AuthStatusSentInvite,
	}

	err := s.storage.RunTransaction(ctx, func(ctxTx context.Context) error {
		if saveErr := s.storage.SaveUser(ctxTx, newUser); saveErr != nil {
			return fmt.Errorf("s.storage.SaveUser: %w", saveErr)
		}
		if saveErr := s.storage.SaveUserAuth(ctxTx, newAuth); saveErr != nil {
			return fmt.Errorf("s.storage.SaveUserAuth: %w", saveErr)
		}
		return nil
	})

	if err != nil {
		return serviceerr.MakeErr(err, " s.storage.RunTransaction")
	}

	err = s.confirmCodeService.SendConfirmCode(ctx, newUser.ID, model.ConfirmCodeTypeActivateInvite)
	if err != nil {
		return serviceerr.MakeErr(err, "s.confirmCodeService.SendConfirmCode")
	}

	return nil
}

// ActivateInvite активация инвайта
func (s *Service) ActivateInvite(ctx context.Context, form form.ActivateInviteForm) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Поиск кода подтверждения в базе
	code, err := s.confirmCodeService.GetActiveConfirmCode(ctx, form.CodeConfirm, model.ConfirmCodeTypeActivateInvite)
	switch {
	case err == nil:
	case errors.Is(err, codes.ErrNotFound):
		return serviceerr.NotFoundf("confirm code not found")
	default:
		return serviceerr.MakeErr(err, "s.confirmCodeService.GetActiveConfirmCode")
	}

	auth, err := s.storage.GetAuth(ctx, code.UserID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("User code not found")
	default:
		return serviceerr.MakeErr(err, "s.storage.GetAuth")
	}

	if auth.Status == model.AuthStatusActivated {
		return serviceerr.Conflictf("user already activated")
	}

	if auth.Status == model.AuthStatusBlocked {
		return serviceerr.PermissionDeniedf("user blocked")
	}

	// Генерация соли
	hash, err := s.passwordService.HashPassword(form.Password)
	if err != nil {
		return serviceerr.MakeErr(err, "s.passwordService.HashPassword")
	}

	form.FirstName = utils.TransformToName(form.FirstName)
	form.Surname = utils.TransformToName(form.Surname)
	form.Patronymic = utils.TransformToNamePtr(form.Patronymic)

	updateUser := model.UpdateUser{
		BaseUpdate: model.NewBaseUpdate(),
		FirstName:  model.NewUpdateField(form.FirstName),
		Surname:    model.NewUpdateField(form.Surname),
		Patronymic: model.NewUpdateField(form.Patronymic),
	}

	updateAuth := model.UpdateAuth{
		BaseUpdate:   model.NewBaseUpdate(),
		PasswordHash: model.NewUpdateField(hash),
		Status:       model.NewUpdateField(model.AuthStatusActivated),
	}

	err = s.storage.RunTransaction(ctx, func(ctxTx context.Context) error {
		if err = s.storage.UpdateUser(ctxTx, code.UserID, updateUser); err != nil {
			return serviceerr.MakeErr(err, "s.storage.UpdateUser")
		}

		if err = s.storage.UpdateUserAuth(ctxTx, code.UserID, updateAuth); err != nil {
			return serviceerr.MakeErr(err, "s.storage.UpdateUserAuth")
		}

		if err = s.confirmCodeService.DeactivateCode(ctxTx, code.UserID, code.Type); err != nil {
			return serviceerr.MakeErr(err, "s.confirmCodeService.DeactivateCode")
		}
		return nil
	})

	if err != nil {
		return serviceerr.MakeErr(err, "s.storage.RunTransaction")
	}

	return nil
}

// Registration регистрация нового пользователя
func (s *Service) Registration(ctx context.Context, form form.RegisterForm) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	if !s.cfg.AllowRegistration {
		return serviceerr.FailPreconditionf("Registration is not available")
	}

	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	if emailExists, err := s.storage.EmailExists(ctx, form.Email); err != nil {
		return serviceerr.MakeErr(err, "s.storage.EmailExists")
	} else if emailExists {
		return serviceerr.Conflictf("Email already exists")
	}

	// Генерация соли
	hash, err := s.passwordService.HashPassword(form.Password)
	if err != nil {
		return serviceerr.MakeErr(err, "s.passwordService.HashPassword")
	}

	newUser := model.User{
		Base: model.NewBase(),
		ID:   uuid.New(),
	}

	newAuth := model.Auth{
		Base:         model.NewBase(),
		UserID:       newUser.ID,
		Email:        form.Email,
		PasswordHash: hash,
		Status:       model.AuthStatusNotActivated,
	}

	err = s.storage.RunTransaction(ctx, func(ctxTx context.Context) error {
		if saveErr := s.storage.SaveUser(ctxTx, newUser); saveErr != nil {
			return fmt.Errorf("s.storage.SaveUser: %w", saveErr)
		}
		if saveErr := s.storage.SaveUserAuth(ctxTx, newAuth); saveErr != nil {
			return fmt.Errorf("s.storage.SaveUserAuth: %w", saveErr)
		}
		return nil
	})

	if err != nil {
		return serviceerr.MakeErr(err, " s.storage.RunTransaction")
	}

	err = s.confirmCodeService.SendConfirmCode(ctx, newUser.ID, model.ConfirmCodeTypeActivateRegistration)
	if err != nil {
		return serviceerr.MakeErr(err, "s.confirmCodeService.SendConfirmCode")
	}

	return nil
}

// ActivateRegistration активация инвайта
func (s *Service) ActivateRegistration(ctx context.Context, form form.ActivateRegisterForm) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	// Поиск кода подтверждения в базе
	code, err := s.confirmCodeService.GetActiveConfirmCode(ctx, form.CodeConfirm, model.ConfirmCodeTypeActivateRegistration)
	switch {
	case err == nil:
	case errors.Is(err, codes.ErrNotFound):
		return serviceerr.NotFoundf("Confirm code not found")
	default:
		return serviceerr.MakeErr(err, "s.confirmCodeService.GetActiveConfirmCode")
	}

	auth, err := s.storage.GetAuth(ctx, code.UserID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.NotFoundf("User code not found")
	default:
		return serviceerr.MakeErr(err, "s.storage.GetAuth")
	}

	if auth.Status == model.AuthStatusActivated {
		return serviceerr.Conflictf("User already activated")
	}

	if auth.Status == model.AuthStatusBlocked {
		return serviceerr.PermissionDeniedf("User blocked")
	}

	updateAuth := model.UpdateAuth{
		BaseUpdate: model.NewBaseUpdate(),
		Status:     model.NewUpdateField(model.AuthStatusActivated),
	}

	err = s.storage.RunTransaction(ctx, func(ctxTx context.Context) error {
		if err = s.storage.UpdateUserAuth(ctxTx, code.UserID, updateAuth); err != nil {
			return serviceerr.MakeErr(err, "s.storage.UpdateUserAuth")
		}

		if err = s.confirmCodeService.DeactivateCode(ctxTx, code.UserID, code.Type); err != nil {
			return serviceerr.MakeErr(err, "s.confirmCodeService.DeactivateCode")
		}
		return nil
	})

	if err != nil {
		return serviceerr.MakeErr(err, "s.storage.RunTransaction")
	}

	return nil
}

// Login авторизация пользователя
func (s *Service) Login(ctx context.Context, form form.LoginForm) (dto.AuthData, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.AuthData{}, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	userAuth, err := s.storage.GetAuthByEmail(ctx, form.Email)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.AuthData{}, serviceerr.PermissionDeniedf("Incorrect username or password")
	default:
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.storage.GetAuthByEmail")
	}

	if check := s.passwordService.CheckPasswordHash(form.Password, userAuth.PasswordHash); !check {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("Incorrect username or password")
	}

	if userAuth.Status == model.AuthStatusBlocked {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("User is blocked")
	}
	if userAuth.Status == model.AuthStatusNotActivated {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("Not activated user account")
	}

	return s.createAuthData(ctx, userAuth)
}

// Logout разлогинить пользователя по рефреш токену
func (s *Service) Logout(ctx context.Context, form form.LogoutForm) error {
	if err := s.validate.Struct(form); err != nil {
		return serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	refreshSession, err := s.sessionService.GetRefreshSessionByToken(form.Token)
	if err != nil {
		return serviceerr.PermissionDeniedErr(fmt.Errorf("invalid token"))
	}
	err = s.storage.UpdateRefreshTokenStatus(ctx, refreshSession.RefreshTokenID, model.RefreshTokenStatusLogout)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, storage.ErrNotFound):
		return serviceerr.PermissionDeniedErr(fmt.Errorf("invalid token"))
	default:
		return serviceerr.MakeErr(err, "s.storage.UpdateRefreshTokenStatus")
	}
}

// EmailAvailable проверка доступности email
func (s *Service) EmailAvailable(ctx context.Context, form form.EmailAvailableForm) (bool, error) {
	if err := s.validate.Struct(form); err != nil {
		return false, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	emailExists, err := s.storage.EmailExists(ctx, form.Email)
	if err != nil {
		return false, serviceerr.MakeErr(err, "s.storage.EmailExists")
	}
	return emailExists, nil
}

// RefreshToken обновление авторизации по рефрештокену
func (s *Service) RefreshToken(ctx context.Context, form form.RefreshForm) (dto.AuthData, error) {
	if err := s.validate.Struct(form); err != nil {
		return dto.AuthData{}, serviceerr.InvalidInputErr(err, "Invalid input parameters")
	}

	refreshSession, err := s.sessionService.GetRefreshSessionByToken(form.Token)
	if err != nil {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("invalid token")
	}
	refreshToken, err := s.storage.GetLastActiveRefreshToken(ctx, refreshSession.RefreshTokenID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.AuthData{}, serviceerr.PermissionDeniedErr(fmt.Errorf("invalid token"))
	default:
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.storage.GetLastActiveRefreshToken")
	}

	userAuth, err := s.storage.GetAuth(ctx, refreshToken.UserID)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.AuthData{}, serviceerr.PermissionDeniedf("Incorrect username or password")
	default:
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.storage.GetAuth")
	}

	if userAuth.Status == model.AuthStatusBlocked {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("User is blocked")
	}

	if userAuth.Status == model.AuthStatusNotActivated {
		return dto.AuthData{}, serviceerr.PermissionDeniedf("Not activated user account")
	}

	err = s.storage.UpdateRefreshTokenStatus(ctx, refreshSession.RefreshTokenID, model.RefreshTokenStatusRevoked)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNotFound):
		return dto.AuthData{}, serviceerr.PermissionDeniedErr(fmt.Errorf("invalid token"))
	default:
		return dto.AuthData{}, serviceerr.MakeErr(err, "s.storage.UpdateRefreshTokenStatus")
	}

	return s.createAuthData(ctx, userAuth)
}
