package auth_handler

import (
	"context"
	"net/http"

	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"
	"google.golang.org/grpc"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	"github.com/photo-pixels/user-account/internal/user_case/form"
	desc "github.com/photo-pixels/user-account/pkg/gen/user_account"
)

// AuthUserCase юзеркейс авторизации
type AuthUserCase interface {
	SendInvite(ctx context.Context, form form.SendInviteForm) error
	ActivateInvite(ctx context.Context, form form.ActivateInviteForm) error
	Registration(ctx context.Context, form form.RegisterForm) error
	ActivateRegistration(ctx context.Context, form form.ActivateRegisterForm) error
	Login(ctx context.Context, form form.LoginForm) (dto.AuthData, error)
	Logout(ctx context.Context, form form.LogoutForm) error
	EmailAvailable(ctx context.Context, form form.EmailAvailableForm) (bool, error)
	RefreshToken(ctx context.Context, form form.RefreshForm) (dto.AuthData, error)
}

// AuthHandler хендлер авторизации
type AuthHandler struct {
	desc.AuthServiceServer
	logger log.Logger
	auth   AuthUserCase
}

// NewHandler новый хендлер
func NewHandler(logger log.Logger,
	auth AuthUserCase,
) *AuthHandler {
	return &AuthHandler{
		logger: logger.Named("auth_handler"),
		auth:   auth,
	}
}

// RegistrationServerHandlers .
func (h *AuthHandler) RegistrationServerHandlers(_ *http.ServeMux) {
}

// RegisterServiceHandlerFromEndpoint .
func (h *AuthHandler) RegisterServiceHandlerFromEndpoint() server.HandlerFromEndpoint {
	return desc.RegisterAuthServiceHandlerFromEndpoint
}

// RegisterServiceServer регистрация
func (h *AuthHandler) RegisterServiceServer(server *grpc.Server) {
	desc.RegisterAuthServiceServer(server, h)
}
