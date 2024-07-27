package user_handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/photo-pixels/platform/log"
	"github.com/photo-pixels/platform/server"
	"google.golang.org/grpc"

	"github.com/photo-pixels/user-account/internal/user_case/dto"
	desc "github.com/photo-pixels/user-account/pkg/gen/api/user_account"
)

// UserUserCase юзеркейс для работы с пользователями
type UserUserCase interface {
	GetUser(ctx context.Context, userID uuid.UUID) (dto.User, error)
}

// UserHandler хендлер для работы с пользователями
type UserHandler struct {
	desc.UserServiceServer
	logger log.Logger
	user   UserUserCase
}

// NewHandler новый хендлер
func NewHandler(logger log.Logger,
	user UserUserCase,
) *UserHandler {
	return &UserHandler{
		logger: logger.Named("user_handler"),
		user:   user,
	}
}

// RegistrationServerHandlers .
func (h *UserHandler) RegistrationServerHandlers(_ *http.ServeMux) {
}

// RegisterServiceHandlerFromEndpoint .
func (h *UserHandler) RegisterServiceHandlerFromEndpoint() server.HandlerFromEndpoint {
	return desc.RegisterAuthServiceHandlerFromEndpoint
}

// RegisterServiceServer .
func (h *UserHandler) RegisterServiceServer(server *grpc.Server) {
	desc.RegisterUserServiceServer(server, h)
}
