package token_handler

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

// TokenUserCase сервис токенов
type TokenUserCase interface {
	GetTokens(ctx context.Context, form form.GetTokens) ([]dto.Token, error)
	GetToken(ctx context.Context, form form.GetToken) (dto.Token, error)
	CreateToken(ctx context.Context, form form.CreateToken) (string, error)
	DeleteToken(ctx context.Context, form form.DeleteToken) error
}

// TokenHandler хендлер для работы с токенами
type TokenHandler struct {
	desc.TokenServiceServer
	logger log.Logger
	token  TokenUserCase
}

// NewHandler новый хендлер
func NewHandler(logger log.Logger,
	token TokenUserCase,
) *TokenHandler {
	return &TokenHandler{
		logger: logger.Named("user_handler"),
		token:  token,
	}
}

// RegistrationServerHandlers .
func (h *TokenHandler) RegistrationServerHandlers(_ *http.ServeMux) {
}

// RegisterServiceHandlerFromEndpoint .
func (h *TokenHandler) RegisterServiceHandlerFromEndpoint() server.HandlerFromEndpoint {
	return desc.RegisterTokenServiceHandlerFromEndpoint
}

// RegisterServiceServer .
func (h *TokenHandler) RegisterServiceServer(server *grpc.Server) {
	desc.RegisterTokenServiceServer(server, h)
}
