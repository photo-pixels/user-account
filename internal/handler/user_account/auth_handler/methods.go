package auth_handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/photo-pixels/user-account/internal/handler"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

// EmailAvailable проверка доступности email
func (h *AuthHandler) EmailAvailable(ctx context.Context, request *desc.EmailAvailableRequest) (*desc.EmailAvailableResponse, error) {
	available, err := h.auth.EmailAvailable(ctx, form.EmailAvailableForm{
		Email: request.Email,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.EmailAvailable")
	}

	return &desc.EmailAvailableResponse{
		Available: available,
	}, nil
}

// SendInvite отправка инвайта для регистрации
func (h *AuthHandler) SendInvite(ctx context.Context, request *desc.SendInviteRequest) (*emptypb.Empty, error) {
	err := h.auth.SendInvite(ctx, form.SendInviteForm{
		Email: request.Email,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.SendInvite")
	}

	return &emptypb.Empty{}, nil
}

// ActivateInvite активация инвайта
func (h *AuthHandler) ActivateInvite(ctx context.Context, request *desc.ActivateInviteRequest) (*emptypb.Empty, error) {
	err := h.auth.ActivateInvite(ctx, form.ActivateInviteForm{
		FirstName:   request.Firstname,
		Surname:     request.Surname,
		Patronymic:  request.Patronymic,
		CodeConfirm: request.CodeConfirm,
		Password:    request.Password,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.ActivateInvite")
	}
	return &emptypb.Empty{}, nil
}

// Registration регистрация
func (h *AuthHandler) Registration(ctx context.Context, request *desc.RegistrationRequest) (*emptypb.Empty, error) {
	err := h.auth.Registration(ctx, form.RegisterForm{
		FirstName:  request.Firstname,
		Surname:    request.Surname,
		Patronymic: request.Patronymic,
		Email:      request.Email,
		Password:   request.Password,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.Registration")
	}
	return &emptypb.Empty{}, nil
}

// ActivateRegistration активация регистрации
func (h *AuthHandler) ActivateRegistration(ctx context.Context, request *desc.ActivateRegistrationRequest) (*emptypb.Empty, error) {
	err := h.auth.ActivateRegistration(ctx, form.ActivateRegisterForm{
		CodeConfirm: request.CodeConfirm,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.ActivateRegistration")
	}
	return &emptypb.Empty{}, nil
}

// Logout выйти из системы
func (h *AuthHandler) Logout(ctx context.Context, request *desc.LogoutRequest) (*emptypb.Empty, error) {
	err := h.auth.Logout(ctx, form.LogoutForm{
		Token: request.RefreshToken,
	})
	if err != nil {
		return nil, handler.HandleError(err, "h.auth.Logout")
	}
	return &emptypb.Empty{}, nil
}

// Login логин пользователя
func (h *AuthHandler) Login(ctx context.Context, request *desc.LoginRequest) (*desc.AuthData, error) {
	res, err := h.auth.Login(ctx, form.LoginForm{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, handler.HandleError(err, "p.authService.Login")
	}
	return mapAuthData(res), nil
}

// RefreshToken обновление токена
func (h *AuthHandler) RefreshToken(ctx context.Context, request *desc.RefreshTokenRequest) (*desc.AuthData, error) {
	res, err := h.auth.RefreshToken(ctx, form.RefreshForm{
		Token: request.RefreshToken,
	})
	if err != nil {
		return nil, handler.HandleError(err, "p.authService.RefreshToken")
	}
	return mapAuthData(res), nil
}
